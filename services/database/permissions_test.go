package database_test

import (
	"fmt"
	"testing"

	"github.com/jokin1999/mobius-damn-snapshot/services/database"
	_ "github.com/mattn/go-sqlite3"
)

func Test_Permission_register(t *testing.T) {
	database.Init()
	res := database.Permission_register("8ced6a55-0cfe-4879-8955-5bdbff3c9a64", "10001", 1)
	fmt.Println(res)
}

func Test_Permission_update(t *testing.T) {
	database.Init()
	database.Permission_update("8ced6a55-0cfe-4879-8955-5bdbff3c9a64", "10000")
	res := database.Permission_update("8ced6a55-0cfe-4879-8955-5bdbff3c9a64", "10001")
	fmt.Println(res)
}

func Test_Permission_general_update(t *testing.T) {
	database.Init()
	res := database.Permission_general_update("8ced6a55-0cfe-4879-8955-5bdbff3c9a64", "10002", 1)
	fmt.Println(res)
}

func Test_Permission_by_uuid(t *testing.T) {
	database.Init()
	res := database.Permission_by_user_uuid("8ced6a55-0cfe-4879-8955-5bdbff3c9a64")
	fmt.Println(res)
}

func Test_Permission_delete_by_user_uuid(t *testing.T) {
	database.Init()
	res := database.Permission_delete_by_user_uuid("8ced6a55-0cfe-4879-8955-5bdbff3c9a64")
	fmt.Println(res)
}
