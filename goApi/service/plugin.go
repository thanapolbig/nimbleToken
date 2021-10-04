package service

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)
func (ep *Endpoint)connectWeb3()(client *ethclient.Client , err error)  {
	client, err = ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client , nil
}

func (ep *Endpoint)connectPrivateKey(PrivateKey string)(privateKey *ecdsa.PrivateKey , err error)  {

	privateKey, err = crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil , errors.New(fmt.Sprintf("connectPrivateKey error : %s",err.Error()))
	}
	return privateKey , nil
}

func (ep *Endpoint) convertWallet1(privateKey *ecdsa.PrivateKey)(fromAddress common.Address , err error)  {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{},errors.New("cannot assert convertWallet")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress ,nil
}