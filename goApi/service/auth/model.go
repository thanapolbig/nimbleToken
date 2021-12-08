package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	PriKeyPath = "certification/key.auth/bootcamp.key"
	PubKeyPath = "certification/key.auth/bootcamp.key.pub"
)

type roleLogin struct {
	LoginUuid    string    `json:"login_uuid"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	LastLogin    time.Time `json:"last_login"`
	Session      string    `json:"session"`
	ExpSession   time.Time `json:"exp_session"`
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type character struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	LoginUuid   string `json:"login_uuid"`
	CharacterId int    `json:"character_id"`
	LoginId     string `json:"login_id"`
	Session     string `json:"session"`
	RoleName    string `json:"role_name"`
}

type Claims struct {
	CharacterId int    `json:"character_id"`
	Session     string `json:"session"`
	RoleName    string `json:"role_name"`
	jwt.StandardClaims
}

type rtClaims struct {
	Session string `json:"session"`
	jwt.StandardClaims
}

type tokenClaims interface {
	Valid() error
}

type useData struct {
	CharacterId int    `json:"character_id"`
	Username    string `json:"username"`
}

type DecodeClaims struct {
	CharacterId int    `json:"character_id"`
	Session     string `json:"session"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}
type parseCode struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type MessageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}

type inputRePassword struct {
	UserName           string `json:"user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
	Otp                string `json:"otp"`
}

type resetPassword struct {
	UserName           string `json:"user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
}

type inputFirstLogin struct {
	NewUserName        string `json:"new_user_name"`
	NewPassWord        string `json:"new_pass_word"`
	ConfirmNewPassWord string `json:"confirm_new_pass_word"`
	CitizenId          string `json:"citizen_id"`
}

type inputReqOTP struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type User struct {
	Id           int       `json:"id"`
	LoginUuid    string    `json:"login_uuid"`
	UserName     string    `json:"user_name"`
	PassWord     string    `json:"pass_word"`
	Name         string    `json:"name"`
	CitizenId    string    `json:"citizen_id"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
	CompanyId    int       `json:"company_id"`
	Email        string    `json:"email"`
	OTP          string    `json:"otp"`
	IsFirstLogin bool      `json:"is_first_login"`
}

type inputRegister struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
	Email            string `json:"email"`
	Gender           int    `json:"gender"`
	SkinId           int    `json:"skin_id"`
	HatId            int    `json:"hat_id"`
	ShirtId          int    `json:"shirt_id"`
	ShoesId          int    `json:"shoes_id"`
}

type user struct {
	Id           int       `json:"id"`
	LoginUuid    string    `json:"login_uuid"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	LastLogin    time.Time `json:"last_login"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
	RoleId       int       `json:"role_id"`
}


type rolePermit struct {
	Id              int    `json:"id"`
	PermitCharacter int    `json:"permit_character"`
	PermitFarm      int    `json:"permit_farm"`
	PermitMarket    int    `json:"permit_market"`
	PermitLottery   int    `json:"permit_lottery"`
	PermitQuest     int    `json:"permit_quest"`
	PermitGachapon  int    `json:"permit_gachapon"`
	PermitGiftcode  int    `json:"permit_giftcode"`
	Description     string `json:"description"`
	PermitSkill     int    `json:"permit_skill"`
}

type DetailCharacterSkill struct {
	Id          int       `json:"id"`
	CharacterId int       `json:"character_id"`
	SkillId     int       `json:"skill_id"`
	Remaining   int       `json:"remaining"`
	UpdateDate  time.Time `json:"update_date"`
	RefreshDate time.Time `json:"refresh_date"`
}