package crust

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	// target url
	Url string

	// request method
	Method string

	// common header or auth header
	CommonHeader map[string][]string

	// request body
	Body url.Values

	// if skip certificate verification check
	InsecureSkipVerify bool

	// request
	req *http.Request
}

// add request header
func (r *Request) AddHeader(key string, value string) {
	r.req.Header.Set(key, value)
}

// execute request
func (r *Request) Run() (*http.Response, error) {

	// new request
	if len(r.Body) == 0 {
		r.req, _ = http.NewRequest(r.Method, r.Url, nil)
	} else {
		r.req, _ = http.NewRequest(r.Method, r.Url, strings.NewReader(r.Body.Encode()))
	}

	r.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// common header
	if len(r.CommonHeader) != 0 {
		for key, value := range r.CommonHeader {
			r.req.Header.Set(key, strings.Join(value, ","))
		}
	}

	// transport settings
	tr := &http.Transport{}

	// if skip certificate verification check
	if r.InsecureSkipVerify {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// send request
	resp, err := (&http.Client{
		Transport: tr,
	}).Do(r.req)

	// check response error
	if err != nil {
		return nil, err
	}

	return resp, nil
}
