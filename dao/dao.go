package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dB *sql.DB
var rdb *redis.Client
var db *gorm.DB

func InitDB() (err error) {
	dsn := "lance:yxh030714@tcp(1.14.43.76:3306)/douban?charset=utf8&parseTime=True"
	// 连接数据库
	dB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = dB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "1.14.43.76:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func InitGormDB() (err error) {
	dB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      "lance:yxh030714@tcp(1.14.43.76:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
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
