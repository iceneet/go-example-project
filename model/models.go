package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"new-go-project/pkg/setting"
)

var db * gorm.DB
//var db1 * gorm.DB
//基本模型
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

//初始化db
func init()  {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	//单数表开启
	db.SingularTable(true)
	//日志开启
	db.LogMode(true)
	fmt.Println(db)
	//SetDB(db1)
}
//DB获取
func GetDB() * gorm.DB {
	return db
}
func CloseDB() {
	defer db.Close()
}

//设置DB
func SetDB(db *gorm.DB)  {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	//单数表开启
	db.SingularTable(true)
	//日志开启
	db.LogMode(true)
}

