package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed auth",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("TestAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected test case error: %v", err)
				return
			}
			if output != test.expect {
				t.Errorf("Undexpected: TestGetAPIKey: %v", output)
				return
			}
		})
	}
}
