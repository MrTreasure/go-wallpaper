package unsplash

import (
	"fmt"
	"testing"
)

func TestGetRandomPhoto(t *testing.T) {
	body, err := GetRandomPhoto()
	fmt.Printf("resp: %s", body)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

}
