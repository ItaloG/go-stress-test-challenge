package internal

import (
	"net/http"
	"sync"
	"time"
)

func PerformStressTest(url string, totalRequests int, concurrency int) (time.Duration, map[int]int) {
	startTime := time.Now()
	var wg sync.WaitGroup
	successfulRequests := 0
	statusCodes := make(map[int]int)
	requestPerThread := totalRequests / concurrency

	resultCh := make(chan int)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go runRequests(i, url, requestPerThread, resultCh, &wg)

	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for statusCode := range resultCh {
		if statusCode == http.StatusOK {
			successfulRequests++
		}
		statusCodes[statusCode]++
	}

	duration := time.Since(startTime)
	return duration, statusCodes
}

func runRequests(i int, url string, totalRequests int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < totalRequests; i++ {
		response, err := http.Get(url)
		if err != nil {
			resultCh <- 0
			continue
		}
		resultCh <- response.StatusCode
	}
}
