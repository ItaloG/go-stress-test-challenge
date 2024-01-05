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
	resultCh := make(chan int)

	if concurrency >= totalRequests {
		for i := 0; i < totalRequests; i++ {
			wg.Add(1)
			go runRequests(url, resultCh, &wg)
		}
	} else {
		isDivisible := totalRequests % concurrency

		if isDivisible == 0 {
			remainingRequest := totalRequests
			manyIterations := totalRequests / concurrency
			for i := 0; i < manyIterations; i++ {
				if remainingRequest <= 0 {
					continue
				}
				for i := 0; i < concurrency; i++ {
					wg.Add(1)
					go runRequests(url, resultCh, &wg)
				}
				remainingRequest = remainingRequest - concurrency
			}
		} else {
			restRequests := totalRequests % concurrency
			for i := 0; i < restRequests; i++ {
				wg.Add(1)
				go runRequests(url, resultCh, &wg)
			}
			remainingRequest := totalRequests - restRequests
			manyIterations := totalRequests / concurrency
			for i := 0; i < manyIterations; i++ {
				if remainingRequest <= 0 {
					continue
				}
				for i := 0; i < concurrency; i++ {
					wg.Add(1)
					go runRequests(url, resultCh, &wg)
				}
				remainingRequest = remainingRequest - concurrency
			}
		}
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

func runRequests(url string, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		resultCh <- 0
		return
	}
	resultCh <- response.StatusCode
}
