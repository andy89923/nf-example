package sbi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andy89923/nf-example/internal/sbi"
	"github.com/andy89923/nf-example/pkg/factory"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func Test_getNewServiceRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// 設置 mock
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

	t.Run("Get Info", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_MESSAGE = "This is the GET method of the new service"
		const EXPECTED_SERVICE = "NewService"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/newservice/info", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.HTTPGetNewServiceInfo(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		expectedBody := `{"message":"` + EXPECTED_MESSAGE + `","service":"` + EXPECTED_SERVICE + `"}`
		if httpRecorder.Body.String() != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, httpRecorder.Body.String())
		}
	})

	t.Run("Post Data", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const TEST_NAME = "Test User"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		jsonBody := `{"name":"` + TEST_NAME + `"}`
		ginCtx.Request, err = http.NewRequest("POST", "/newservice/data", strings.NewReader(jsonBody))
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}
		ginCtx.Request.Header.Set("Content-Type", "application/json")

		server.HTTPPostNewServiceData(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}
		expectedBody := `{"message":"Data received","name":"` + TEST_NAME + `"}`
		if httpRecorder.Body.String() != expectedBody {
			t.Errorf("Expected body %s, got %s", expectedBody, httpRecorder.Body.String())
		}
	})
}
