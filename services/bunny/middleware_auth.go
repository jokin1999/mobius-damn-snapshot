package bunny

import (
	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/crust"
)

func mw_admin(c *gin.Context) {
	h := c.Request.Header.Get("A")
	if h != "Joe" {
		c.JSON(404, gin.H{
			"code": 404,
		})
		c.Abort()
	}
}

func mw_auth_api(c *gin.Context) {
	u, err := c.Request.Cookie("u")
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
		})
		c.Abort()
	}
	a, err := c.Request.Cookie("a")
	if err != nil {
		c.JSON(401, gin.H{
			"code": 401,
		})
		c.Abort()
	}
	username := u.Value
	auth := a.Value
	// check auth
	real_auth := crust.GenAuthToken_by_username(username)
	if auth != real_auth {
		c.JSON(401, gin.H{
			"code": 401,
		})
		c.Abort()
		return
	}
}
