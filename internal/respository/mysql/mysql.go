package mysql

import (
	"fmt"
	"time"

	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Model struct {
	ID        int64 `gorm:"primary_key;auto_increment;not null;type:bigint(20)" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func DBConnect(host, database, user, password string, enableLog bool) (*gorm.DB, error) {
	var err error
	addr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=%s&allowNativePasswords=true",
		user, password, host, database, "UTC")
	var logLevel gormLogger.LogLevel
	if enableLog {
		logLevel = gormLogger.Info
	} else {
		logLevel = gormLogger.Error
	}

	db, err = gorm.Open(mysqlDriver.Open(addr), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if enableLog {
		db.Debug()
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(60)
	err = db.AutoMigrate()

	if err != nil {
		return nil, err
	}

	return db, nil
}
