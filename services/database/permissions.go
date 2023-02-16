package database

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// register a permission
func Permission_register(user_uuid string, vmid string, status uint) int64 {
	stmt, err := DB.Prepare("INSERT INTO permissions(user_uuid, vmid, status) values (?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(user_uuid, vmid, status)
	if err != nil {
		log.Fatal(err)
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lid
}

// update a permission
func Permission_update(user_uuid string, vmid string) bool {
	stmt, err := DB.Prepare("UPDATE permissions SET vmid=? WHERE user_uuid=?;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(vmid, user_uuid)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// general permission register
func Permission_general_update(user_uuid string, vmid string, status uint) bool {
	p := Permission_by_user_uuid(user_uuid)
	fmt.Println(p.User_uuid)
	if p.User_uuid == "" {
		if Permission_register(user_uuid, vmid, status) != 0 {
			return true
		} else {
			return false
		}
	} else {
		return Permission_update(user_uuid, vmid)
	}
}

// permission
type Permission struct {
	User_uuid string
	Vmid      string
	Status    uint
}

func Permission_by_user_uuid(user_uuid string) Permission {
	stmt, err := DB.Prepare("SELECT * FROM permissions WHERE user_uuid=?;")
	if err != nil {
		log.Fatal(err)
	}

	row, err := stmt.Query(user_uuid)
	if err != nil {
		log.Fatal(err)
	}

	var p Permission

	row.Next()
	row.Scan(&p.User_uuid, &p.Vmid, &p.Status)

	defer row.Close()

	return p
}

// delete a permission
func Permission_delete_by_user_uuid(user_uuid string) bool {
	stmt, err := DB.Prepare("DELETE FROM permissions where user_uuid=?;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user_uuid)
	if err != nil {
		log.Fatal(err)
	}

	return true
}
