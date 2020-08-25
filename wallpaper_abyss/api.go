package wallpaper_abyss

import "github.com/go-resty/resty/v2"

const ACCESS_TOKEN = "7d699547c7918470bfc572c0fa518074"

var client *resty.Client

func init() {
	client = resty.New().SetQueryParam("auth", ACCESS_TOKEN)
}
