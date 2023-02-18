package crust

import "log"

func ApiUrl(path string) string {
	// return "http://192.168.90.21:8080/api/v2/c3" + path
	u := "http://192.168.90.41:8080/api/v2/default" + path
	log.Println(u)
	return u
}
