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

func Test_getHelloRoutes(t *testing.T) {
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

	t.Run("GetHello", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = "Hello from feature/hello!"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.GetHello(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("PostHello", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const TEST_MESSAGE = "test"
		const EXPECTED_BODY = "Received message: test"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("POST", "/"+TEST_MESSAGE, nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}
		ginCtx.Params = []gin.Param{{Key: "message", Value: TEST_MESSAGE}}

		server.PostHello(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
