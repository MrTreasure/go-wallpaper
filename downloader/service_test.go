package downloader

import (
	"fmt"
	"go-wallpaper/unsplash"
	"reflect"
	"testing"
	"time"
)

func TestDownloadPhoto(t *testing.T) {
	start := time.Now()
	photo, err := unsplash.GetRandomPhoto()
	if err != nil {
		t.Fatalf("get url error: %v", err)
	}
	_, err = unsplash.GetDownloadURL(photo.ID)

	err = DownloadPhoto(photo.ID, photo.Urls.Small)
	spentTime := time.Since(start)
	fmt.Printf("time spent %d s", spentTime/time.Second)
	if err != nil {
		t.Fatalf("download url error: %v", err)
	}
}

func TestGetBaseURL(t *testing.T) {
	path := getBaseURL()
	if path == "" {
		t.Fatalf("get baseURL error")
	}
	fmt.Printf("path: %s", path)
}

func TestDivideBuffer(t *testing.T) {
	size, chunkSize := 23, 4
	out := dividerBuffer(size, chunkSize)
	want := [][]int{[]int{0, 4}, []int{5, 9}, []int{10, 14}, []int{15, 19}, []int{20, 23}}

	for index, val := range out {
		if !reflect.DeepEqual(val, out[index]) {
			t.Fatalf("divide faild, want %v, got %v", want, out)
		}
	}
}

func TestGetSize(t *testing.T) {
	photo, err := unsplash.GetRandomPhoto()
	if err != nil {
		t.Fatalf("get url error: %v", err)
	}
	url, err := unsplash.GetDownloadURL(photo.ID)
	size, err := getContentSize(url)
	fmt.Printf("size: %d", size)
	if size == 0 {
		t.Fatalf("error size: %d", size)
	}
}
