package main

import (
	"net/http"
	"testing"
)

func Test_Basic_Strapi_Get(t *testing.T) {
	resp, err := http.Get("http://thedummydogo.com/")
	if err != nil {
		t.Fatalf("Test_Basic_Strapi_Get failed, expected err to be nil, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Test_Basic_Strapi_Get failed, expected resp.StatusCode to be 200, got %v", resp.StatusCode)
	}
}
