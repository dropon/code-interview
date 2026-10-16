package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestRateLimit(t *testing.T) {
	srv := httptest.NewServer(newHandler())
	defer srv.Close()

	burst := func(n int) map[int]int {
		codes := map[int]int{}
		var mu sync.Mutex
		var wg sync.WaitGroup
		for range n {
			wg.Go(func() {
				res, err := http.Get(srv.URL + "?from=1&to=1")
				if err != nil {
					t.Error(err)
					return
				}
				res.Body.Close()
				mu.Lock()
				codes[res.StatusCode]++
				mu.Unlock()
			})
		}
		wg.Wait()
		return codes
	}

	codes := burst(12)
	if codes[http.StatusOK] != 10 || codes[http.StatusTooManyRequests] != 2 {
		t.Fatalf("first burst: got %v, want 10x200 and 2x429", codes)
	}

	// requests finish within ~1s and slots are released 1s after start at the latest
	time.Sleep(1500 * time.Millisecond)
	codes = burst(10)
	if codes[http.StatusOK] != 10 {
		t.Fatalf("after slots freed: got %v, want 10x200", codes)
	}
}
