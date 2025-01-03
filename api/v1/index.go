package v1

import (
	"github.com/secretpot/simple-golang-project/api/v1/test"
	"github.com/secretpot/simple-golang-project/api/v1/text"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	test.Register(r)
	text.Register(r)
}
