package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://garabage.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://yahoo.com",
		"waat://garabage.com",
	}

	want := map[string]bool{
		"http://google.com":   true,
		"http://yahoo.com":    true,
		"waat://garabage.com": false,
	}

	got := ChekcWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}