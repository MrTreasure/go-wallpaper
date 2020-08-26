package downloader

import (
	"fmt"
	"testing"
	"time"
)

func TestDownloadPhoto(t *testing.T) {
	start := time.Now()
	err := master("eyJhcHBfaWQiOjE1OTcyNn0", "https://images.alphacoders.com/109/1097312.png")
	spentTime := time.Since(start)
	fmt.Printf("time spent %d s", spentTime/time.Second)
	if err != nil {
		t.Fatalf("download url error: %v", err)
	}
}

func TestDisplayProgress(t *testing.T) {
	err := displayProgress("eyJhcHBfaWQiOjE1OTcyNn0", "https://images.alphacoders.com/109/1097312.png")
	if err != nil {
		t.Fatalf("error %v", err)
	}
}
