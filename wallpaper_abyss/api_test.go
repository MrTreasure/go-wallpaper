package wallpaper_abyss

import (
	"fmt"
	"testing"
)

func TestGetPhotoByRandom(t *testing.T) {
	url, err := GetPhotoListByRandom()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	fmt.Printf("url: %s\n", url)
}
