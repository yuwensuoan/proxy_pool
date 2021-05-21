package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proxy_pool/app/services"
)

type ProxyController struct {
	BaseController
}

/**
获取代理列表
 */
func (ProxyController) GetList (ctx *gin.Context) {
	ctx.String(http.StatusOK, "list")
}

/**
获取单条代理
 */
func (ProxyController) GetFirst (ctx *gin.Context) {
	services.ProxyService{}.GetList()
	ctx.String(http.StatusOK, "first")
}

func (ProxyController) Delete (ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete")
}