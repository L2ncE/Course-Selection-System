package dao

import (
	"CSA/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitGormDB() (err error) {
	dB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			global.Settings.GormInfo.Name, global.Settings.GormInfo.Password, global.Settings.GormInfo.Host, global.Settings.GormInfo.Port, global.Settings.GormInfo.DBName), // DSN data source name
		DefaultStringSize:        171,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("连接失败：%v\n", err)
	}
	db = dB
	return err
}
