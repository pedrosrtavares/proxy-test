package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.Handle("/push", Handle())
	http.ListenAndServe(":3000", nil)
}

func Handle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		header := request.Header
		fmt.Println(header)
		host := request.URL.Host
		fmt.Println(host)
		path := request.URL.Path
		fmt.Println(path)
		body := request.Body
		defer body.Close()

		req1, err := http.NewRequest("GET", "http://backstage.squair.com.br", body)
		if err != nil {
			panic(err)
		}

		req2, err := http.NewRequest("GET", "http://backstage.squair.io", body)
		if err != nil {
			panic(err)
		}

		res1, err := client.Do(req1)
		if err != nil {
			panic(err)
		}
		fmt.Println(res1.StatusCode)

		res2, err := client.Do(req2)
		if err != nil {
			panic(err)
		}
		fmt.Println(res2.StatusCode)

		writer.WriteHeader(200)
	}
}
