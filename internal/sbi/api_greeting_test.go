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

func Test_getGreetingRoutes(t *testing.T) {
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

	t.Run("GetGreeting", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = "Greetings!\n"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("GET", "/greeting", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.GetGreeting(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("PostFarewell", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const EXPECTED_BODY = "Farewells!\n"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		ginCtx.Request, err = http.NewRequest("POST", "/greeting", nil)
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.PostFarewell(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Farewell to", func(t *testing.T) {
		const EXPECTED_STATUS = http.StatusOK
		const NAME = "Alisa"
		const EXPECTED_BODY = "Farewell ~ " + NAME + "!"

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)

		var err error
		jsonBody := `{"name":"` + NAME + `"}`
		ginCtx.Request, err = http.NewRequest("POST", "/greeting/to", strings.NewReader(jsonBody))
		if err != nil {
			t.Errorf("Failed to create request: %s", err)
			return
		}

		server.Greetingto(ginCtx)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
