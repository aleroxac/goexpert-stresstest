package repo

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

type RequestRepository struct{}

func NewRequestRepository() *RequestRepository {
	return &RequestRepository{}
}

func doRequest(url string) (http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("Fail to create the request: %v", err)
		return http.Response{}, err
	}

	// start := time.Now()
	res, err := http.DefaultClient.Do(req)
	// duration := time.Since(start)
	if err != nil {
		log.Printf("Fail to make the request: %v", err)
		return *res, err
	}
	defer res.Body.Close()

	// log.Printf("Request to %s took %s", url, duration)

	ctx_err := ctx.Err()
	if ctx_err != nil {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Printf("Max timeout reached: %v", err)
			return *res, err
		}
	}

	return *res, nil
}

func (r *RequestRepository) Do(url string, requests, concurrency int) ([]int, time.Duration, error) {
	var status_code_list []int
	var mu sync.Mutex
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	start := time.Now()

	for i := 0; i < requests; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// log.Printf("Starting request #%d", i)
			res, err := doRequest(url)
			if err != nil {
				log.Printf("Request #%d failed: %v", i, err)
				return
			}

			mu.Lock()
			status_code_list = append(status_code_list, res.StatusCode)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	duration := time.Since(start)

	return status_code_list, duration, nil
}
