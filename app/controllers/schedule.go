package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Schedule struct {

}

func (Schedule) Stop(ctx *gin.Context) {
	ctx.String(http.StatusOK, "success!")
}