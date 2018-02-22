package token_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/micro_ico_ethereum/app/controller/common_helper"
	"github.com/masato25/micro_ico_ethereum/app/model/tokens"
	tokenc "github.com/masato25/micro_ico_ethereum/contracts/token"
	"github.com/masato25/micro_ico_ethereum/services/ethereumrpc"
)

type CreateErc20TokenInput struct {
	Name         string `form:"name" json:"name"`
	Symbol       string `form:"symbol" json:"symbol"`
	TotallSupply int64  `form:"totall_supply" json:"totall_supply"`
	EthKeyPass   string `form:"ethereum_password" json:"ethereum_password"`
}

func CreateErc20Token(c *gin.Context) {
	inputs := CreateErc20TokenInput{}
	if err := c.Bind(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//auth user session
	user, err := common_helper.GetUserBySession(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ethInstance, err := ethereumrpc.NewEthDepolyStruct(user.EthereumAddress, user.EthereumPrivateKeyEncode, inputs.EthKeyPass)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := ethInstance.DepolyContactToBlockChain(inputs.TotallSupply, inputs.Name, inputs.Symbol)
	err = createTokenItem(inputs, user.EthereumAddress, response.Address.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
	return
}

func createTokenItem(input CreateErc20TokenInput, ownerAddr, contractAddr string) (err error) {
	tokenItem := tokens.Token{
		TokenName:       input.Name,
		TokenSymbol:     input.Symbol,
		TotallSupply:    input.TotallSupply,
		ContractAddress: contractAddr,
		OwnerAddress:    ownerAddr,
		Abi:             tokenc.TokenABI,
	}
	dt := db.Create(&tokenItem)
	err = dt.Error
	return err
}

type TokenListJSONOuput struct {
	ID              string
	TokenName       string
	TokenSymbol     string
	ContractAddress string
	OwnerAddress    string
	TotallSupply    int64
	Abi             string
}

func TokenListJSON(c *gin.Context) {
	var minetokens []tokens.Token
	db.Find(&minetokens).Limit(10)
	var outputs []TokenListJSONOuput = make([]TokenListJSONOuput, len(minetokens))
	for indx, tos := range minetokens {
		outputs[indx] = TokenListJSONOuput{
			ID:              tos.IDString(),
			TokenName:       tos.TokenName,
			TokenSymbol:     tos.TokenSymbol,
			ContractAddress: tos.ContractAddress,
			OwnerAddress:    tos.OwnerAddress,
			TotallSupply:    tos.TotallSupply,
			Abi:             tos.Abi,
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": outputs})
	return
}

func TokenList(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "token_list", gin.H{"login": lg})
	return
}

func TokenCreate(c *gin.Context) {
	lg, _ := c.Get("login")
	c.HTML(http.StatusOK, "create_token", gin.H{"login": lg})
	return
}
