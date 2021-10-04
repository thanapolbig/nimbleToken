package service

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

type Input struct {
	Key 		string		`json:"address"`
	Value		float64		`json:"value"`
}

type InputBalanceOf struct {
	Address string `json:"address"`
}

type InputMint struct {
	AddressHeader	string	`json:"address_header"`
	Address		string		`json:"address"`
	Value		string	`json:"value"`
}

type InputAppove struct {
	AddressSpender string `json:"addressSpender"`
	PrivateKey     string `json:"privateKey"`
	Value          int    `json:"value"`
}

type InputClaimReward struct {
	Sender     string `json:"Sender"`
	PrivateKey string `json:"privateKey"`
	Value      string    `json:"value"`
}

type InputAllowance struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
}

type InputBurn struct {
	PrivateKey string `json:"privatekey"`
	Value      string `json:"value"`
}

type InputKeyValue struct {
	PrivateKey string `json:"privatekey"`
	Value      string `json:"value"`
}

type InputKeyArray struct {
	PrivateKey string 		`json:"privatekey"`
	DataList	[]string	`json:"data_list"`
}
type Vote struct {
	PrivateKey string `json:"privatekey"`
	ToAddress  string `json:"toAddress"`
	Value      string `json:"value"`
}

type TokenSession struct {
	Contract     interface{}
	CallOpts     *bind.CallOpts
	TransactOpts *bind.TransactOpts
}


