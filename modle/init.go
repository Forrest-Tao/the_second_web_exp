package modle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Database(constString string) {

	db, err := gorm.Open("mysql", constString)
	if err != nil {
		fmt.Println(err)
		panic("数据库连接错误")
		return
	}
	fmt.Println("数据库连接成功")
	db.LogMode(true)             //打印数据库日志
	if gin.Mode() == "release" { //发行版
		db.LogMode(false)
	}
	db.SingularTable(true)      //表名后不加s
	db.DB().SetMaxIdleConns(20) //设置连接池
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migrate()
}
