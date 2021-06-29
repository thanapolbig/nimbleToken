package main

import (

	"github.com/gin-gonic/gin"


)

type _Data struct {

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	data := _Data{}

	//r.POST("/getBalance", data.GetBalance)
	//r.POST("/transferEth", data.transferEth)
	r.GET("/ping", Pong)

	return r
}



func main() {
	r := setupRouter()
	r.Run()
}

