package sqlClient

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var DB *gorm.DB

func init() {
	log.Debug("init db")
	DB = OpenDb()
}
func OpenDb() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "@Gx.1101434570", "127.0.0.1", "3308", "jxserver")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{ //建立连接时指定打印info级别的sql
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
	if err != nil {
		log.Error("数据库连接失败")
		log.Error(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Error(err.Error())
	}
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大打开连接数
	sqlDB.SetMaxIdleConns(100)                 // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置空闲连接最大存活时间

	return db
}
