package main

import (
	v1 "github.com/secretpot/simple-golang-project/api/v1"

	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

func main() {
	v1.Register(engine.Group("/api/v1"))

	engine.Run(":8000")
}
