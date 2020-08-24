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

const PATH = "Photos/"
const BUFFER_SIZE = 500 * 1024

func init() {
	client = resty.New()
}

func DownloadPhoto(photoName, url string) error {
	_, err := client.R().SetOutput(PATH + photoName).Get(url)
	return err
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

func dividerBuffer(size, chunkSize int) [][]int {
	turn := size / chunkSize
	out := make([][]int, turn)

	for i := 0; i < turn; i++ {
		out[i] = make([]int, 2)
		if i == 0 {
			out[i][0] = 0
		} else {
			out[i][0] = out[i-1][1] + 1
		}
		next := out[i][0] + chunkSize
		if next > size {
			next = size
		}
		out[i][1] = next
	}
	return out
}
