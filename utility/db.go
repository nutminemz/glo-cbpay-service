package utility

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"gitlab.com/firstkungz/log-go"
)

var DB *gorm.DB

func Initialize() {
	connection := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", viper.GetString("cloud-sql.username"), viper.GetString("cloud-sql.password"), viper.GetString("cloud-sql.gormhost"), viper.GetString("cloud-sql.dbname"))
	log.Infof("Connecting To Db : %s", connection)
	db, err := gorm.Open("mysql", connection)
	db.DB().SetMaxIdleConns(viper.GetInt("cloud-sql.max-idle-conns"))
	db.DB().SetMaxOpenConns(viper.GetInt("cloud-sql.max-open-conns"))
	db.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("cloud-sql.max-life-time-minutes")) * time.Minute)
	if err != nil {
		log.Fatal("Init DB Error", err)
	}
	DB = db
}

func GetSqlClient() *gorm.DB {
	return DB
}

func Close() {
	DB.Close()
}
