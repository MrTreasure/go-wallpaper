package unsplash

import (
	"github.com/go-resty/resty/v2"
	"log"
)

const (
	ACCESS_TOKEN = "OfJriVm_1LgBX3GlWlmtFZwd0srr5ZtRSzo2jxm-ulI"
	API          = "https://api.unsplash.com/"
)

var client *resty.Client

func init() {
	client = resty.New().SetHostURL(API).SetQueryParam("client_id", ACCESS_TOKEN)
}

func GetRandomPhoto() (string, error) {
	res, err := client.R().Get("photos/random")
	if err != nil {
		log.Printf("GetRandomPhoto error: %v", err)
		return "", err
	}
	return string(res.Body()), nil
}
