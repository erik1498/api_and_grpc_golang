package utils

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(cfg Config, logger Zlog) *gorm.DB {
	username := cfg.GetString(DbUser)
	password := cfg.GetString(DbPassword)
	host := cfg.GetString(DbHost)
	port := cfg.GetString(DbPort)
	dbname := cfg.GetString(DbName)
	parseTime := cfg.GetString(DbParseTime)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s", username, password, host, port, dbname, parseTime)
	logger.LogInfo(nil, "%s", dsn)
	var err error
	logger.LogInfo(nil, "%s", "Db Connect")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogError(err, nil, "%s", "cannot connect to dsn")
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.LogError(err, nil, "%s", "cannot connect to DB")
	}
	configureConnectionPool(sqlDB)
	return db
}

func configureConnectionPool(dbPool *sql.DB) {
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxIdleConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)
}
