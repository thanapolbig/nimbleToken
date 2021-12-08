package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	log "github.com/sirupsen/logrus"
)

//var DB *gorm.DB

type Store struct {
	DB *gorm.DB
}
var Token Store

func InitDB(c *Configs) {
	var err error

	log.Infoln(c.Mssql.ConnectionMasterDB)

	db, err := gorm.Open("mssql", c.Mssql.ConnectionMasterDB)

	if err != nil {
		log.Errorln("Failed to connect database : ", c.Mssql.ConnectionMasterDB)
		log.Errorln("Error : ", err)
		panic("Failed to connect database MS SQL")
	}

	Token.DB = db
}


//func InitDB(configPath string) *gorm.DB {
//	var err error
//	// path to look for the config file in
//	viper.AddConfigPath(configPath)
//	if err = viper.ReadInConfig();
//		err != nil {
//		log.Errorln("Fatal Error Config File: ", err)
//		panic("Fatal Error Config File")
//	}
//	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
//		viper.GetString("mssql.server"),
//		viper.GetString("mssql.username"),
//		viper.GetString("mssql.password"),
//		viper.GetString("mssql.port"),
//		viper.GetString("mssql.database"))
//	log.Infoln(connectionString)
//	db, err := gorm.Open(viper.GetString("mssql.databaseType"), connectionString)
//
//	if err != nil {
//		log.Errorln("Failed to connect database : ", connectionString)
//		log.Errorln("Error : ", err)
//		panic("Failed to connect database MS SQL")
//	}
//
//	DB = db
//	return DB
//}
