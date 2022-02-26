package internal

import (
	"fmt"
	"net/http"
	"time"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running middleware now!")

	fmt.Println("putting middleware to sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("sleep over! middleware is awake now")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:9002/", nil)
	if err != nil {
		fmt.Println(err)
	}

	req = req.WithContext(r.Context())
	c := &http.Client{}
	_, _ = c.Do(req)

	fmt.Println("middleware ran succesfully")
}
