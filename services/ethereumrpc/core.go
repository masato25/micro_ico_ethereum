package ethereumrpc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/masato25/micro_ico_ethereum/config"
	"github.com/masato25/micro_ico_ethereum/contracts/token"
	"github.com/masato25/micro_ico_ethereum/lib/cryptolib"
	log "github.com/sirupsen/logrus"
)

var d = time.Now().Add(40000 * time.Millisecond)
var conn *ethclient.Client

// test key copy from truffle will delete 23::29 when compelete this code base
const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	address          = `0x627306090abaB3A6e1400e9345bC60c78a8BEf57`
	key              = `c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3`
	password         = `password`
)

func SetTimeOut() {
	conf := config.MyConfig().Ether
	timeMs := conf.TimeOutMS
	d = time.Now().Add(time.Duration(timeMs) * time.Millisecond)
}

func DepolyContactToBlockChain(totallsupply int64, Name string, tokenSymbol string) {
	auth, err := cryptolib.GenerateKeyStore(address, key, password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, tokenc, err := token.DeployToken(auth, conn, big.NewInt(int64(totallsupply)), Name, 0, tokenSymbol)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	name, err := tokenc.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}

// 0xf12b5dd4ead5f743c6baa640b0216200e89b60da
func GetContractName(contractAddr string) {
	// Instantiate the contract and display its name
	tokenc, err := token.NewToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := tokenc.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)
}

func invokeContract(contractAddr string, action string) (err error) {
	auth, err := cryptolib.GenerateKeyStore(address, key, password)
	tokenc, err := token.NewToken(common.HexToAddress(contractAddr), conn)
	if strings.Contains(action, "transfer") {
		params := strings.Split(action, " ")
		var amount int
		amount, err = strconv.Atoi(params[2])
		tx, err := tokenc.Transfer(auth, common.HexToAddress(params[1]), big.NewInt(int64(amount)))
		if err != nil {
			log.Fatalf("Failed to request token transfer: %v", err)
		}
		log.Infof("Transfer pending: 0x%x\n", tx.Hash())
	} else if strings.Contains(action, "blanceOf") {
		var result *big.Int
		result, err = tokenc.BalanceOf(&bind.CallOpts{Pending: true}, auth.From)
		log.Infof("result: %s", result)
	} else {
		err = fmt.Errorf("action not found")
		log.Error(err.Error())
	}
	return
}

func Conn() (err error) {
	conf := config.MyConfig()
	// Create an IPC based RPC connection to a remote node
	conn, err = ethclient.Dial(conf.Ether.RPCHost)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return
	// DepolyContactToBlockChain(conn)
	// GetContractName(conn, "0xf25186b5081ff5ce73482ad761db0eb0d25abfbf")
	// invokeContract(conn, "0xf25186b5081ff5ce73482ad761db0eb0d25abfbf", "blanceOf")
	// invokeContract(conn, "0xf25186b5081ff5ce73482ad761db0eb0d25abfbf", "transfer 0xf17f52151EbEF6C7334FAD080c5704D77216b732 200")
}

func GetConn() *ethclient.Client {
	return conn
}
