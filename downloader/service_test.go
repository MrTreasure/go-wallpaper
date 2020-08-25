package downloader

import (
	"fmt"
	"testing"
	"time"
)

func TestDownloadPhoto(t *testing.T) {
	start := time.Now()
	err := master("eyJhcHBfaWQiOjE1OTcyNn0", "https://images.alphacoders.com/539/539674.jpg")
	spentTime := time.Since(start)
	fmt.Printf("time spent %d s", spentTime/time.Second)
	if err != nil {
		t.Fatalf("download url error: %v", err)
	}
}
