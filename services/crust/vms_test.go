package crust_test

import (
	"log"
	"testing"

	"github.com/jokin1999/mobius-damn-snapshot/services/crust"
)

func Test_GetVms(t *testing.T) {
	r, e := crust.GetVms()
	if e != nil {
		log.Println(e)
	}
	log.Println(r)
}
