package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type testCase struct {
		header  http.Header
		want    string
		wantErr error
	}

	goodHeader := http.Header{}
	goodHeader.Add("Authorization", "ApiKey 01")

	badHeader := http.Header{}
	badHeader.Add("Authorization", "Winchester")

	tests := []testCase{
		{header: http.Header{}, want: "", wantErr: ErrNoAuthHeaderIncluded},
		{header: goodHeader, want: "01", wantErr: nil},
		{header: badHeader, want: "", wantErr: errors.New("malformed authorization header")},
	}

	for _, test := range tests {
		got, gotErr := GetAPIKey(test.header)

		if !reflect.DeepEqual(test.want, got) {
			t.Fatalf("expected: %v, got: %v", test.want, got)
		}

		if !reflect.DeepEqual(test.wantErr, gotErr) {
			t.Fatalf("expected Error: %v, got Error: %v", test.wantErr, gotErr)
		}
	}
}
