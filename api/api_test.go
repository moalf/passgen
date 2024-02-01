package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetPassword(t *testing.T) {
	tests := []struct {
		name   string
		want   int
		length int
		method string
		path   string
	}{
		// simple
		{"single 12-chars password", 1, 12, http.MethodGet, "/"},
		{"two 13-chars passwords", 2, 13, http.MethodGet, "/2/13"},
		{"three 20-chars passwords", 3, 20, http.MethodGet, "/3/20"},

		// request for one 65-chars password returns single 12-chars one
		{"max password length free limit", 1, 12, http.MethodGet, "/1/65"},
		// request for 11 passwords returns single 12-chars one
		{"max number of passwords free limit", 1, 12, http.MethodGet, "/11"},
		// request for 20 passwords of 65-chars returns single 12-chars one
		{"beyond default free limits", 1, 12, http.MethodGet, "/20/65"},
	}

	for _, tc := range tests {
		request := httptest.NewRequest(tc.method, tc.path, nil)
		responseRecorder := httptest.NewRecorder()

		GetPassword(responseRecorder, request)
		response := responseRecorder.Result()
		defer response.Body.Close()

		content, err := io.ReadAll(response.Body)

		if err != nil {
			t.Errorf("Error: %v", err)
		}

		var data Data
		err = json.Unmarshal(content, &data)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if tc.want != len(data.Passwords) {
			t.Errorf("'%s' failed, wanted %d password(s), got %d", tc.name, tc.want, len(data.Passwords))
		}

		if strconv.Itoa(tc.length) != data.Details.Specs.Length {
			t.Errorf("'%s' failed, wanted password length of %d, got %s", tc.name, tc.length, data.Details.Specs.Length)
		}
	}
}
