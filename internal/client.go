package internal

import (
	"context"
	"log"
	"net/http"
	"time"
)

func Client(w http.ResponseWriter, r *http.Request) {
	log.Println("[CLIENT] running now!")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9001/", nil)
	if err != nil {
		log.Printf("[CLIENT] %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	_, _ = c.Do(req)

	log.Println("[CLIENT] ran succesfully")
}
