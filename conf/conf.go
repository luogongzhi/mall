package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"mall/dao"
	"strings"
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	dao.Database(loadMysqlData(file))
}

func loadMysqlData(file *ini.File) string {
	dbHost := file.Section("mysql").Key("DbHost").String()
	dbPort := file.Section("mysql").Key("DbPort").String()
	dbUser := file.Section("mysql").Key("DbUser").String()
	dbPassWord := file.Section("mysql").Key("DbPassWord").String()
	dbName := file.Section("mysql").Key("DbName").String()
	return strings.Join([]string{dbUser, ":", dbPassWord, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
}
