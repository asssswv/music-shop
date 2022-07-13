package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handelRequest(w *httptest.ResponseRecorder, request *http.Request) {
	router := getRouter()
	router.ServeHTTP(w, request)
}

func TestWork(t *testing.T) {
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok")
	}
}

func TestAlbumDetail(t *testing.T) {
	albumID := 3
	request, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok")
	}
}

func TestAlbumNotFound(t *testing.T) {
	albumID := 999
	request, _ := http.NewRequest("GET", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestDeleteAlbum(t *testing.T) {
	albumID := 1
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusNoContent {
		t.Fatal("status must be 204")
	}
}

func TestDeleteAlbumNotFound(t *testing.T) {
	albumID := 999
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateAlbumNotFound(t *testing.T) {
	albumID := 999
	request, _ := http.NewRequest("PUT", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateAlbum(t *testing.T) {
	albumID := 2
	request, _ := http.NewRequest("PUT", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(`{"title": "test"}"`))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok")
	}
}

func TestCreateBadStorage(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusBadRequest {
		t.Fatal("status must be 400", w.Code)
	}
}

func TestCreateAlbum(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4", "title": "a", "artist": "asd", "price": 49.99}`))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201", w.Code)
	}
}
