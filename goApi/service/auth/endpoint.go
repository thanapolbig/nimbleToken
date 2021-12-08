package auth

import (
	"github.com/miguelmota/ethereum-development-with-go/app"
	"github.com/miguelmota/ethereum-development-with-go/service/loger"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
	EM *app.ErrorMessage
	CV *app.Configs
}

func NewEndpoint(conf *app.Configs, em *app.ErrorMessage) *Endpoint {
	return &Endpoint{
		CV: conf,
		EM: em,
	}
}

func (ep *Endpoint) SignIn(c *gin.Context) {
	defer c.Request.Body.Close()
	loginId := 0
	loger := loger.Initialize(loginId, "auth")

	var request credentials //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	loger.LogInfo("auth/signin", 0)
	loger.Debugf("Input Username: [%+v]", request.Username)

	srv := NewAuthBindingLog(ep.CV, ep.EM, loger)

	result, err := srv.EnCode(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result.AccessToken == "" {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) Register(c *gin.Context) {
	defer c.Request.Body.Close()
	loginId := 0
	loger := loger.Initialize(loginId, "auth")

	var request inputRegister
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		log.Errorf("ShouldBindBodyWith : %s", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	loger.LogInfo("auth/register", 0)
	loger.LogInfo(request.Username, 0)
	loger.Debugf("Input Username: [%+v]", request.Username)
	loger.Debugf("Input Email: [%+v]", request.Email)
	loger.Debugf("Input Gender: [%+v]", request.Gender)
	loger.Debugf("Input Skin Id: [%+v]", request.SkinId)
	loger.Debugf("Input Hat Id: [%+v]", request.HatId)
	loger.Debugf("Input Shirt Id: [%+v]", request.ShirtId)
	loger.Debugf("Input Shoes Id: [%+v]", request.ShoesId)

	srv := NewRegisterBindingLog(ep.CV, ep.EM, loger)
	msg, err := srv.RegisterTransaction(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(msg.Status, msg)
	return
}

func (ep *Endpoint) Refresh(c *gin.Context) {
	defer c.Request.Body.Close()

	var request RefreshInput //model รับ input จาก body
	log.Info("[Input : ...]")

	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	srv := NewAuthService(ep.CV, ep.EM)
	result, err := srv.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result.AccessToken == "" {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infof("AccessToken : %+s", result.AccessToken)
	log.Infof("RefreshToken : %+s", result.AccessToken)

	c.JSON(http.StatusOK, result)
	return

}