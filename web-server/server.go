package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Server: hello handler started ✅")
	defer fmt.Println("server: hello handler has ended ✅")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err, " ❌")
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Println(w, name, h)
// 		}
// 	}
// }

func main() {
	http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
