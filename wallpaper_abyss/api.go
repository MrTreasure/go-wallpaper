package wallpaper_abyss

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const ACCESS_TOKEN = "7d699547c7918470bfc572c0fa518074"
const BASE_URL = "https://wall.alphacoders.com/api2.0/get.php"

var client *resty.Client

func init() {
	client = resty.New().SetQueryParam("auth", ACCESS_TOKEN).SetHostURL(BASE_URL)
	SetDefaultConfig()
}

func SetDefaultConfig() {
	client.SetQueryParam("width", "1920").SetQueryParam("height", "1080").SetQueryParam("operator", "min")
}

func GetPhotoListByRandom() (string, error) {
	resp, err := client.R().SetQueryParam("method", "random").Get("")
	if err != nil {
		return "", err
	}
	randomEntity := &Res{}
	err = json.Unmarshal(resp.Body(), randomEntity)
	if err != nil {
		return "", err
	}
	return randomEntity.Wallpapers[0].URLImage, nil
}
