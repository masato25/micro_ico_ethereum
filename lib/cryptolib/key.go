package cryptolib

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func GenerateKeyStore(address string, keyhex string, passwd string) (auth *bind.TransactOpts, err error) {
	//conver hex to address
	addr := common.HexToAddress(address)
	//conver hex to ecdsa
	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA(keyhex)
	if err != nil {
		return
	}
	keygen := keystore.Key{
		Address:    addr,
		PrivateKey: privateKey,
	}
	var keyjson []byte
	keyjson, err = keystore.EncryptKey(&keygen, passwd, veryLightScryptN, veryLightScryptP)
	if err != nil {
		return
	}
	auth, err = bind.NewTransactor(strings.NewReader(string(keyjson)), passwd)
	return
}
