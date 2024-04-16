package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Testapi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "This is a test API")
}