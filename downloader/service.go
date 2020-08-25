package downloader

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"path/filepath"
)

var (
	client *resty.Client
)

const path = "./Photos/"

func init() {
	client = resty.New()
}

func getBaseURL() string {
	ex, err := os.Executable()
	fmt.Println(ex)
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
