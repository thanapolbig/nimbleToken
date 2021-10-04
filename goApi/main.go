package main

import (
	"flag"
	"fmt"
	"github.com/miguelmota/ethereum-development-with-go/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

//
//type Data struct {
//
//}
//
//type Input struct {
//	PrivateKey 		string		`json:"key"`
//	Value		int64		`json:"value"`
//}
//
//func setupRouter() *gin.Engine {
//	r := gin.Default()
//	data := Data{}
//
//	r.POST("/getBalance", data.GetBalance)
//	r.POST("/transferEth", data.TransferEth)
//
//	return r
//}
//
//
//func (d *Data)GetBalance(c *gin.Context)  {
//	client, err := ethclient.Dial("http://127.0.0.1:8545")
//	if err != nil {
//		log.Fatal(err)
//	}
//	account := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
//	balance, err := client.BalanceAt(context.Background(), account, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(balance) // 25893860171173005034
//
//	fbalance := new(big.Float)
//	fbalance.SetString(balance.String())
//	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
//	fmt.Println(ethValue) // 25.729324269165216041
//
//	c.JSON(http.StatusOK, ethValue)
//}
//




func main() {
	//r := setupRouter()
	//r.Run()


	portApp := flag.String("portApp", "8080", "portApp number")
	//start http APP server
	r := handler.Routes{} //new object
	handleRoute := r.InitTransactionRoute()
	AppSrv := &http.Server{
		Addr:    fmt.Sprint(":", *portApp), //":8080"
		Handler: handleRoute,
	}
	go func() {
		if err := AppSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("transaction listen: %s\n", err)
		} else if err != nil {
			log.Panicf("transaction listen error: %s\n", err)
		}
		log.Infof("transaction listen at: %s", *portApp)
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals // wait for SIGINT


}

