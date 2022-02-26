package handler

import (
	"context"
	"fmt"
	"net/http"
)

func slowAPICall(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://google.com", nil)
	if err != nil {
		fmt.Println(err)
	}

	req = req.WithContext(ctx)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running server now!")
	err := slowAPICall(r.Context())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("server ran succesfully")
}
