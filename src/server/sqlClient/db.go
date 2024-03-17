package sqlClient

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path"
	"server/tool"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var DB *gorm.DB

func init() {
	tool.Debug("init db" + "127.0.0.1:" + "3308")
	DB = OpenDb()
}
func OpenDb() *gorm.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "@Gx.1101434570", "127.0.0.1", "3308", "jxserver")
	now := time.Now()

	filename := fmt.Sprintf("%d%02d%02d.log",
		now.Year(),
		now.Month(),
		now.Day(),
	)

	logfile, _ := os.OpenFile(path.Join("./logs", filename), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	newLogger := logger.New(
		log.New(logfile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
		},
	)
	var isDebug = false
	args := os.Args[1:] // 去除第一个参数，即程序本身的路径
	for i, arg := range args {
		fmt.Println("参数", i, "是", arg)
		if arg == "debug" {
			isDebug = true
		}
	}
	if isDebug {
		newLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{ //建立连接时指定打印info级别的sql
		Logger: newLogger, //配置日志级别，打印出所有的sql
	})
	if err != nil {
		tool.Error("数据库连接失败")
		tool.Error(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		tool.Error(err.Error())
	}
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大打开连接数
	sqlDB.SetMaxIdleConns(100)                 // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 设置空闲连接最大存活时间

	return db
}
