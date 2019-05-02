package module

import (
	"github.com/aloxc/gapporder/config"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func init() {
	dbName := os.Getenv(config.ORDER_MYSQL_DATABASE_NAME)
	host := os.Getenv(config.ORDER_MYSQL_HOST)
	user := os.Getenv(config.ORDER_MYSQL_USER)
	password := os.Getenv(config.ORDER_MYSQL_PASSWORD)
	if dbName == "" {
		dbName = config.ORDER_MYSQL_DATABASE_NAME_DEFAULT
	}
	if host == "" {
		host = config.ORDER_MYSQL_HOST_DEFAULT
	}
	if user == "" {
		user = config.ORDER_MYSQL_USER_DEFAULT
	}
	if password == "" {
		password = config.ORDER_MYSQL_PASSWORD_DEFAULT
	}

	logs.Info("准备初始化数据库表")
	ds := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&loc=Local"
	logs.Info(ds)
	// set default database
	orm.RegisterDataBase("default", "mysql", ds, 30)
	// register model
	orm.RegisterModel(new(Order))
	// create table
	orm.RunSyncdb("default", true, true)
}
