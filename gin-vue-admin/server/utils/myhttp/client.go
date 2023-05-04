package myhttp

import "github.com/go-resty/resty/v2"

// resty httpclient
func HttpClinet() *resty.Client {
	return resty.New()
}
