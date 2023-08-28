package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
)

// https://gorm.io/zh_CN/docs/create.html

type User struct {
	Id   int
	Name string
	Age  int
}

func (User) TableName() string {
	return "user"
}

func InitDb() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	if err != nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.NamingStrategy = schema.NamingStrategy{
		TablePrefix:         "",
		SingularTable:       true,
		NameReplacer:        nil,
		NoLowerCase:         false,
		IdentifierMaxLength: 0,
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	return db
}

func TestConn(t *testing.T) {
	db := InitDb()
	user := &User{
		Name: "jiayun2",
		Age:  13,
	}
	result := db.Create(&user)
	fmt.Printf("user.id=%d\n", user.Id)
	fmt.Printf("result.RowsAffected=%d\n", result.RowsAffected)
}

func TestTransitional(t *testing.T) {
	db := InitDb()
	user := &User{
		Name: "jiayun2",
		Age:  13,
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func TestInsertIgnore(t *testing.T) {
	db := InitDb()
	user := &User{
		Name: "jiayun2",
		Age:  13,
	}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	fmt.Printf("user.id=%d\n", user.Id)
	fmt.Printf("result.RowsAffected=%d\n", result.RowsAffected)

	// Do nothing on conflict
	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error; err != nil {
		fmt.Println(err.Error())
	}
}
