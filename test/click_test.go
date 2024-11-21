package test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const (
	numRequests    = 500
	bannerID       = 1
	serverAddr     = "http://localhost:9999"
	expectedStatus = http.StatusOK
	expectedBody   = "click for banner 1 incremented"
	concurrency    = 10
)

func TestIncrementClickAPI(t *testing.T) {
	e := echo.New()

	server := httptest.NewServer(e)
	defer server.Close()

	sem := make(chan struct{}, concurrency)

	sendRequest := func(wg *sync.WaitGroup) {
		defer wg.Done()

		sem <- struct{}{}
		defer func() { <-sem }()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/counter/%d", serverAddr, bannerID), nil)
		if err != nil {
			t.Errorf("failed create new request: %s", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Errorf("failed making http request: %s", err)
			return
		}

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("failed reading response body: %s", err)
			return
		}

		assert.Equal(t, expectedStatus, resp.StatusCode)
		assert.Equal(t, string(resBody), expectedBody)
	}

	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go sendRequest(&wg)
	}

	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("time taken for %d requests: %v", numRequests, duration)
}
