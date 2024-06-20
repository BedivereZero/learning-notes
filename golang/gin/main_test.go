package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewLogMiddleware(t *testing.T) {
	e := gin.Default()

	registerAPI(e)

	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/a/b/c", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	e.ServeHTTP(w, req)

	t.Log("code", w.Code)
}
