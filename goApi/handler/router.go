package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelmota/ethereum-development-with-go/service"
	"net/http"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {

	service := service.NewEndpoint()

	r.transaction = []route{

		{
			Name:        "GetBalance : POST ",
			Description: "GetBalance",
			Method:      http.MethodPost,
			Pattern:     "/getBalance",
			Endpoint:    service.GetBalance,
		},
		{
			Name:        "TransferEth : POST ",
			Description: "TransferEth",
			Method:      http.MethodPost,
			Pattern:     "/transferEth",
			Endpoint:    service.TransferEth,
		},
		{
			Name:        "BalanceOf : POST ",
			Description: "BalanceOf",
			Method:      http.MethodPost,
			Pattern:     "/contractRead",
			Endpoint:    service.BalanceOf,
		},
		{
			Name:        "ContractMint : POST ",
			Description: "ContractMint",
			Method:      http.MethodPost,
			Pattern:     "/contractMint",
			Endpoint:    service.ContractMint,
		},
		{
			Name:        "Appove : POST ",
			Description: "Appove",
			Method:      http.MethodPost,
			Pattern:     "/appove",
			Endpoint:    service.Appove,
		},
		{
			Name:        "Allowance : POST ",
			Description: "Allowance",
			Method:      http.MethodPost,
			Pattern:     "/allowance",
			Endpoint:    service.Allowance,
		},
		{
			Name:        "TransferFrom : POST ",
			Description: "TransferFrom",
			Method:      http.MethodPost,
			Pattern:     "/claimReward",
			Endpoint:    service.TransferFrom,
		},
		{
			Name:        "Burn : POST ",
			Description: "Burn",
			Method:      http.MethodPost,
			Pattern:     "/burn",
			Endpoint:    service.Burn,
		},


	}

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}