package bunny

import (
	"github.com/gin-gonic/gin"
)

// register api v1
func apiv1(r *gin.Engine) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})

	// api := r.Group("/api/v1")
	{
		// login
	}
}
