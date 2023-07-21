package main

import (
	"encoding/json"
	"fmt"
	"github.com/W1llyu/ourjson"
	"github.com/bitly/go-simplejson"
	"reflect"
	"strconv"
	"testing"
)

func TestOurJson(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	jsonStr := `{
        "user": {
            "name": "aa",
            "age": 10,
            "phone": "12222222222",
            "emails": [
                "aa@164.com",
                "aa@165.com"
            ],
            "address": [
                {
                    "number": "101",
                    "now_live": true
                },
                {
                    "number": "102",
                    "now_live": null
                }
            ],
            "account": {
                "balance": 999.9
            }
        }
    }
    `
	jsonObject, err := ourjson.ParseObject(jsonStr)
	fmt.Println(jsonObject, err)

	user := jsonObject.GetJsonObject("user")
	fmt.Println(user)

	name, err := user.GetString("name")
	fmt.Println(name, err)

	phone, err := user.GetInt64("phone")
	fmt.Println(phone, err)

	age, err := user.GetInt64("age")
	fmt.Println(age, err)

	account := user.GetJsonObject("account")
	fmt.Println(account)

	balance, err := account.GetFloat64("balance")
	fmt.Println(balance, err)

	email1, err := user.GetJsonArray("emails").GetString(0)
	fmt.Println(email1, err)

	address := user.GetJsonArray("address")
	fmt.Println(address)

	address1nowLive, err := user.GetJsonArray("address").GetJsonObject(0).GetBoolean("now_live")
	fmt.Println(address1nowLive, err)

	address2, err := address.Get(1)
	fmt.Println(address2, err)

	address2NowLive, err := address2.JsonObject().GetNullBoolean("now_live")
	fmt.Println(address2NowLive, err)

	jsonObject, err = ourjson.ParseObject(jsonStr)
	marshal, err := json.Marshal(*jsonObject)
	if err != nil {
		panic(err)
	}
	fmt.Println("---------", marshal)
}

var json_str string = `
{"rc" : 0,
  "error" : "Success",
  "type" : "stats",
  "progress" : 100,
  "job_status" : "COMPLETED",
  "result" : {
    "total_hits" : 803254,
    "starttime" : 1528434707000,
    "endtime" : 1528434767000,
    "fields" : [ ],
    "timeline" : {
      "interval" : 1000,
      "start_ts" : 1528434707000,
      "end_ts" : 1528434767000,
      "rows" : [ {
        "start_ts" : 1528434707000,
        "end_ts" : 1528434708000,
        "number" : "x12887"
      }, {
        "start_ts" : 1528434720000,
        "end_ts" : 1528434721000,
        "number" : "x13028"
      }, {
        "start_ts" : 1528434721000,
        "end_ts" : 1528434722000,
        "number" : "x12975"
      }, {
        "start_ts" : 1528434722000,
        "end_ts" : 1528434723000,
        "number" : "x12879"
      }, {
        "start_ts" : 1528434723000,
        "end_ts" : 1528434724000,
        "number" : "x13989"
      } ],
      "total" : 803254
    },
      "total" : 8
  }
}`

func TestSimpleJson(t *testing.T) {

	res, err := simplejson.NewJson([]byte(json_str))

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	//获取json字符串中的 result 下的 timeline 下的 rows 数组
	rows, err := res.Get("result").Get("timeline").Get("rows").Array()

	//遍历rows数组
	for _, row := range rows {
		//对每个row获取其类型，每个row相当于 C++/Golang 中的map、Python中的dict
		//每个row对应一个map，该map类型为map[string]interface{}，也即key为string类型，value是interface{}类型
		if each_map, ok := row.(map[string]interface{}); ok {

			//可以看到each_map["start_ts"]类型是json.Number
			//而json.Number是golang自带json库中decode.go文件中定义的: type Number string
			//因此json.Number实际上是个string类型
			fmt.Println("reflect.TypeOf(each_map[start_ts])=", reflect.TypeOf(each_map["start_ts"]))

			if start_ts, ok := each_map["start_ts"].(json.Number); ok {
				start_ts_int, err := strconv.ParseInt(string(start_ts), 10, 0)
				if err == nil {
					fmt.Println(start_ts_int)
				}
			}

			if number, ok := each_map["number"].(string); ok {
				fmt.Println(number)
			}

		}
	}
}
