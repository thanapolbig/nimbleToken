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
		log.Errorf("[AddWorkday.connectWeb3]  : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[AddWorkday.connectPrivateKey] %+v: ",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[AddWorkday.convertWallet] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)


	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf("[AddWorkday.PendingNonceAt] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("[AddWorkday.SuggestGasPrice] : %+v",err)
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
		log.Errorf("[AddWorkday.EstimateGas] : %+v",err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Errorf("[AddWorkday.NetworkID] : %+v",err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Errorf("[AddWorkday.SignTx] : %+v",err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Errorf("[AddWorkday.SendTransaction] : %+v",err)
	}

	//check Value
	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[AddWorkday.NewApi] : %+v",err)
	}

	totalDay, err := instance.Workday(&bind.CallOpts{})
	if err != nil {
		log.Errorf("[AddWorkday.Workday] : %+v",err)
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
		log.Errorf("[ConfigMint.connectWeb3] : %+v",err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[ConfigMint.connectPrivateKey] : %+v",err)
	}
	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[ConfigMint.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf("[ConfigMint.PendingNonceAt] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("[ConfigMint.SuggestGasPrice] : %+v",err)
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
		log.Errorf("[ConfigMint.EstimateGas] : %+v",err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Errorf("[ConfigMint.NetworkID] : %+v",err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Errorf("[ConfigMint.SignTx] : %+v",err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Errorf("[ConfigMint.SendTransaction] : %+v",err)
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
		log.Errorf("[MintToken.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[MintToken.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[MintToken.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Errorf("[MintToken.PendingNonceAt] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("[MintToken.SuggestGasPrice] : %+v",err)
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
		log.Errorf("[MintToken.EstimateGas] : %+v",err)
	}

	tx := types.NewTransaction(nonce, masterChefAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Errorf("[MintToken.NetworkID] : %+v",err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Errorf("[MintToken.SignTx] : %+v",err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Errorf("[MintToken.SendTransaction] : %+v",err)
	}

	//check
	instance, err := NimbleToken_interface.NewApi(nimbleToken, client)
	if err != nil {
		log.Errorf("[MintToken.NewApi] : %+v",err)
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Errorf("[MintToken.TotalSupply] : %+v",err)
	}
	balanceOf, err := instance.BalanceOf(&bind.CallOpts{},syrupAddress)
	if err != nil {
		log.Errorf("[MintToken.BalanceOf] : %+v",err)
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
		log.Errorf("[AutoClaimCheckin.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[AutoClaimCheckin.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[AutoClaimCheckin.convertWallet1] : %+v",err)
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
		log.Errorf("[AutoClaimCheckin.NewApi] : %+v",err)
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
		log.Errorf("[AutoClaimCheckin.AutoClaimCheckin] : %+v",err)
	}
	log.Infof("claimCheckinLog : %s",claimCheckin)

	tokenInstance,err := NimbleToken_interface.NewApi(nimbleToken,client)
	if err != nil {
		log.Errorf("[AutoClaimCheckin.NewApi] : %+v",err)
	}

	balanceOf,err := tokenInstance.BalanceOf(session.CallOpts,list[0])
	if err != nil {
		log.Errorf("[AutoClaimCheckin.BalanceOf] : %+v",err)
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
		log.Errorf("[AddVote.connectWeb3] : %+v",err)
	}

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[AddVote.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[AddVote.convertWallet1] : %+v",err)
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
		log.Errorf("[AddVote.NewApi] : %+v",err)
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
		log.Errorf("[AddVote] : %+v",err)
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
		log.Errorf("[Vote.connectWeb3] : %+v",err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[Vote.connectPrivateKey] : %+v",err)
	}
	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[Vote.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[Vote.NewApi] : %+v",err)
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
		log.Errorf("[Vote.sendTransaction] : %+v",err)
	}

	c.JSON(http.StatusOK, vote)
	return
}

func (ep *Endpoint) CheckScoreVote(c *gin.Context)  {
	var request InputKeyValue
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[CheckScoreVote.connectWeb3] : %+v",err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[CheckScoreVote.connectPrivateKey] : %+v",err)
	}
	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[CheckScoreVote.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)
	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[CheckScoreVote.NewApi] : %+v",err)
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
	scoreVote,err := instance.GetScore(session.CallOpts,fromAddress)
	if err != nil {
		log.Errorf("[CheckScoreVote.GetScore] : %+v",err)
	}
	log.Println("scoreVote : ",scoreVote)

	rightScoreVote,err := instance.GetRightScore(session.CallOpts,fromAddress)
	if err != nil {
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
	}
	log.Println("rightScoreVote : ",rightScoreVote)

	score := ReturnScore{
		Status: "success",
		MessageCode: success,
		RightScore: rightScoreVote,
		ScoreVoteTotal: scoreVote,
	}

	c.JSON(http.StatusOK, score)

}

func (ep *Endpoint)AutoClaimScoreVote(c *gin.Context)  {

	var request InputKeyArray //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
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
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
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

	claimRewardScoreVote, err := instance.ClaimRewardScoreVote(session.TransactOpts,list)
	if err != nil {
		log.Errorf("[CheckScoreVote.GetRightScore] : %+v",err)
	}
	log.Infof("claimRewardScoreVote : %s", claimRewardScoreVote)


	c.JSON(http.StatusOK, "success")

	return
}

func (ep *Endpoint)CreateEvent(c *gin.Context)  {

	var request Event
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[CreateEvent] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[CreateEvent.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[CreateEvent.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[CreateEvent.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	t := request.TimeStart.Unix()
	log.Println("time : ",t)
	unixTime := new(big.Int)
	unixTime.SetInt64(t)
	log.Println("unixTime : ",unixTime)

	reward := new(big.Int)
	reward.SetString(request.Reward,10)
	log.Println("reward : " , reward)

	Auth := bind.NewKeyedTransactor(privateKey)
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

	approve,err := ep.approve(client,session,EventAddress,reward)
	if err != nil {
		log.Error(err)
	}
	log.Info(approve)

	//---------------
	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[CreateEvent.NewApi] : %+v",err)
	}


	CreateEvent, err := instance.CreateEvent(session.TransactOpts,request.EventName,request.Detial,reward,unixTime)
	if err != nil {
		log.Errorf("[CreateEvent.sendtransaction] : %+v",err)
	}
	log.Infof("CreateEvent : %s", CreateEvent.Data())


	c.JSON(http.StatusOK, "success")

	return
}

func (ep *Endpoint)StartEvent(c *gin.Context)  {

	var request InputKeyValue
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[CreateEvent] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[StartEvent.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[StartEvent.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[StartEvent.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[StartEvent.NewApi] : %+v",err)
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

	eid := new(big.Int)
	eid.SetString(request.Value,10)
	log.Println("eid : " , eid)

	StartEvent, err := instance.StartEvent(session.TransactOpts,eid)
	if err != nil {
		log.Errorf("[StartEvent.sendtransaction] : %+v",err)
	}
	log.Infof("StartEvent : %s", StartEvent)

	c.JSON(http.StatusOK, "start success")

	return
}

func (ep *Endpoint)JoinEvent(c *gin.Context)  {

	var request InputKeyValue
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[JoinEvent] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[JoinEvent.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[JoinEvent.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[JoinEvent.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[JoinEvent.NewApi] : %+v",err)
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

	eid := new(big.Int)
	eid.SetString(request.Value,10)
	log.Println("eid : " , eid)

	JoinEvent, err := instance.JoinEvent(session.TransactOpts,eid)
	if err != nil {
		log.Errorf("[JoinEvent.sendtransaction] : %+v",err)
	}
	log.Infof("JoinEvent : %s", JoinEvent)

	c.JSON(http.StatusOK, "JoinEvent success")

	return
}

func (ep *Endpoint)CloseEvent(c *gin.Context)  {

	var request InputKeyValue
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[CloseEvent] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[CloseEvent.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[CloseEvent.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[CloseEvent.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[CloseEvent.NewApi] : %+v",err)
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

	eid := new(big.Int)
	eid.SetString(request.Value,10)
	log.Println("eid : " , eid)

	JoinEvent, err := instance.CloseEvent(session.TransactOpts,eid)
	if err != nil {
		log.Errorf("[CloseEvent.sendtransaction] : %+v",err)
	}
	log.Infof("JoinEvent : %s", JoinEvent)

	c.JSON(http.StatusOK, "CloseEvent success")

	return
}

func (ep *Endpoint)AcceptEvent(c *gin.Context)  {

	var request AcceptEvent //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[AcceptEvent] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[AcceptEvent.GetRightScore] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[AcceptEvent.GetRightScore] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[AcceptEvent.GetRightScore] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)
	//list address
	var list []common.Address
	for i:=0;i< len(request.DataList);i++{
		list = append(list, common.HexToAddress(request.DataList[i]))
	}
	log.Infof("[AcceptEvent]list address : %s",list)
	//eid
	eid := new(big.Int)
	eid.SetString(request.Eid,10)
	log.Println("eid : " , eid)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[AcceptEvent.GetRightScore] : %+v",err)
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

	AcceptEvent, err := instance.AcceptEvent(session.TransactOpts,eid,list)
	if err != nil {
		log.Errorf("[AcceptEvent.GetRightScore] : %+v",err)
	}
	log.Infof("AcceptEvent : %s", AcceptEvent)

	c.JSON(http.StatusOK, "AcceptEvent success")

	return
}

func (ep *Endpoint)CreateEventByAdmin(c *gin.Context)  {

	var request Event
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[CreateEventByAdmin] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[CreateEventByAdmin.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[CreateEventByAdmin.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[CreateEventByAdmin.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	t := request.TimeStart.Unix()
	log.Println("time : ",t)
	unixTime := new(big.Int)
	unixTime.SetInt64(t)
	log.Println("unixTime : ",unixTime)

	reward := new(big.Int)
	reward.SetString(request.Reward,10)
	log.Println("reward : " , reward)



	Auth := bind.NewKeyedTransactor(privateKey)
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

	approve,err := ep.approve(client,session,EventAddress,reward)
	if err != nil {
		log.Error(err)
	}
	log.Info(approve)

	//---------------
	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[CreateEventByAdmin.NewApi] : %+v",err)
	}

	CreateEventByAdmin, err := instance.CreateEventAdmin(session.TransactOpts,request.EventName,request.Detial,reward,unixTime)
	if err != nil {
		log.Errorf("[CreateEventByAdmin.sendtransaction] : %+v",err)
	}
	log.Infof("CreateEventByAdmin : %s", CreateEventByAdmin.Data())

	c.JSON(http.StatusOK, "CreateEventByAdmin success")

	return
}

func (ep *Endpoint)AcceptEventAdmin(c *gin.Context)  {

	var request AcceptEvent //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[AcceptEventAdmin] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[AcceptEventAdmin.GetRightScore] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[AcceptEventAdmin.GetRightScore] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[AcceptEventAdmin.GetRightScore] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)
	//list address
	var list []common.Address
	for i:=0;i< len(request.DataList);i++{
		list = append(list, common.HexToAddress(request.DataList[i]))
	}
	log.Infof("[AcceptEventAdmin]list address : %s",list)
	//eid
	eid := new(big.Int)
	eid.SetString(request.Eid,10)
	log.Println("eid : " , eid)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[AcceptEventAdmin.GetRightScore] : %+v",err)
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

	AcceptEventAdmin, err := instance.AcceptEventAdmin(session.TransactOpts,eid,list)
	if err != nil {
		log.Errorf("[AcceptEventAdmin.GetRightScore] : %+v",err)
	}
	log.Infof("AcceptEventAdmin : %s", AcceptEventAdmin)

	c.JSON(http.StatusOK, "AcceptEventAdmin success")

	return
}

func (ep *Endpoint)EventInfo(c *gin.Context)  {

	var request InputKeyValue
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("[EventInfo] %+v",err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[EventInfo.connectWeb3] : %+v",err)
	}
	fmt.Println(client)

	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[EventInfo.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[EventInfo.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	Auth := bind.NewKeyedTransactor(privateKey)

	instance, err := _masterChef.NewApi(masterChefAddress, client)
	if err != nil {
		log.Errorf("[EventInfo.NewApi] : %+v",err)
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

	eid := new(big.Int)
	eid.SetString(request.Value,10)
	log.Println("eid : " , eid)

	EventInfo, err := instance.EventInfo(session.CallOpts,eid)
	if err != nil {
		log.Errorf("[CloseEvent.sendtransaction] : %+v",err)
	}
	log.Infof("EventInfo : %s", EventInfo)

	c.JSON(http.StatusOK, EventInfo)

	return
}
