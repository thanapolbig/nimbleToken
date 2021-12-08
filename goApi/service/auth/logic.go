package auth

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/miguelmota/ethereum-development-with-go/app"
	"io/ioutil"
	"time"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm"

	"github.com/miguelmota/ethereum-development-with-go/service/loger"
	log "github.com/sirupsen/logrus"
)

var srv *authService

type authService struct {
	conf   *app.Configs
	em     *app.ErrorMessage
	repo   *authRepo
	claims *Claims
	rtclaims *rtClaims
	loger  *loger.Loger
}

func Init(conf *app.Configs, em *app.ErrorMessage) {
	srv = &authService{
		conf: conf,
		em:   em,
		repo: &authRepo{conf: conf},
	}
}

func NewAuthService(conf *app.Configs, em *app.ErrorMessage) *authService {
	repo := authRepo{}

	return &authService{
		conf: conf,
		em:   em,
		repo: repo.InitAuthRepo(conf, em),
	}
}

func NewAuthBindingLog(conf *app.Configs, em *app.ErrorMessage, loger *loger.Loger) *authService {
	repo := authRepo{}
	return &authService{
		conf: conf,
		em:   em,
		repo: repo.InitAuthRepo(conf, em),
		loger: loger,
	}
}

func encryptPassword(password string) string {
	s := password
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashPassword : " + sha1_hash)
	return sha1_hash
}

func compilePassword(password string, loginUuid string) string {
	//request password + loginUuid[15:20]
	//uuid Start 0
	detailPass := password + loginUuid[15:20]
	fmt.Println("compilePassword : " + detailPass)
	return detailPass
}

func (srv *authService) EnCode(request credentials) (result parseCode, err error) {
	var detail character
	detail, err = srv.repo.getCharacter(request)
	loginId, _ := strconv.ParseInt(detail.LoginId, 10, 64)
	
	srv.loger.LogInfo("Check Character In Database ...", int(loginId))
	if err != nil && err == gorm.ErrRecordNotFound {
		err = srv.em.Auth.LoginNotFound
		srv.loger.LogErrorf("Check Character Failed: [%+v]", err.Error(), int(loginId))
		return
	}
	srv.loger.LogDebugf("Detail Target: [%+v]", detail.Username, int(loginId))
	srv.loger.LogDebugf("Detail LoginUuid: [%+v]", detail.LoginUuid, int(loginId))
	srv.loger.LogDebugf("Detail CharacterId: [%+v]", detail.CharacterId, int(loginId))
	srv.loger.LogDebugf("Detail LoginIdOwnerFarm: [%+v]", detail.LoginId, int(loginId))
	srv.loger.LogDebugf("Detail Session: [%+v]", detail.Session, int(loginId))
	srv.loger.LogDebugf("Detail RoleName: [%+v]", detail.RoleName, int(loginId))

	srv.loger.LogInfo("Hashing Password ...", int(loginId))
	detailPass := compilePassword(request.Password, detail.LoginUuid)
	hashPass := encryptPassword(detailPass)

	srv.loger.LogInfo("Check Password In Database ...", int(loginId))
	// expect Password && Target
	if detail.Password != hashPass || request.Username != detail.Username {
		err = srv.em.Auth.InvalidUsernamePassword
		srv.loger.LogErrorf("Check Character Failed: [%+v]", err.Error(), int(loginId))
		//return err
		return
	}

	session := uuid.New().String()
	//expirationTime := time.Now().Add(time.Duration(viper.GetInt("expirationTime.expHour")) * time.Hour)

	srv.loger.LogInfo("Create Access Token ...", int(loginId))
	expirationTime := time.Now().Add(1 * time.Hour)
	//expirationTime := time.Now().Add(time.Duration(viper.GetInt("expirationTime.expHour")) * time.Hour)
	claims := &Claims {
		CharacterId: detail.CharacterId,
		Session:     session,
		RoleName:    detail.RoleName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token,err := srv.CreateToken(claims)
	if err != nil {
		err = srv.em.Auth.CreateAccessTokenFail
		srv.loger.LogErrorf("Create Access Token Failed: [%+v]", err.Error(), int(loginId))
		return
	}
	err = srv.UpdateSession(request,session)
	if err != nil {
		err = srv.em.Auth.AddNewSessionFail
		srv.loger.LogErrorf("Update Session Failed: [%+v]", err.Error(), int(loginId))
		return
	}

	srv.loger.LogInfo("Create Refresh Token ...", int(loginId))
	expirationTime = time.Now().Add(3 * time.Hour)
	rtclaims := &rtClaims {
		Session:     session,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	rftoken,err := srv.CreateToken(rtclaims)
	if err != nil {
		err = srv.em.Auth.CreateRefreshTokenFail
		srv.loger.LogErrorf("Create Refresh Token Failed: [%+v]", err.Error(), int(loginId))
		return
	}

	result = parseCode {
		AccessToken: token,
		RefreshToken: rftoken,
	}
	srv.loger.LogInfof("Access Token: [%+v]", result.AccessToken, int(loginId))
	srv.loger.LogInfof("Refresh Token: [%+v]", result.RefreshToken, int(loginId))
	srv.loger.LogInfo("Signin Successfully", int(loginId))
	return
}

func (srv *authService)CreateToken(claims tokenClaims)  (RefreshTokenString string,err error){
	//cliams ที่ส่งเข้ามา จะกำหนดว่าได้ access_token หรือ refresh_token
	log.Info("[CreateToken ...]")
	var signKey *rsa.PrivateKey

	signBytes, err := ioutil.ReadFile(PriKeyPath)
	if err != nil {
		err = srv.em.Auth.GetPrivateKeyFail
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		err = srv.em.Auth.ParsePrivateKeyFail
		return
	}

	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	RefreshTokenString, err = RefreshToken.SignedString(signKey)
	if err != nil {
		err = srv.em.Auth.StatusInternalServerError
		log.Error(err)
		// If there is an error in creating the JWT return an internal server error
		return
	}
	return
}

func (srv *authService)UpdateSession(request credentials,session string)  (err error){
	log.Info("[UpdateSession in database ...]")

	log.Info("[Update Time ...]")
	timeNow := time.Now()
	//expiration Token +1Hour
	timeToken := timeNow.Add(1 * time.Hour)

	// update lastLogin_date , session , exp_session
	update, err := srv.repo.updateLastLogin(request, timeNow, session, timeToken)
	if err != nil && err != gorm.ErrRecordNotFound {
		err = srv.em.Auth.LoginNotFound
		log.Error(err)
		return
		//return
	}
	log.Info("LastLoginDate : ", update.LastLogin)
	log.Info("Expiration Token : ", update.ExpSession)
	return
}

func (srv *authService)RefreshAccessToken(RefreshToken string) (result parseCode, err error) {
	log.Info("[refresh access token ...]")

	log.Info("[check refresh_token ...]")
	var verifyKey *rsa.PublicKey
	verifyBytes, err := ioutil.ReadFile(PubKeyPath)
	if err != nil {
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return
	}
	chclaims := &Claims{}
	tkn, err := jwt.ParseWithClaims(RefreshToken, chclaims, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if tkn == nil || !tkn.Valid {
		err = srv.em.Auth.InvalidRefreshToken
		log.Error(err)
		return
	}
	if chclaims.CharacterId != 0 {
		err = srv.em.Auth.InvalidRefreshToken
		log.Error(err)
		return
	}
	detail, err := srv.repo.checkSessionRf(chclaims.Session)
	if err != nil {
		err = srv.em.Auth.InvalidRefreshToken
		log.Error(err)
		return
	}

	log.Info("[create new access_token ...]")
	session := uuid.New().String()
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		CharacterId: detail.CharacterId,
		Session:     session,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token,err:=srv.CreateToken(claims)
	if err != nil {
		err = srv.em.Auth.CreateAccessTokenFail
		log.Error(err)
		return
	}
	updateData:= credentials{Username: detail.Username}
	err=srv.UpdateSession(updateData,session)
	if err != nil {
		err = srv.em.Auth.AddNewSessionFail
		log.Error(err)
		return
	}

	result = parseCode{
		AccessToken: token,
		RefreshToken: RefreshToken,
	}
	return
}

func CheckSession(session string, characterId int) (checkSession bool, err error) {
	//validate Session
	var detail character
	detail, err = getSession(characterId)
	if err != nil && err == gorm.ErrRecordNotFound {
		checkSession = false
		log.Error("Login not found.")
		return
	}
	log.Infof("Detail : %+v", detail)
	if detail.Session != session {
		checkSession = false
		log.Error("An unexpected error occurred on get session.")
		return
	}
	checkSession = true
	return
}

func (srv *authService) GetClaimCurrent(c *gin.Context) {
	if claims, ok := c.Get("claim"); ok {
		srv.claims = claims.(*Claims)
	}
}

func (srv *authService) GetCharacterIdInClaim() int {
	return srv.claims.CharacterId
}

func (srv *authService) GetRoleNameInClaim() string {
	return srv.claims.RoleName
}
