package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "new-go-project/docs"
	_ "new-go-project/model"
	"new-go-project/router"
	v1 "new-go-project/router/api/v1"
)
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8088
// @BasePath
func main() {
	r := router.InitRouter()
	r.GET("/test",v1.FirstTest)
	r.Run(":8088")
}