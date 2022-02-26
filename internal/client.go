package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Client(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nrunning client now!")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9001/", nil)
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()

	req = req.WithContext(ctx)
	c := &http.Client{}
	_, _ = c.Do(req)

	fmt.Println("client ran succesfully")
}
