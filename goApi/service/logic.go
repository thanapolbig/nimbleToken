package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/miguelmota/ethereum-development-with-go/NimbleToken_interface"
	"github.com/miguelmota/ethereum-development-with-go/interface"
	"golang.org/x/crypto/sha3"

	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"net/http"
)
type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}


func (ep *Endpoint)GetBalance(c *gin.Context)  {
	var request Input //model รับ input จาก body
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[GetBalance.connectWeb3] : %+v",err)
	}
	account := common.HexToAddress(request.Key)
	log.Infof("%s" , account)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893860171173005034

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	c.JSON(http.StatusOK, ethValue)
}


func (ep *Endpoint)GetBalanceToken(c *gin.Context)  {
	var request InputBalanceOf //model รับ input จาก body
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[GetBalance.connectWeb3] : %+v",err)
	}
	account := common.HexToAddress(request.Address)
	log.Infof("%s" , account)

	instance, err := NimbleToken_interface.NewApi(nimbleToken, client)
	if err != nil {
		log.Errorf("[GetBalance.NewApi] : %+v",err)
	}

	balanceOf, err := instance.BalanceOf(&bind.CallOpts{},account)
	if err != nil {
		log.Errorf("[GetBalance.BalanceOf] : %+v",err)
	}

	//balance, err := client.BalanceAt(context.Background(), account, nil)
	fmt.Println(balanceOf) // 25893860171173005034

	fbalance := new(big.Float)
	fbalance.SetString(balanceOf.String())
	tokenValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(3)))
	fmt.Println(tokenValue) // 25.729324269165216041

	c.JSON(http.StatusOK, tokenValue)
}

func (ep *Endpoint)TransferEth(c *gin.Context){

	//ดึงค่าจาก body
	var request Input //model รับ input จาก body
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("Body PrivateKey : %s", request.Key)
	log.Infof("Body Value  : %d", request.Value)


	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba")
	if err != nil {
		log.Fatal("error private : ",err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Infof("%s",fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s",nonce)

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(request.Key)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

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

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func (ep *Endpoint)TransferToken(c *gin.Context){

	//ดึงค่าจาก body
	var request InputAppoveTransfer //model รับ input จาก body
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client,err := ep.connectWeb3()
	if err != nil {
		log.Errorf("[TransferToken.connectWeb3] : %+v",err)
	}
	log.Infof("client : %s",client)
	privateKey,err := ep.connectPrivateKey(request.PrivateKey)
	if err != nil{
		log.Errorf("[TransferToken.connectPrivateKey] : %+v",err)
	}

	fromAddress,err := ep.convertWallet1(privateKey)
	if err != nil{
		log.Errorf("[TransferToken.convertWallet1] : %+v",err)
	}
	log.Infof("From : %s",fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(request.AddressSpender)
	//tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString(string(request.Value), 10)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &nimbleToken,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, nimbleToken, value, gasLimit, gasPrice, data)

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

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc

}

func (ep *Endpoint)BalanceOf(c *gin.Context)  {

	var request InputBalanceOf //model รับ input จาก body
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := _interface.NewApi(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instance)

	address := common.HexToAddress(request.Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{false,common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),nil,nil}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: nimble"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: nimble"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 3"

	fmt.Printf("wei: %s\n", bal) //

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) //

	c.JSON(http.StatusOK, value)
	return

}

func (ep *Endpoint)Appove(c *gin.Context) {
	var request InputAppoveTransfer
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(request.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("fromAddress : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(request.AddressSpender)
	fmt.Println(toAddress)
	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	transferFnSignature := []byte("approve(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("50000", 10) // จำนวน token

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

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

}

func (ep *Endpoint)Allowance(c *gin.Context)  {
	var request InputAllowance
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := _interface.NewApi(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}


	fromAddress	 := common.HexToAddress(request.FromAddress)
	toAddress := common.HexToAddress(request.ToAddress)

	Allowance, err := instance.Allowance(&bind.CallOpts{},fromAddress,toAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Allowance: %v\n", Allowance)
}

func (ep *Endpoint)TransferFrom(c *gin.Context) {
	var request InputClaimReward
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(request.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contracts := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	senderAddress := common.HexToAddress(request.Sender)
	log.Infof("%s",senderAddress)


	transferFnSignature := []byte("transferFrom(address,address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	//recipient	FromAddress
	log.Infof("Recipientaddress :  %s" , fromAddress)
	paddedRecipientAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedRecipientAddress))
	//sender TOAddress
	log.Infof("Senderaddress :  %s" , senderAddress)
	paddedSenderAddress := common.LeftPadBytes(senderAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedSenderAddress))
	//value
	amount := new(big.Int)
	amount.SetString(request.Value, 10) // จำนวน token

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedSenderAddress...)
	data = append(data, paddedRecipientAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &contracts,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	tx := types.NewTransaction(nonce, contracts, value, gasLimit, gasPrice, data)

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

}

func (ep *Endpoint)Burn(c *gin.Context){

	var request InputBurn
	log.Infof("input : %s", request)
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(request.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("fromAddress : %s",fromAddress)

	value := big.NewInt(0) // in wei (0 eth) จำนวน eth
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")

	transferFnSignature := []byte("burn(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(fromAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString(request.Value, 10) // จำนวน token

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

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

}