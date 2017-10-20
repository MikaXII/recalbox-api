package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// getRouter for test files
func getRouter() *gin.Engine {
	return GetMainEngine()
}

func TestAPIInfo(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	r := getRouter()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", w.Code)
	}
}
