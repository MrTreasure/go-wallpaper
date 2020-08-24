package downloader

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-wallpaper/unsplash"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var (
	client *resty.Client
)

const PATH = "./Photos/"
const BUFFER_SIZE = 500 * 1024

func init() {
	client = resty.New()
}

func DownloadPhoto(photoName, url string) error {
	size, err := getContentSize(url)
	fmt.Printf("DownloadPhoto-image size: %d\n", size/1024*1024)
	if err != nil {
		return err
	}
	chunks := dividerBuffer(size, BUFFER_SIZE)
	var wg sync.WaitGroup
	wg.Add(len(chunks))
	totalChunk := make([][]byte, len(chunks))

	fmt.Printf("chunks %v\n", chunks)

	for index, val := range chunks {
		fmt.Printf("DownloadPhoto-image download start %d\n", index)
		go downloadChunk(&wg, val, totalChunk, url, index)
	}

	wg.Wait()

	fmt.Printf("DownloadPhoto-image all downloaded\n")

	buffer := &bytes.Buffer{}

	for _, chunkByte := range totalChunk {
		buffer.Write(chunkByte)
	}

	err = ioutil.WriteFile(photoName+".jpg", buffer.Bytes(), 777)

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
	fmt.Printf("chunks in divide: %v \n", out)
	return out
}

func downloadChunk(wg *sync.WaitGroup, chunk []int, pools [][]byte, url string, index int) {
	Range := fmt.Sprintf("bytes=%d-%d", chunk[0], chunk[1])
	resp, err := client.R().SetHeader("Range", Range).Get(url)
	if err != nil {
		panic(err)
	}
	pools[index] = resp.Body()
	fmt.Printf("chunk done: %d\n", index)
	wg.Done()
}

func getContentSize(url string) (int, error) {
	resp, err := client.SetRetryCount(3).SetRetryWaitTime(5*time.Second).R().SetQueryParam("client_id", unsplash.ACCESS_TOKEN).Head(url)
	if err != nil || resp == nil {
		fmt.Printf("getContentSize error:\n %s\n", err.Error())
	}
	size, err := strconv.Atoi(resp.Header().Get("Content-Length"))
	return size, err
}
