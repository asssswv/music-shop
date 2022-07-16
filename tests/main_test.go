package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/asadbek280604/server_on_golang_example/pkg/handler"
)

func handelRequest(w *httptest.ResponseRecorder, request *http.Request) {
	router := handler.GetRouter()
	router.ServeHTTP(w, request)
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

func TestWork(t *testing.T) {
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok")
	}
}

func TestAlbumDetail(t *testing.T) {
	albumID := 4
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

func TestUpdateAlbumNotFound(t *testing.T) {
	albumID := 999
	request, _ := http.NewRequest("PUT", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404", w.Code)
	}
}

func TestUpdateAlbum(t *testing.T) {
	albumID := 4
	request, _ := http.NewRequest("PUT", fmt.Sprintf("/albums/%d", albumID), strings.NewReader(`{"title": "test"}"`))
	w := httptest.NewRecorder()
	handelRequest(w, request)
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok")
	}
}

func TestDeleteAlbum(t *testing.T) {
	albumID := 4
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
