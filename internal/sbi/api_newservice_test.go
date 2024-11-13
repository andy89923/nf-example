package sbi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 模擬 Server 結構
type mockServer struct {
	Server
}

func TestHTTPGetNewServiceInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	s := &mockServer{}
	r := gin.Default()
	r.GET("/info", s.HTTPGetNewServiceInfo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/info", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "This is the GET method of the new service", response["message"])
	assert.Equal(t, "NewService", response["service"])
}

func TestHTTPPostNewServiceData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	s := &mockServer{}
	r := gin.Default()
	r.POST("/data", s.HTTPPostNewServiceData)

	data := map[string]string{"name": "Test User"}
	jsonData, _ := json.Marshal(data)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/data", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Data received", response["message"])
	assert.Equal(t, "Test User", response["name"])
}
