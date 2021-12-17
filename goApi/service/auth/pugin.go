package auth

import (
	"github.com/miguelmota/ethereum-development-with-go/app"
	"time"

	"github.com/jinzhu/gorm"
)

type registerRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *registerRepo) InitRegisterRepo(conf *app.Configs, em *app.ErrorMessage) *registerRepo {
	return &registerRepo{
		conf: conf,
		em:   em,
	}
}

type authRepo struct {
	conf *app.Configs
	em   *app.ErrorMessage
}

func (repo *authRepo) InitAuthRepo(conf *app.Configs, em *app.ErrorMessage) *authRepo {
	return &authRepo{
		conf: conf,
		em:   em,
	}
}

func (repo *authRepo) getCharacter(request credentials) (detail character, err error) {
	if err = app.Token.DB.Select("l.username ,l.password ,l.login_uuid ,w.id AS wallet_id ,w.login_id ,l.session ,r.role_name").
		Table("login AS l ,role AS r,wallet AS w ").
		Where("l.username = ? AND l.id = w.login_id AND l.role_id = r.id", request.Username).
		Find(&detail).Error; err != nil {
		return
	}
	return
}

func (repo *authRepo) updateLastLogin(request credentials, timeNow time.Time, session string, timeToken time.Time) (update roleLogin, err error) {
	if err = app.Token.DB.Table("login").
		Where("username = ?", request.Username).
		Update(map[string]interface{}{
			"last_login":  timeNow,
			"session":     session,
			"exp_session": timeToken,
		}).
		Error; err != nil {
		return
	}
	update.LastLogin = timeNow
	update.ExpSession = timeToken
	return update, nil
}

func getSession(request int) (detail character, err error) {

	if err = app.Token.DB.Select("l.username ,l.password ,l.login_uuid ,w.id AS wallet_id ,w.login_id ,l.session").
		Table("wallet AS w ,login AS l").
		Where("w.id = ? AND w.login_id = l.id ", request).
		Find(&detail).Error; err != nil {
		return
	}

	return
}

func (repo *registerRepo) checkEmail(request inputRegister) (result user, err error) {
	if err = app.Token.DB.
		Table("NimbleToken.dbo.login").
		Where("email = ?", request.Email).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) checkUsername(request inputRegister) (result user, err error) {
	if err = app.Token.DB.
		Table("NimbleToken.dbo.login").
		Where("username = ?", request.Username).
		Find(&result).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) createUser(tx *gorm.DB, request user) (err error) {
	if tx == nil {
		tx = app.Token.DB
	}
	if err = tx.Table("NimbleToken.dbo.login").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func (repo *registerRepo) getUserId(tx *gorm.DB, request string) (user user, err error) {
	if tx == nil {
		tx = app.Token.DB
	}
	if err = tx.Table("NimbleToken.dbo.login").
		Where("username = ?", request).
		Find(&user).Error; err != nil {
		return
	}
	return
}


func (repo *authRepo) checkSessionRf(request string) (result useData, err error) {
	if err = app.Token.DB.Select("l.username ,c.id AS character_id").
		Table("NimbleToken.dbo.character AS c ,NimbleToken.dbo.login AS l").
		Where("l.session = ? AND c.login_id = l.id", request).
		Find(&result).Error; err != nil {
		return
	}
	return
}


