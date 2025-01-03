package text

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	g := r.Group("/text")
	{
		g.GET("/echo", echo)
	}
}
