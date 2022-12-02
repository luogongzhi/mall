package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"mall/model"
	"strings"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	driverName, dataSourceName := loadMysqlData(file)
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	err = engine.Sync(new(model.Cart), new(model.MemberAddress), new(model.Member), new(model.Order), new(model.Product), new(model.ProductAttr))
	if err != nil {
		fmt.Println("表初始失败")
	}

	engine.ShowSQL(true)
	db = engine
}

func loadMysqlData(file *ini.File) (string, string) {
	driverName := file.Section("mysql").Key("DriverName").String()
	dbHost := file.Section("mysql").Key("DbHost").String()
	dbPort := file.Section("mysql").Key("DbPort").String()
	dbUser := file.Section("mysql").Key("DbUser").String()
	dbPassWord := file.Section("mysql").Key("DbPassWord").String()
	dbName := file.Section("mysql").Key("DbName").String()
	dataSourceName := strings.Join([]string{dbUser, ":", dbPassWord, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	return driverName, dataSourceName
}

func NewDBClient() *xorm.Engine {
	return db
}