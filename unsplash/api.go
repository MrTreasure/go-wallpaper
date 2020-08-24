package unsplash

import (
	"encoding/json"
	"fmt"
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

func GetRandomPhoto() (*Photo, error) {
	photo := &Photo{}
	res, err := client.R().Get("photos/random")
	if err != nil {
		log.Printf("GetRandomPhoto error: %v", err)
		return nil, err
	}

	err = json.Unmarshal(res.Body(), photo)
	if err != nil {
		log.Printf("unmarshal error: %v", err)
		return nil, err
	}
	fmt.Println(photo.ID)
	return photo, nil
}

func GetDownloadURL(id string) (string, error) {
	downloadURL := &DownloadUrl{}
	res, err := client.R().SetPathParams(map[string]string{
		"id": id,
	}).Get("/photos/{id}/download")

	if err != nil {
		log.Printf("ger download url error: %v", err)
		return "", err
	}
	_ = json.Unmarshal(res.Body(), downloadURL)
	return downloadURL.URL, nil
}
