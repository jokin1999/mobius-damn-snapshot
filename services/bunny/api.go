package bunny

import (
	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/crust"
)

// register api v1
func apiv1(r *gin.Engine) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})

	api := r.Group("/api/v1")
	{
		// register
		api.POST("/reg", crust.Reguser)
		api.POST("/del", crust.Deluser)
		api.POST("/login", crust.Login)
	}
}
