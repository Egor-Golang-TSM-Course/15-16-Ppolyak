package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(ctx context.Context, r *http.Request, url string) {
	client := http.DefaultClient
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err")
	}

	r = r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Response from %s: %v\n", url, resp.Status)
}

func main() {
	r := &http.Request{}
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	url := []string{
		"https://www.opera.com",
		"https://www.google.com",
	}

	for _, url := range url {
		wg.Add(1)
		go request(ctx, r, url)
		wg.Done()
	}

	wg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	}

}
