package unsplash

import (
	"fmt"
	"testing"
)

func TestGetRandomPhoto(t *testing.T) {
	url, err := GetRandomPhoto()
	fmt.Printf("resp: %s", url)
	if err != nil || url == "" {
		t.Fatalf("error: %v", err)
	}
}
