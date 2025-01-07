package processor_test

import (
	"net/http/httptest"
	"testing"

	nf_context "github.com/andy89923/nf-example/internal/context"
	"github.com/andy89923/nf-example/internal/sbi/processor"
	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

func Test_Note(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	processorNf := processor.NewMockProcessorNf(mockCtrl)
	processor, err := processor.NewProcessor(processorNf)
	if err != nil {
		t.Errorf("Failed to create processor: %s", err)
		return
	}

	t.Run("Find and Show Note", func(t *testing.T) {
		const INPUT_NAME = "User_Guide"
		const EXPECTED_STATUS = 200
		const EXPECTED_BODY = "Title: " + INPUT_NAME +
			"\nContent:\n\t" +
			"/<Title>/ Show the content of the note with title <Title>.\n\n"

		processorNf.EXPECT().Context().Return(&nf_context.NFContext{
			NoteData: map[string]string{
				"User_Guide": "/<Title>/ Show the content of the note with title <Title>.\n",
			},
		})

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		processor.FindNote(ginCtx, INPUT_NAME)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Find Note That Does Not Exist", func(t *testing.T) {
		const INPUT_NAME = "Andy"
		const EXPECTED_STATUS = 404
		const EXPECTED_BODY = "[" + INPUT_NAME + "] not found in Notebook\n"

		processorNf.EXPECT().Context().Return(&nf_context.NFContext{
			NoteData: map[string]string{
				"User_Guide": "/<Title>/ Show the content of the note with title <Title>.\n",
			},
		})

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		processor.FindNote(ginCtx, INPUT_NAME)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Update Note", func(t *testing.T) {
		const INPUT_NAME = "new_note_title"
		const INPUT_CONTENT = "Content"
		const EXPECTED_STATUS = 200
		const EXPECTED_BODY = "Title: " + INPUT_NAME +
			"\nContent:\n\t" + INPUT_CONTENT + "\n"

		processorNf.EXPECT().Context().Return(&nf_context.NFContext{
			NoteData: map[string]string{
				"User_Guide": "/<Title>/ Show the content of the note with title <Title>.\n",
			},
		}).AnyTimes()

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		processor.UpdateNote(ginCtx, INPUT_NAME, INPUT_CONTENT)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

	t.Run("Append with Whitespace Prefix on Note", func(t *testing.T) {
		const INPUT_NAME = "new_note_title"
		const INPUT_CONTENT = "can't_contain_whitespace."
		const EXPECTED_STATUS = 200
		const EXPECTED_BODY = "Title: " + INPUT_NAME +
			"\nContent:\n\tContent " + INPUT_CONTENT + "\n"

		processorNf.EXPECT().Context().Return(&nf_context.NFContext{
			NoteData: map[string]string{
				"User_Guide": "/<Title>/ Show the content of the note with title <Title>.\n",
			},
		}).AnyTimes()

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		processor.NoteWhitespaceAppend(ginCtx, INPUT_NAME, INPUT_CONTENT)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})
}
