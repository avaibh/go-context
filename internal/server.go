package handler

import (
	"context"
	"fmt"
	"net/http"
)

func api(ctx context.Context) error {
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
	err := api(r.Context())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("server ran succesfully")
}
