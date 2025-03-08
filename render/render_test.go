package render

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/saeedafzal/sandbox-service/tester/assert"
)

func TestPlainText(t *testing.T) {
	recorder := httptest.NewRecorder()

	status := http.StatusOK
	body := "hello world"
	PlainText(recorder, status, body)

	assert.Equals(t, status, recorder.Code)
	assert.Equals(t, "text/plain; charset=utf-8", recorder.Header().Get("Content-Type"))
	assert.Equals(t, body, recorder.Body.String())
}
