package crust

import (
	"log"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jokin1999/mobius-damn-snapshot/services/database"
)

func Rollback_AllowList(c *gin.Context) {
	u, _ := c.Request.Cookie("u")
	username := u.Value
	user := database.User_by_username(username)
	uuid := user.Uuid

	userperms := database.Permission_by_user_uuid(uuid)
	if uuid != userperms.User_uuid {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Permission mismatch",
		})
	}

	// check permission
	perms := strings.Split(userperms.Vmid, ",")

	// get vm list
	vms, err := GetVms()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Vm list mismatch",
		})
		return
	}

	// filter vms
	permVms := map[string]string{}
	for _, vmid := range perms {
		log.Println(vmid)
		if v, ok := vms[vmid]; ok {
			_v := strings.Split(v, "-")
			if len(_v) > 3 {
				_v = _v[len(_v)-3:]
				v = strings.Join(_v, "-")
			}
			permVms[vmid] = v
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": permVms,
	})

}

func Rollback(c *gin.Context) {
	u, _ := c.Request.Cookie("u")
	username := u.Value
	vmid := c.PostForm("vmid")
	user := database.User_by_username(username)
	uuid := user.Uuid

	userperms := database.Permission_by_user_uuid(uuid)
	if uuid != userperms.User_uuid {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Permission mismatch",
		})
		return
	}

	// check permission
	perms := strings.Split(userperms.Vmid, ",")
	rollback_flag := false
	for _, v := range perms {
		if v == vmid {
			rollback_flag = true
			break
		}
	}

	if rollback_flag {
		// rollback vm
		node := "pve08"
		snapname := "origin"
		_, err := (&Request{
			Url:                ApiUrl("/node/" + node + "/vmid/" + vmid + "/snapshot/" + snapname + "/rollback"),
			Method:             "POST",
			InsecureSkipVerify: true,
			Body: url.Values{
				"start": {"1"},
			},
		}).Run()
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "rollback failed",
			})
			return
		}

		//NOTE - No checking rollback status here, should add one
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "Not your vm",
		})
	}

}
