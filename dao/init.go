package dao

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var _db *gorm.DB

func Database(dsn string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                   // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                  // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Second * 300) // 设置连接可复用的最大时间

	_db = db
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

func NewTransactionDBClient(ctx context.Context) *gorm.DB {
	db := _db.Begin()
	return db.WithContext(ctx)
}
