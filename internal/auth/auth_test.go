package auth

import (
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

	tests := []testCase{
		{header: http.Header{}, want: "", wantErr: ErrNoAuthHeaderIncluded},
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
