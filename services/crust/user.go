package crust

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/database"
)

// register user
func Reguser(c *gin.Context) {
	// status = 1 active
	username := c.PostForm("username")
	password := c.PostForm("password")
	// username string, password string
	res := database.User_register(username, password, 1)
	if !res {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Failed to register user",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
		})
	}
}

// delete user
func Deluser(c *gin.Context) {
	// status = 1 active
	username := c.PostForm("username")
	// username string, password string
	res := database.User_delete_by_username(username)
	if !res {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Failed to delete user",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
		})
	}
}

// user login
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// username and password shoule not be empty
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Empty username or password",
		})
		return
	}

	user := database.User_by_username(username)
	if user.Username == username && user.Password == password {
		// auth := genAuth(username, password)
		auth := gen_sha256(username)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
			"data": auth,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "Invalid username or password",
		})
	}
}

func genAuth(username string, password string) string {
	return gen_md5(gen_sha1(gen_md5(username) + gen_sha256(password)))
}

func gen_md5(str string) string {
	hex := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", hex)
}

func gen_sha1(str string) string {
	hex := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", hex)
}

func gen_sha256(str string) string {
	hex := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", hex)
}
