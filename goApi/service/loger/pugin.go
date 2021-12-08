package loger

import (
	"github.com/miguelmota/ethereum-development-with-go/app"
)

func createLog(request logModel) (err error) {

	if err = app.Token.DB.
		Table("log_application").
		Save(&request).Error; err != nil {
		return
	}

	return
}
