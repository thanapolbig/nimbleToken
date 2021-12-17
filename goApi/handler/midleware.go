package handler

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"

	"github.com/miguelmota/ethereum-development-with-go/service/loger"

	"github.com/dgrijalva/jwt-go"

	"github.com/miguelmota/ethereum-development-with-go/service/auth"

	"github.com/gin-gonic/gin"

	"github.com/miguelmota/ethereum-development-with-go/app"
)

type Middleware struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func NewMiddleware(conf *app.Configs, em *app.ErrorMessage) *Middleware {
	return &Middleware{
		CV: conf,
		EM: em,
	}
}

const apikey = "8af8e5a938ec9e8162ec532b77c3a0c3e3dbc1f61710ce5dbe7f51cf4018137a"

// ValidateRequestHeader validate for request header
func (m *Middleware) ValidateRequestHeader(c *gin.Context) (err error){
	//return func(c *gin.Context) {
	//log := logController.NewLogController()
	apiKey := c.Request.Header.Get("x-api-key")
	if apiKey == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "x-api-key Not Found")
		return
	}

	m.ValidateAPIKEY(apiKey, c)
	err = m.authWithClaims(c)
	return
	//}
}

func (m *Middleware) ValidateAPIKEY(k string, c *gin.Context) {
	//return func(c *gin.Context) {
	log := loger.NewLogController()

	if k != apikey {
		if k == "" {
			log.Error("no x-api-key")
			c.Abort()
			c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"Status":             http.StatusNotImplemented,
				"MessageCode":        "00000",
				"MessageDescription": "access denied",
			})
		} else {
			log.Error("x-api-key not correct")
			c.Abort()
			c.JSON(http.StatusNotImplemented, map[string]interface{}{
				"Status":             http.StatusNotImplemented,
				"MessageCode":        "00000",
				"MessageDescription": "access denied",
			})
		}
		return
	}
	return
}

func (m *Middleware) authWithClaims(c *gin.Context) (err error) {

	token, ok := c.Request.Header["Authorization"]
	if len(token) == 0 || !ok {
		err = m.EM.Auth.InvalidToken
		return
	}

	claims := &auth.Claims{}
	var verifyKey *rsa.PublicKey

	verifyBytes, err := ioutil.ReadFile(auth.PubKeyPath)
	if err != nil {
		err = m.EM.Auth.GetPublicKeyFail
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		err = m.EM.Auth.ParsePublicKeyFail
		return
	}
	tkn, err := jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if tkn == nil || !tkn.Valid {
		err = m.EM.Auth.AuthorizationExpiration
		return
	}
	c.Set("claim", claims)
	//check session
	session, err := auth.CheckSession(claims.Session,claims.WalletId)
	if err != nil || !session  {
		err = m.EM.Auth.GetSessionFail
		return
	}
	return

}
