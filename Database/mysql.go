package database

import (
	"fmt"
	"ginDemoProject/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author: kylo_cheok
 * @Email:  maggic0816@gmail.com
 * @Date:   2022/5/14 13:12
 * @Desc:   Grace under pressure
 */

const (
	USERNAME = "root"
	PASSWORD = "zhUoyue816"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "gin_db_mysql"
	CHARSET  = "utf8mb4"
)

func connectMysql() (p *gorm.DB, e error) {
	//dsn := "root:zhUoyue816@tcp(127.0.0.1:3306)/gin_db_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}
	//defer db.Close()
	return db, nil
}

func QueryMysql(value string) Models.User {
	db, _ := connectMysql()
	var user Models.User
	db.Raw("select * from user where username = ?", value).Scan(&user)
	return user
}

func UpdateMysql(key string, value string) {
	db, _ := connectMysql()
	db.Raw("update user set ? where username = ?", key, value)
}

func CreateMysql(name string, password string) error {
	db, _ := connectMysql()
	user := Models.User{Username: name, Password: password}
	db.Table("user").Select("username", "password").Create(&user)

	if err := db.Table("user").Select("username", "password").Create(&user).Error; err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}

//var db *gorm.DB
//
//type Model struct {
//	ID         int `gorm:"primary_key" json:"id"`
//	CreatedOn  int `json:"created_on"`
//	ModifiedOn int `json:"modified_on"`
//}
//
//func init() {
//	var (
//		err                                             error
//		DBNAME, USERNAME, PASSWORD, HOST, PORT, CHARSET string
//	)
//
//	sec, err := setting.Cfg.GetSection("database")
//	if err != nil {
//		log.Fatal(2, "Fail to get section 'database': %v", err)
//	}
//
//	DBNAME = sec.Key("DATABASE").String()
//	USERNAME = sec.Key("USERNAME").String()
//	PASSWORD = sec.Key("PASSWORD").String()
//	HOST = sec.Key("HOST").String()
//	PORT = sec.Key("PORT").String()
//	CHARSET = sec.Key("CHARSET").String()
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DBNAME, CHARSET)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Println(err)
//	}
//
//	db.AutoMigrate(&todoModel{})
//}
