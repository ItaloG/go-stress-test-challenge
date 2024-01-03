package internal

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPServer struct{}

func (m *MockHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestPerformStressTest(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 100
	concurrency := 5

	duration, statusCodes := PerformStressTest(url, totalRequests, concurrency)

	assert.True(t, duration.Seconds() > 0)
	assert.Equal(t, totalRequests, statusCodes[http.StatusOK])
	assert.Equal(t, 0, statusCodes[http.StatusNotFound])
}

func TestRunRequests(t *testing.T) {
	mockServer := httptest.NewServer(&MockHTTPServer{})
	defer mockServer.Close()

	url := mockServer.URL
	totalRequests := 10

	var wg sync.WaitGroup
	wg.Add(1)

	resultCh := make(chan int)

	go runRequests(1, url, totalRequests, resultCh, &wg)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	statusCount := 0
	for statusCode := range resultCh {
		if statusCode == http.StatusOK {
			statusCount++
		}
	}

	assert.Equal(t, totalRequests, statusCount)
}
