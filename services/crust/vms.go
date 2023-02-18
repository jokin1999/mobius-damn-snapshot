package crust

import (
	"io"
	"log"
	"net/url"

	"github.com/tidwall/gjson"
)

var Vms map[string]string = map[string]string{}

func GetVms() (map[string]string, error) {
	resp, err := (&Request{
		Url:                ApiUrl("/cluster/resources?type=vm"),
		Method:             "GET",
		InsecureSkipVerify: true,
		Body:               url.Values{},
	}).Run()

	if err != nil {
		log.Println(err)
		return map[string]string{}, err
	} else {
		defer resp.Body.Close()
		// parse body
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return map[string]string{}, err
		}
		gjson.Get(string(respBytes), "data").ForEach(func(_, value gjson.Result) bool {
			if value.Get("template").String() != "0" {
				return true
			}
			if value.Get("plugintype").String() != "" {
				return true
			}
			if value.Get("type").String() != "qemu" {
				return true
			}
			Vms[value.Get("vmid").String()] = value.Get("name").String()
			return true
		})
		return Vms, nil
	}
}
