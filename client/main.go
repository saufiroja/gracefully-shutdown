package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	now := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := fmt.Sprint(i)
			checkReadiness(id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\nTTC: %v\n\n", time.Since(now))
}

func checkReadiness(id string) {
	const url = "http://localhost:8080/count"
	client := &http.Client{}

	fmt.Println("START REQUEST ID:", id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-REQUEST-ID", id)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR REQUEST ID:", id, err.Error())
	} else {
		fmt.Println("DONE  REQUEST ID:", id, http.StatusText(resp.StatusCode))
	}
}
