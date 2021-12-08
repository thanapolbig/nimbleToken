package loger

import (
	"time"
)

const (
	InfoLevel = "info"
	DebugLevel = "debug"
	ErrorLevel = "error"
)

type logModel struct {
	SeriesId	string		`json:"series_id"`
	LoginId		int			`json:"login_id"`
	LogLevel	string		`json:"log_level"`
	Service		string		`json:"service"`
	Message		string		`json:"message"`
	DateTime	time.Time	`json:"date_time"`
}


type MessageResponse struct {
	Status             int    `json:"status"`
	MessageCode        string `json:"message_code"`
	MessageDescription string `json:"message_description"`
}
