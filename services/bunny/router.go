package bunny

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/public"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiv1(r)

	r.Any("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/ui")
	})

	r.Any("/ui", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/ui/index.html")
	})

	r.StaticFS("/ui", http.FS(public.DistFS))

	return r
}
