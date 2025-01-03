package text

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func echo(ctx *gin.Context) {
	message := ctx.Query("message")

	if len(message) > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "echo: " + message,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no message",
		})
	}
}
