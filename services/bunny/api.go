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
		// admin
		api.POST("/reg", mw_admin, crust.Reguser)
		api.POST("/del", mw_admin, crust.Deluser)
		api.POST("/perm/add", mw_admin, crust.Permadd)
		api.POST("/perm/del", mw_admin, crust.Permdel)
		api.POST("/perm", mw_admin, crust.Perm)

		// login
		api.POST("/login", crust.Login)

		// users
		api.POST("/snap/rollback", mw_auth_api, crust.Login)

		// rollback allow list
		api.POST("/vms", mw_auth_api, crust.Rollback_AllowList)
		api.POST("/rollback", mw_auth_api, crust.Rollback)
	}
}
