package crust

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/database"
)

// add permissions of a user
func Permadd(c *gin.Context) {
	username := c.PostForm("username")
	vmid := c.PostForm("vmid")

	// check if username is valid
	user := database.User_by_username(username)
	if user.Username != username {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "Failed to find username",
		})
		return
	}
	res := database.Permission_general_update(user.Uuid, vmid, 1)
	if !res {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Failed to update permission",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
		})
	}
}

// del permissions of a user
func Permdel(c *gin.Context) {
	username := c.PostForm("username")

	// check if username is valid
	user := database.User_by_username(username)
	if user.Username != username {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "Failed to find username",
		})
		return
	}
	res := database.Permission_delete_by_user_uuid(user.Uuid)
	if !res {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Failed to delete permission",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
		})
	}
}

// show permission
func Perm(c *gin.Context) {
	username := c.PostForm("username")

	// check if username is valid
	user := database.User_by_username(username)
	if user.Username != username {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "Failed to find username",
		})
		return
	}
	res := database.Permission_by_user_uuid(user.Uuid)
	fmt.Println(res)
	if res.User_uuid != user.Uuid {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Failed to find permission",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
			"data": res,
		})
	}
}
