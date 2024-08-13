package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.NotEmpty(t, list)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status == http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusBadRequest, status)
	}

	body := responseRecorder.Body.String()
	assert.Contains(t, body, "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

}
