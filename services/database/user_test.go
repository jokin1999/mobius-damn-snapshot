package database_test

import (
	"fmt"
	"testing"

	"github.com/jokin1999/mobius-damn-snapshot/services/database"
	_ "github.com/mattn/go-sqlite3"
)

func Test_User_register(t *testing.T) {
	database.Init()
	res := database.User_register("tom", "123456", 1)
	fmt.Println(res)
}

func Test_User_delete_by_uuid(t *testing.T) {
	database.Init()
	res := database.User_delete_by_uuid("09ef0a4b-ae0b-4729-bbb0-f7c3ba103f04")
	fmt.Println(res)
}

func Test_User_delete_by_username(t *testing.T) {
	database.Init()
	res := database.User_delete_by_username("tom")
	fmt.Println(res)
}

func Test_User_change_password_by_username(t *testing.T) {
	database.Init()
	res := database.User_change_password_by_username("tom", "1234")
	fmt.Println(res)
}

func Test_User_by_username(t *testing.T) {
	database.Init()
	res := database.User_by_username("tom")
	fmt.Println(res)
}
