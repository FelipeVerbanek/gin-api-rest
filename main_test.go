package main

import (
	"gin-api-rest/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRoutesTest() *gin.Engine {
	return gin.Default()
}

func TestCheckStatusCodeNotFound(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:nome", controllers.NotFound)

	req, _ := http.NewRequest("GET", "/gui", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusNotFound {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resp.Code, http.StatusOK)
	}
}
