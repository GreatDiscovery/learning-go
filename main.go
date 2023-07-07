package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"learning-go/docs"
	"net/http"
	"strconv"
)

// @title 开发文档
// @version 0.0.1
// @BasePath /api/v1/

// @title  haimait.com开发文档
// @version 1.0
// @description  Golang api of demo
// @termsOfService haimait.com

// @contact.name API Support
// @contact.url haimait.com
// @contact.email ×××@qq.com
// @BasePath /api/v1/
func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.New()
	swagHandler := true
	if swagHandler {
		// 文档界面访问URL
		// http://127.0.0.1:8080/swagger/index.html
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 创建路由组
	v1 := r.Group("/api/v1")

	v1.GET("/getUser/:id", getUser)

	r.Run()
}

// @Tags 测试
// @Summary  获取指定getUser记录1
// @Description 获取指定getUser记录2
// @Accept  json
// @Product json
// @Param   id     query    int     true        "用户id"
// @Param   name   query    string  false        "用户name"
// @Success 200 {object} string	"{"code": 200, "data": [...]}"
// @Router /getUser/:id [get]
func getUser(c *gin.Context) {
	var r req
	Id := c.DefaultQuery("id", "0")
	r.Id, _ = strconv.Atoi(Id)
	r.Name = c.DefaultQuery("name", "")
	Age, _ := strconv.Atoi(c.DefaultQuery("age", "0"))
	r.Age = Age
	fmt.Println(r)
	c.JSON(http.StatusOK, r)
}

type req struct {
	Id   int    `json:"id" form:"id" example:"1"`
	Name string `json:"name" form:"name" example:"用户name"`
	Age  int    `json:"age" form:"age" example:"123"`
}
