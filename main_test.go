package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	rootHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	responseMsg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(responseMsg, []byte(rootResponseMsg)) {
		t.Fatalf("expected '%s', got '%s'", rootResponseMsg, responseMsg)
	}
}
