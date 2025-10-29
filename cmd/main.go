package main

import (
	"context"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequestWithContext(context.TODO(), "GET", "http://localhost:8080/api/v1/user/getUser", nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{

		Timeout: 3 * time.Second,
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

}
