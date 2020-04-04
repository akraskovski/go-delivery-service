package delivery_service

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	_ = NewAPIServer(NewDeliveryServiceConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	handleHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Welcome to the club, buddy\n")
}
