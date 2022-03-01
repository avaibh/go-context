package internal

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1,
		},
		Timeout: 10 * time.Second,
	}
}

func getHTTPRequest(ctx context.Context, timeout time.Duration) (*http.Request, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("[CLIENT] reused/old connection?: %+v\n", connInfo.Reused)
		},
	}
	request, _ := http.NewRequestWithContext(httptrace.WithClientTrace(ctx, trace), http.MethodGet, "http://localhost:9001/", nil)
	return request, cancel
}

func Client(w http.ResponseWriter, r *http.Request) {
	log.Println("[CLIENT] running now!")

	// custom http client to resue the TCP connection
	c := httpClient()

	// 1st request, no conn in pool, GotConn.Reused should be false
	req, cancelContext := getHTTPRequest(r.Context(), time.Millisecond*80)
	defer cancelContext()
	_, err := c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	// 2nd request, first request should have timed out, no conn in pool, GetConn.Reused should be false
	req, cancelContext = getHTTPRequest(r.Context(), time.Millisecond*80)
	defer cancelContext()
	_, err = c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	// 3rd request, previous request should have timed out, no conn in pool. GetConn.Reused should be false
	req, cancelContext = getHTTPRequest(r.Context(), time.Second*80)
	defer cancelContext()
	_, err = c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	// 4th request, previous request hould not have timed out, 1 conn in poll, GetConn.Reused should be true
	req, cancelContext = getHTTPRequest(r.Context(), time.Second*80)
	defer cancelContext()
	_, err = c.Do(req)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}

	// 1st/2nd request show: Example above shows that when context is canceled, connection is killed and not put in conn pool
	// 3rd/4th request how: After a successful request, the connection is put in the pool and reused for subsequent requests.

	log.Println("[CLIENT] stopped")
}
