package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册web路由
func RegisterWebRouter(root *gin.Engine)  {
	webRouter := root.Group("/")
	{
		webRouter.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "")
		})
	}
}