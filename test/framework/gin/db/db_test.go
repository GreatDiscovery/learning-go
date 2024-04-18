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

func TestQueryAllRecord(t *testing.T) {
	db := InitDb()
	var results []User
	var data []User

	batchSize := 2 // 每次拉取的批次大小
	offset := 0
	for {
		// 查询数据
		db.Limit(batchSize).Offset(offset).Find(&data)

		// 处理数据
		for _, item := range data {
			// 处理单条数据
			results = append(results, item)
		}

		// 如果当前批次的数据不足 batchSize 个，说明已经没有更多数据了
		if len(data) < batchSize {
			break
		}
		// 增加 offset，准备下一批次查询
		offset += batchSize
	}
	fmt.Printf("len(user)=%v", len(results))
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
