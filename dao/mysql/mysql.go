package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 定义一个初始化数据库的函数
var db *sqlx.DB

type MySQLConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	DBname        string
	MaxOpenConnes int
	MaxIdleConnes int
}

func Init() (err error) {
	var mysqlConfig = MySQLConfig{
		User:          viper.GetString("mysql.user"),
		Password:      viper.GetString("mysql.password"),
		Host:          viper.GetString("mysql.host"),
		Port:          viper.GetInt("mysql.port"),
		DBname:        viper.GetString("mysql.dbname"),
		MaxOpenConnes: viper.GetInt("mysql.max_open_conns"),
		MaxIdleConnes: viper.GetInt("mysql.max_idle_conns"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBname,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(mysqlConfig.MaxOpenConnes)
	db.SetMaxIdleConns(mysqlConfig.MaxIdleConnes)
	return
}

func Close() {
	_ = db.Close()
}
