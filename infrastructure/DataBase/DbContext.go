package DataBase

import (
	"StudentRecord/domain/aggregates/userAggregates"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const connStr = "root:Yxy139@tcp(49.232.204.239:3306)/study_record?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	go func() {
		db := NewDb()
		// 执行数据库迁移
		err := db.AutoMigrate(
			&userAggregates.User{},
			&userAggregates.Account{})
		if err != nil {
			panic("数据库迁移失败:error:" + err.Error())
		}
		fmt.Println("数据库迁移成功")
	}()

}
func NewDb() *gorm.DB {
	// 连接字符串
	db, err := gorm.Open(mysql.Open(connStr))
	if err != nil {
		panic("数据库连接失败:error:" + err.Error())
	}
	return db
}
