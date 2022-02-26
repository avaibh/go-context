package internal

import (
	"log"
	"net/http"
	"time"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	log.Println("[MIDDLEWARE] running now!")

	log.Println("[MIDDLEWARE] putting to sleep")
	time.Sleep(2 * time.Second)
	log.Println("[MIDDLEWARE] sleep over! Middleware is awake now")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9002/", nil)
	if err != nil {
		log.Printf("[MIDDLEWARE] %v", err)
	}

	req = req.WithContext(r.Context())
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		log.Printf("[MIDDLEWARE] %v", err)
	}

	log.Println("[MIDDLEWARE] stopped")
}
