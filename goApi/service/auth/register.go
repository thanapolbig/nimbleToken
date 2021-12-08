package auth

import (
	"github.com/miguelmota/ethereum-development-with-go/app"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/google/uuid"

	"github.com/miguelmota/ethereum-development-with-go/service/loger"
)

var rs *registerService

type registerService struct {
	conf  *app.Configs
	em    *app.ErrorMessage
	repo  *registerRepo
	loger *loger.Loger
}

func NewRegisterBindingLog(conf *app.Configs, em *app.ErrorMessage, loger *loger.Loger) *registerService {
	repo := registerRepo{}
	return &registerService{
		conf:  conf,
		em:    em,
		repo:  repo.InitRegisterRepo(conf, em),
		loger: loger,
	}
}

func (rs *registerService) RegisterTransaction(request inputRegister) (msg MessageResponse, err error) {
	rs.loger.LogInfo("Register Validation ...", 0)

	rs.loger.LogInfo("Trim Space Target ...", 0)
	request.Username = strings.TrimSpace(request.Username)

	rs.loger.LogInfo("Empty Target Validation ...", 0)
	//validate username empty
	if request.Username == "" {
		err = rs.em.Register.ValidateFail.ValidateUsernameEmpty
		rs.loger.LogErrorf("Empty Target Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Target Length Validation ...", 0)
	//validate username max 15
	if len(request.Username) > 15 {
		err = rs.em.Register.ValidateFail.ValidateUsernameLength
		rs.loger.LogErrorf("Target Length Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Password Length Validation ...", 0)
	//validate password max 20
	if len(request.Password) > 20 || len(request.RepeatedPassword) > 20 {
		err = rs.em.Register.ValidateFail.ValidatePasswordLength
		rs.loger.LogErrorf("Password Length Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Email Length Validation ...", 0)
	//validate email max 30
	if len(request.Email) > 30 {
		err = rs.em.Register.ValidateFail.ValidateEmailLength
		rs.loger.LogErrorf("Email Length Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Password Match Validation ...", 0)
	//validate match password
	if request.Password != request.RepeatedPassword {
		err = rs.em.Register.ValidateFail.PasswordNotMatch
		rs.loger.LogErrorf("Password Match Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Duplicate Email Validation ...", 0)
	//validate duplicate email
	var userFound user
	userFound, err = rs.repo.checkEmail(request)
	if (userFound != user{}) || (err != nil && err != gorm.ErrRecordNotFound) {
		err = rs.em.Register.ValidateFail.DuplicateEmail
		rs.loger.LogErrorf("Duplicate Email Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	rs.loger.LogInfo("Duplicate Target Validation ...", 0)
	//validate duplicate username
	userFound, err = rs.repo.checkUsername(request)
	if (userFound != user{}) || (err != nil && err != gorm.ErrRecordNotFound) {
		err = rs.em.Register.ValidateFail.DuplicateUsername
		rs.loger.LogErrorf("Duplicate Target Validation Failed: [%+v]", err.Error(), 0)
		return
	}

	//create transaction -> user, character, inventory, farm
	if err == gorm.ErrRecordNotFound {
		rs.loger.LogInfo("Transaction Started ...", 0)
		tx := app.Token.DB.Begin()
		err = nil
		loginUuid := uuid.New().String()
		password := encryptPassword(compilePassword(request.Password, loginUuid))
		var user = user{
			LoginUuid: loginUuid,
			Username:  request.Username,
			Password:  password,
			Email:     request.Email,
			RoleId:    2,
		}
		varTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		user.RegisterDate = varTime
		user.CreateDate = varTime
		user.UpdateDate = varTime

		rs.loger.LogInfo("Create User ...", 0)
		//create user
		err = rs.repo.createUser(tx, user)
		if err != nil {
			err = rs.em.Register.CreateFail.CreateUsername
			rs.loger.LogErrorf("Create User Failed: [%+v]", err.Error(), 0)
			tx.Rollback()
			return
		}

		rs.loger.LogInfo("Get User Id ...", 0)
		//get user_id for character_id
		user, err = rs.repo.getUserId(tx, string(request.Username))
		if err != nil {
			err = rs.em.Register.CreateFail.CreateUsername
			rs.loger.LogErrorf("Get User Id Failed: [%+v]", err.Error(), user.Id)
			tx.Rollback()
			return
		}

		rs.loger.LogInfo("Register Successfully ...", user.Id)
		tx.Commit()
	}
	return
}
