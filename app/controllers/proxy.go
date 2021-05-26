package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proxy_pool/app/services"
)

// proxy 控制器
type ProxyController struct {
	BaseController
	Service services.ProxyService
}

// 获取代理列表
func (P ProxyController) GetList (ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "10")
	ctx.JSON(http.StatusOK, P.Service.GetList(page, pageSize))
}

// 获取1条代理
func (P ProxyController) GetFirst (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, P.Service.GetFirst())
}

// 删除一条代理
func (ProxyController) Delete (ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete")
}