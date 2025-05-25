package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	headers := http.Header {
		"Authorization": []string{"ApiKey my-test-key"},
	}
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey() unexpected error: %v", err)
	}
	want := "my-test-key"
	if got != want {
		t.Errorf("GetAPIKey() %v, want %v", got, want)
	}
}