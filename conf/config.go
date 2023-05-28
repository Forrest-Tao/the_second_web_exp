package conf

import (
	"experimen_2/modle"
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误,亲检查文件的路径")
		return
	}
	LoadServer(file)
	LoadMysql(file)
	//dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True"
	path := DbUser + ":" + DbPassWord +
		"@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&parseTime=true"
	fmt.Println(path)
	/*	= strings.Join([]string{DbUser, ":", DbPassWord,
		"@tcp(", DbHost, DbHost, ":", ")/", DbName})
	*/
	//连接数据库
	modle.Database(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
