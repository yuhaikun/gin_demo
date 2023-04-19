package dao

import (
	"bubble/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//func InitMySQL() (err error) {
//	DB, err = gorm.Open("mysql", "root:8871527yhk@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local")
//	if err != nil {
//		return
//	}
//	return DB.DB().Ping()
//}

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
