package downloader

import (
	"fmt"
	"go-wallpaper/unsplash"
	"reflect"
	"testing"
)

func TestDownloadPhoto(t *testing.T) {
	photo, err := unsplash.GetRandomPhoto()
	if err != nil {
		t.Fatalf("get url error: %v", err)
	}
	url, err := unsplash.GetDownloadURL(photo.ID)
	err = DownloadPhoto(photo.Description, url)
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
