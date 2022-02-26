package internal

import (
	"context"
	"log"
	"net/http"
)

func api(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	c := &http.Client{}
	_, err = c.Do(req)
	return err
}

func Server(w http.ResponseWriter, r *http.Request) {
	log.Println("[SERVER] running now!")
	err := api(r.Context())
	if err != nil {
		log.Printf("[SERVER] %v", err)
	}

	log.Println("[SERVER] stopped")
}
