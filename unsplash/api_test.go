package unsplash

import (
	"fmt"
	"testing"
)

func TestGetRandomPhoto(t *testing.T) {
	photo, err := GetRandomPhoto()
	if err != nil || photo == nil {
		t.Fatalf("error: %v", err)
	}
	fmt.Printf("resp: %s", photo.ID)
}
