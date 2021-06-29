package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"nimbleToken/goApi/service"
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
			Name:        "CallBack : POST ",
			Description: "CallBack",
			Method:      http.MethodPost,
			Pattern:     "/getBalance",
			Endpoint:    service.GetBalance,
		},
	}

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}