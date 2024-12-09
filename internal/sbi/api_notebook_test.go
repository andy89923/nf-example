package sbi_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andy89923/nf-example/internal/sbi"
	"github.com/andy89923/nf-example/pkg/factory"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func Test_getNotebookRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	nfApp := sbi.NewMocknfApp(mockCtrl)
	nfApp.EXPECT().Config().Return(&factory.Config{
		Configuration: &factory.Configuration{
			Sbi: &factory.Sbi{
				Port: 8000,
			},
		},
	}).AnyTimes()
	server := sbi.NewServer(nfApp, "")

	t.Run("No name provided", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusBadRequest
		const EXPECTED_BODY = "No name provided"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/notebook/", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.HTTPShowNote(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("No name provided", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusBadRequest
		const EXPECTED_BODY = "No name provided"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("POST", "/notebook/", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.HTTPUpdateNote(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("No name provided", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusBadRequest
		const EXPECTED_BODY = "No name provided"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("POST", "/notebook//append/", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.HTTPUpdateNote(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
