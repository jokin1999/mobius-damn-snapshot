package crust

import (
	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/database"
)

// register user
func Reguser(c *gin.Context) {
	// status = 1 active
	username := c.PostForm("username")
	password := c.PostForm("password")
	// username string, password string
	database.User_register(username, password, 1)
}
