package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running middleware now!")
	time.Sleep(2 * time.Second)
	fmt.Println("2 sec over, calling server now")
	req, err := http.NewRequest(http.MethodGet, "http://localhost:9002/", nil)
	if err != nil {
		fmt.Println(err)
	}

	req = req.WithContext(r.Context())
	c := &http.Client{}
	_, _ = c.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("middleware ran succesfully")
}
