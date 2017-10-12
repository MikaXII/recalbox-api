package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/MikaXII/recalbox-api/models"
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

func TestMD5ToString(t *testing.T) {
	hash := models.NewHash("./hash.txt", "MD5")
	if hash.StringValue != "669fe6c675fd978dc36239db70d340d4" {
		t.Fatalf("Non expected MD5, should be %s, get : %s", "669fe6c675fd978dc36239db70d340d4", hash.StringValue)
	}
}

func TestSHA1ToString(t *testing.T) {
	hash := models.NewHash("./hash.txt", "SHA1")
	if hash.StringValue != "bd4cc7b0953bc366c42517515611a10988afafb9" {
		t.Fatalf("Non expected SHA1, should be %s, get : %s", "bd4cc7b0953bc366c42517515611a10988afafb9", hash.StringValue)
	}
}
func TestCRC32ToString(t *testing.T) {
	hash := models.NewHash("./hash.txt", "CRC")
	if hash.StringValue != "35d3c9b8" {
		t.Fatalf("Non expected CRC, should be %s, get : %s", "35d3c9b8", hash.StringValue)
	}
}
