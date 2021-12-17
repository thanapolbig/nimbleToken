package handler

import (
	"github.com/miguelmota/ethereum-development-with-go/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Validator struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func NewValidator(conf *app.Configs, em *app.ErrorMessage) *Validator {
	return &Validator{
		CV: conf,
		EM: em,
	}
}

type serviceValue int


func (v *Validator) GetCharacterInfoPermit(c *gin.Context) {
	middleware := NewMiddleware(v.CV, v.EM)
	err := middleware.ValidateRequestHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	return
}

