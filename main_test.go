package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWork(t *testing.T) {
	router := getRouter()
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestAlbumDetail(t *testing.T) {
	router := getRouter()
	albumID := 3
	request, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status not ok")
	}
}

func TestAlbumNotFound(t *testing.T) {
	router := getRouter()
	albumID := 999
	request, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("must be 404")
	}
}
