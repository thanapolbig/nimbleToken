package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_masterChef "github.com/miguelmota/ethereum-development-with-go/MasterChef_interface"
	"github.com/miguelmota/ethereum-development-with-go/NimbleToken_interface"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"math/big"
	"net/http"
)

var nimbleToken = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
var syrupAddress = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
var EventAddress = common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
var masterChefAddress = common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9")

func (ep *Endpoint)AddWorkday(c *gin.Context)  {

	var request InputKeyValue //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)


	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("addWorkday(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0x99c94479

	amount := new(big.Int)
	amount.SetString(request.Value,10) // จำนวน day
	log.Println("amount : " , amount)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x0000000000000000000000000000000000000000000000000000000000000000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &masterChefAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//check Value
	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	totalDay, err := instance.Workday(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("totalDay: %v\n", totalDay)

	c.JSON(http.StatusOK, totalDay)
	return
}

func (ep *Endpoint)ConfigMint(c *gin.Context)  {

	var request InputKeyValue //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}
	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("configMint(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0x99c94479

	amount := new(big.Int)
	amount.SetString(request.Value,10) // จำนวน day
	log.Println("mint amount : " , amount)



	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x0000000000000000000000000000000000000000000000000000000000000000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)


	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &masterChefAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, "success")
	return
}

func (ep *Endpoint)MintToken(c *gin.Context)  {

	var request InputKeyValue //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}


	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("mint()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0x99c94479

	var data []byte
	data = append(data, methodID...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &masterChefAddress,
		Data: data,
	})
	fmt.Println("gasLimit : ", gasLimit)
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	//check
	instance, err := NimbleToken_interface.NewApi(nimbleToken, client)
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	balanceOf, err := instance.BalanceOf(&bind.CallOpts{},syrupAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("totalSupply: ", totalSupply)
	log.Println("balanceOf : ", balanceOf)
	c.JSON(http.StatusOK, totalSupply)
	return
}

func (ep *Endpoint)AutoClaimCheckin(c *gin.Context)  {

	var request InputKeyArray //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}


	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	var list []common.Address
	for i:=0;i< len(request.DataList);i++{
		list = append(list, common.HexToAddress(request.DataList[i]))
	}

	log.Infof("list address : %s",list)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	session := &TokenSession{
		Contract: masterChefAddress,
		CallOpts: &bind.CallOpts{
			Pending: true,
		},
		TransactOpts: &bind.TransactOpts{
			From:     Auth.From,
			Signer:   Auth.Signer,
			GasLimit: 3141592,
		},
	}

	claimCheckin, err := instance.AutoClaimCheckin(session.TransactOpts,list)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("claimCheckinLog : %s",claimCheckin)

	tokenInstance,err := NimbleToken_interface.NewApi(nimbleToken,client)
	if err != nil {
		log.Fatal(err)
	}

	balanceOf,err := tokenInstance.BalanceOf(session.CallOpts,list[0])
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, balanceOf)

	return
}

func (ep *Endpoint) AddVote(c *gin.Context)  {

	var request InputKeyArray //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	var list []common.Address
	for i:=0;i< len(request.DataList);i++{
		list = append(list, common.HexToAddress(request.DataList[i]))
	}

	log.Infof("list address : %s",list)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Fatal("instance masterChef error : ",err)
	}

	session := &TokenSession{
		Contract: masterChefAddress,
		CallOpts: &bind.CallOpts{
			Pending: true,
		},
		TransactOpts: &bind.TransactOpts{
			From:     Auth.From,
			Signer:   Auth.Signer,
			GasLimit: 3141592,
		},
	}

	AddVote, err := instance.AddVote(session.TransactOpts,list)
	if err != nil {
		log.Fatal("instance.AddVote error : ",err)
	}
	log.Infof("AddVote : %s",AddVote)

	c.JSON(http.StatusOK, AddVote)

	return
}


func (ep *Endpoint)Vote(c *gin.Context)  {
	var request Vote
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client,err := ep.connectWeb3()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Fatal(err)
	}
	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Fatal(err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Fatal("instance masterChef error : ",err)
	}

	session := &TokenSession{
		Contract: masterChefAddress,
		CallOpts: &bind.CallOpts{
			Pending: true,
		},
		TransactOpts: &bind.TransactOpts{
			From:     Auth.From,
			Signer:   Auth.Signer,
			GasLimit: 3141592,
		},
	}

	toAddress := common.HexToAddress(request.ToAddress)

	amount := new(big.Int)
	amount.SetString(request.Value,10)
	log.Println("mint amount : " , amount)

	vote,err := instance.Vote(session.TransactOpts,toAddress,amount)
	if err != nil {
		log.Fatal("instance.Vote error : ",err)
	}

	scoreVote,err := instance.GetScore(session.CallOpts,toAddress)
	if err != nil {
		log.Fatal("instance.scoreVote error : ",err)
	}
	log.Println("scoreVote : ",scoreVote)

	c.JSON(http.StatusOK, vote)
	return
}



