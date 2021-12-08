package service

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/ethereum-development-with-go/NimbleToken_interface"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math/big"
)
func (ep *Endpoint)connectWeb3()(client *ethclient.Client , err error)  {
	client, err = ethclient.Dial("http://127.0.0.1:8545")
	//client, err = ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")

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

func (ep *Endpoint) approve(client *ethclient.Client,session *TokenSession,eventAddress common.Address,reward *big.Int)(approve *types.Transaction,err error)  {
	instance, err := NimbleToken_interface.NewApi(nimbleToken, client)
	if err != nil {
		log.Errorf("[CreateEvent.NimbleToken_interface] : %+v",err)
	}

	approve,err = instance.Approve(session.TransactOpts,eventAddress,reward)
	if err != nil {
		log.Errorf("[CreateEvent.approve] : %+v",err)
	}
	return approve,nil
}