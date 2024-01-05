package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPServer struct{}

func (m *MockHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestPerformStressWithConcurrencyMoreThanTotalRequestsTest(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 10
	concurrency := 15

	duration, statusCodes := PerformStressTest(url, totalRequests, concurrency)

	assert.True(t, duration.Seconds() > 0)
	assert.Equal(t, totalRequests, statusCodes[http.StatusOK])
	assert.Equal(t, 0, statusCodes[http.StatusNotFound])
}

func TestPerformStressWithConcurrencyEqualToTotalRequestsTest(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 10
	concurrency := 10

	duration, statusCodes := PerformStressTest(url, totalRequests, concurrency)

	assert.True(t, duration.Seconds() > 0)
	assert.Equal(t, totalRequests, statusCodes[http.StatusOK])
	assert.Equal(t, 0, statusCodes[http.StatusNotFound])
}

func TestPerformStressWithConcurrencyLessThanTotalRequestsAndIsDivisibleTest(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 1000
	concurrency := 20

	duration, statusCodes := PerformStressTest(url, totalRequests, concurrency)

	assert.True(t, duration.Seconds() > 0)
	assert.Equal(t, totalRequests, statusCodes[http.StatusOK])
	assert.Equal(t, 0, statusCodes[http.StatusNotFound])
}

func TestPerformStressWithConcurrencyLessThanTotalRequestsAndIsNotDivisibleTest(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 10
	concurrency := 6

	duration, statusCodes := PerformStressTest(url, totalRequests, concurrency)

	assert.True(t, duration.Seconds() > 0)
	assert.Equal(t, totalRequests, statusCodes[http.StatusOK])
	assert.Equal(t, 0, statusCodes[http.StatusNotFound])
}
