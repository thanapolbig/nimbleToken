package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/miguelmota/ethereum-development-with-go/app"
	"github.com/miguelmota/ethereum-development-with-go/service"
	"github.com/miguelmota/ethereum-development-with-go/service/auth"
	"net/http"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute(cv *app.Configs, em *app.ErrorMessage) http.Handler {

	service := service.NewEndpoint()
	validateRole := NewValidator(cv, em)

	r.transaction = []route{

		{
			Name:        "GetBalance : POST ",
			Description: "GetBalance",
			Method:      http.MethodPost,
			Pattern:     "/getBalance",
			Endpoint:    service.GetBalance,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "TransferEth : POST ",
			Description: "TransferEth",
			Method:      http.MethodPost,
			Pattern:     "/transferEth",
			Endpoint:    service.TransferEth,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "BalanceOf : POST ",
			Description: "BalanceOf",
			Method:      http.MethodPost,
			Pattern:     "/balanceOf",
			Endpoint:    service.BalanceOf,
			Validation:  validateRole.GetCharacterInfoPermit,
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
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AddWorkday : POST ",
			Description: "AddWorkday",
			Method:      http.MethodPost,
			Pattern:     "/addWorkday",
			Endpoint:    service.AddWorkday,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "ConfigMint : POST ",
			Description: "ConfigMint",
			Method:      http.MethodPost,
			Pattern:     "/configMint",
			Endpoint:    service.ConfigMint,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "MintToken : POST ",
			Description: "MintToken",
			Method:      http.MethodPost,
			Pattern:     "/mintToken",
			Endpoint:    service.MintToken,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AutoClaimCheckin : POST ",
			Description: "AutoClaimCheckin",
			Method:      http.MethodPost,
			Pattern:     "/autoClaimCheckin",
			Endpoint:    service.AutoClaimCheckin,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AddVote : POST ",
			Description: "AddVote",
			Method:      http.MethodPost,
			Pattern:     "/addVote",
			Endpoint:    service.AddVote,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "Vote : POST ",
			Description: "Vote",
			Method:      http.MethodPost,
			Pattern:     "/vote",
			Endpoint:    service.Vote,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "CheckScoreVote : POST ",
			Description: "CheckScoreVote",
			Method:      http.MethodPost,
			Pattern:     "/checkScoreVote",
			Endpoint:    service.CheckScoreVote,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AutoClaimScoreVote : POST ",
			Description: "AutoClaimScoreVote",
			Method:      http.MethodPost,
			Pattern:     "/autoClaimScoreVote",
			Endpoint:    service.AutoClaimScoreVote,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "CreateEvent : POST ",
			Description: "CreateEvent",
			Method:      http.MethodPost,
			Pattern:     "/createEvent",
			Endpoint:    service.CreateEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "StartEvent : POST ",
			Description: "StartEvent",
			Method:      http.MethodPost,
			Pattern:     "/startEvent",
			Endpoint:    service.StartEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "JoinEvent : POST ",
			Description: "JoinEvent",
			Method:      http.MethodPost,
			Pattern:     "/joinEvent",
			Endpoint:    service.JoinEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "CloseEvent : POST ",
			Description: "CloseEvent",
			Method:      http.MethodPost,
			Pattern:     "/closeEvent",
			Endpoint:    service.CloseEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AcceptEvent : POST ",
			Description: "AcceptEvent",
			Method:      http.MethodPost,
			Pattern:     "/acceptEvent",
			Endpoint:    service.AcceptEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "CreateEventByAdmin : POST ",
			Description: "CreateEventByAdmin",
			Method:      http.MethodPost,
			Pattern:     "/createEventByAdmin",
			Endpoint:    service.CreateEventByAdmin,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "GetBalanceToken : POST ",
			Description: "GetBalanceToken",
			Method:      http.MethodPost,
			Pattern:     "/getBalanceToken",
			Endpoint:    service.GetBalanceToken,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "AcceptEventAdmin : POST ",
			Description: "AcceptEventAdmin",
			Method:      http.MethodPost,
			Pattern:     "/acceptEventAdmin",
			Endpoint:    service.AcceptEventAdmin,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "EventInfo : POST ",
			Description: "EventInfo",
			Method:      http.MethodPost,
			Pattern:     "/eventInfo",
			Endpoint:    service.EventInfo,
			Validation:  validateRole.GetCharacterInfoPermit,
		},
		{
			Name:        "SearchEvent : POST ",
			Description: "SearchEvent",
			Method:      http.MethodPost,
			Pattern:     "/searchEvent",
			Endpoint:    service.SearchEvent,
			Validation:  validateRole.GetCharacterInfoPermit,
		},

	}

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern,e.Validation, e.Endpoint)
	}

	return ro
}

func (rAuth Routes) InitTransactionRouteAuth(cv *app.Configs, em *app.ErrorMessage) http.Handler {

	auth := auth.NewEndpoint(cv, em)

	txAuth := []route{
		{
			Name:        "register",
			Description: "register",
			Method:      http.MethodPost,
			Pattern:     "/register",
			Endpoint:    auth.Register,
		},
		{
			Name:        "signin",
			Description: "signin",
			Method:      http.MethodPost,
			Pattern:     "/signin",
			Endpoint:    auth.SignIn,
		},
		{
			Name:        "Refresh",
			Description: "Refresh",
			Method:      http.MethodPost,
			Pattern:     "/refresh",
			Endpoint:    auth.Refresh,
		},
	}

	roAuth := gin.New()

	storeAuth := roAuth.Group("/auth")
	for _, e := range txAuth {
		storeAuth.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return roAuth
}
