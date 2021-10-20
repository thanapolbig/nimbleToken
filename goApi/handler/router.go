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
			Pattern:     "/balanceOf",
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
		{
			Name:        "AddWorkday : POST ",
			Description: "AddWorkday",
			Method:      http.MethodPost,
			Pattern:     "/addWorkday",
			Endpoint:    service.AddWorkday,
		},
		{
			Name:        "ConfigMint : POST ",
			Description: "ConfigMint",
			Method:      http.MethodPost,
			Pattern:     "/configMint",
			Endpoint:    service.ConfigMint,
		},
		{
			Name:        "MintToken : POST ",
			Description: "MintToken",
			Method:      http.MethodPost,
			Pattern:     "/mintToken",
			Endpoint:    service.MintToken,
		},
		{
			Name:        "AutoClaimCheckin : POST ",
			Description: "AutoClaimCheckin",
			Method:      http.MethodPost,
			Pattern:     "/autoClaimCheckin",
			Endpoint:    service.AutoClaimCheckin,
		},
		{
			Name:        "AddVote : POST ",
			Description: "AddVote",
			Method:      http.MethodPost,
			Pattern:     "/addVote",
			Endpoint:    service.AddVote,
		},
		{
			Name:        "Vote : POST ",
			Description: "Vote",
			Method:      http.MethodPost,
			Pattern:     "/vote",
			Endpoint:    service.Vote,
		},
		{
			Name:        "CheckScoreVote : POST ",
			Description: "CheckScoreVote",
			Method:      http.MethodPost,
			Pattern:     "/checkScoreVote",
			Endpoint:    service.CheckScoreVote,
		},
		{
			Name:        "AutoClaimScoreVote : POST ",
			Description: "AutoClaimScoreVote",
			Method:      http.MethodPost,
			Pattern:     "/autoClaimScoreVote",
			Endpoint:    service.AutoClaimScoreVote,
		},
		{
			Name:        "CreateEvent : POST ",
			Description: "CreateEvent",
			Method:      http.MethodPost,
			Pattern:     "/createEvent",
			Endpoint:    service.CreateEvent,
		},
		{
			Name:        "StartEvent : POST ",
			Description: "StartEvent",
			Method:      http.MethodPost,
			Pattern:     "/startEvent",
			Endpoint:    service.StartEvent,
		},
		{
			Name:        "JoinEvent : POST ",
			Description: "JoinEvent",
			Method:      http.MethodPost,
			Pattern:     "/joinEvent",
			Endpoint:    service.JoinEvent,
		},
		{
			Name:        "CloseEvent : POST ",
			Description: "CloseEvent",
			Method:      http.MethodPost,
			Pattern:     "/closeEvent",
			Endpoint:    service.CloseEvent,
		},
		{
			Name:        "AcceptEvent : POST ",
			Description: "AcceptEvent",
			Method:      http.MethodPost,
			Pattern:     "/acceptEvent",
			Endpoint:    service.AcceptEvent,
		},
		{
			Name:        "CreateEventByAdmin : POST ",
			Description: "CreateEventByAdmin",
			Method:      http.MethodPost,
			Pattern:     "/createEventByAdmin",
			Endpoint:    service.CreateEventByAdmin,
		},

	}

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}