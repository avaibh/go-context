package internal

import (
	"context"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func Client(w http.ResponseWriter, r *http.Request) {
	log.Println("[CLIENT] running now!")

	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("[CLIENT] conn was reused: %t\n", info.Reused) },
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()
	traceCtx := httptrace.WithClientTrace(ctx, clientTrace)

	// 1st request
	req, _ := http.NewRequestWithContext(traceCtx, http.MethodGet, "http://localhost:9001", nil)
	_, _ = http.DefaultClient.Do(req)

	// 2nd request
	req, _ = http.NewRequestWithContext(traceCtx, http.MethodGet, "http://localhost:9001", nil)
	_, _ = http.DefaultClient.Do(req)

	log.Println("[CLIENT] ran succesfully")
}
