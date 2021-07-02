package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Data struct {

}

type Input struct {
	Key 		string		`json:"key"`
	Value		int64		`json:"value"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	data := Data{}

	r.POST("/getBalance", data.GetBalance)
	//r.POST("/transferEth", data.transferEth)

	return r
}


func (d *Data)GetBalance(c *gin.Context)  {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
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



func main() {
	r := setupRouter()
	r.Run()
}

