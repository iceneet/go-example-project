package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-go-project/model"
)

var db = model.GetDB()

// @Summary   测试swagger
// @Tags test
// @Produce  json
// @Success 200 {string} model.Result "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /test  [GET]
func  FirstTest(ctx *gin.Context){
	var menu model.Menu
	db.Last(&menu)
	result := model.ResultInit(200,"测试成功",menu)
	ctx.JSON(http.StatusOK, result)
}
