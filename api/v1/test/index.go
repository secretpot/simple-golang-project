package test

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	g := r.Group("/test")
	{
		g.GET("/hello", hello)
		g.GET("/ping", ping)
	}
}
