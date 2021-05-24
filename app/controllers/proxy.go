package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proxy_pool/app/services"
	"proxy_pool/app/utils"
)

// proxy 控制器
type ProxyController struct {
	BaseController
	Service services.ProxyService
}

// 获取代理列表
func (P ProxyController) GetList (ctx *gin.Context) {
	data := P.Service.GetList()

	ctx.JSON(http.StatusOK, utils.Response{}.Success(data))
}

// 获取1条代理
func (ProxyController) GetFirst (ctx *gin.Context) {
	data := services.ProxyService{}.GetFirst()
	ctx.JSON(http.StatusOK, utils.Response{}.Success(data))
}

// 删除一条代理
func (ProxyController) Delete (ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete")
}