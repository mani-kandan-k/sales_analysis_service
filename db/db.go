package db

import (
	"database/sql"
	"fmt"
	"sales_analysis_service/logger"
	"sales_analysis_service/readtoml"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GMysqlDB *sql.DB
var GMysqlGormDB *gorm.DB

type DbConfigStruct struct {
	User     string
	Port     string
	Server   string
	Password string
	Database string
}

func GetDbConfig() (lDbConfigData DbConfigStruct) {
	logger.Info("GetDbConfig (+)")
	lDbConfig := readtoml.ReadToml("./toml/dbconfig.toml")
	lDbConfigData.User = readtoml.GetConfigValue(lDbConfig, "USER")
	lDbConfigData.Password = readtoml.GetConfigValue(lDbConfig, "PASSWORD")
	lDbConfigData.Server = readtoml.GetConfigValue(lDbConfig, "SERVER")
	lDbConfigData.Port = readtoml.GetConfigValue(lDbConfig, "PORT")
	lDbConfigData.Database = readtoml.GetConfigValue(lDbConfig, "DATABASE")
	logger.Info("GetDbConfig (-)")
	return
}

func DataBaseInit(pDbConfigData DbConfigStruct) (lErr error) {
	logger.Info("DataBaseInit (+)")
	lConnString := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", pDbConfigData.User, pDbConfigData.Password, pDbConfigData.Server, pDbConfigData.Port, pDbConfigData.Database)
	if GMysqlGormDB, lErr = gorm.Open(mysql.Open(lConnString), &gorm.Config{}); lErr != nil {
		logger.Err(lErr)
		return
	}
	if GMysqlDB, lErr = GMysqlGormDB.DB(); lErr != nil {
		logger.Err(lErr)
		return
	}
	GMysqlDB.SetMaxIdleConns(5)
	GMysqlDB.SetMaxOpenConns(5)
	GMysqlDB.SetConnMaxIdleTime(time.Duration(10) * time.Minute)
	logger.Info("DataBaseInit (-)")
	return
}

func CloseDbConnection() {
	if GMysqlDB != nil {
		GMysqlDB.Close()
	}
}
