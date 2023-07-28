package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

// 测试mysql的column中有json类型数据，gorm如何处理

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}

func TestJsonChildArray(t *testing.T) {
	DB := InitDb()
	DB.AutoMigrate(&UserWithJSON{})
	user := UserWithJSON{
		Name:       "json-1",
		Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
	}
	result := DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	// 1. select this insert row
	var tmpUser UserWithJSON
	result = DB.Model(&user).Where("id = ?", user.ID).Find(&tmpUser)
	if result.Error != nil {
		panic(result.Error)
	}

	var attributesMap map[string]interface{}
	marshalJSON, err := tmpUser.Attributes.MarshalJSON()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(marshalJSON, &attributesMap)
	if err != nil {
		panic(err)
	}

	var arr []string
	fmt.Printf("map=%v\n", attributesMap)
	if v, ok := attributesMap["tags"]; ok {
		fmt.Printf("tags=%v\n", v)
		fmt.Printf("tags.type=%v\n", reflect.TypeOf(v))
		for _, i2 := range v.([]interface{}) {
			arr = append(arr, i2.(string))
		}
		arr = append(arr, "tag7")
	}

	DB.Model(&user).Where("id = ?", user.ID).UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("tags", arr))

	var arr2 []string
	tag8 := "tag8"
	if v, ok := attributesMap["others"]; ok {
		exist := false
		for _, i2 := range v.([]interface{}) {
			arr2 = append(arr2, v.(string))
			if tag8 == i2.(string) {
				exist = true
			}
		}
		if !exist {
			arr2 = append(arr2, tag8)
		}
	} else {
		arr2 = append(arr2, tag8)
	}
	DB.Model(&user).Where("id = ?", user.ID).UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("others", arr2))
}

func TestJsonType(t *testing.T) {
	DB := InitDb()
	DB.AutoMigrate(&UserWithJSON{})
	DB.Create(&UserWithJSON{
		Name:       "json-1",
		Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
	})

	// Check JSON has keys
	user := UserWithJSON{}
	DB.Find(&user, datatypes.JSONQuery("attributes").HasKey("role"))
	fmt.Printf("user=%v\n", user)
	DB.Find(&user, datatypes.JSONQuery("attributes").HasKey("orgs", "orga"))
	fmt.Printf("user=%v\n", user)
	// MySQL
	// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.role') IS NOT NULL
	// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.orgs.orga') IS NOT NULL
	result := DB.Delete(&UserWithJSON{}, 1)
	if result.Error != nil {
		panic(result.Error)
	}
}
