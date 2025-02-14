package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const targetUrl = "http://localhost:8080"

func main() {
	target, err := url.Parse(targetUrl)
	if err != nil {
		fmt.Println("Error while parsing the url")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Proxying request : ", r.Method, r.URL.Path)
		proxy.ServeHTTP(w, r)
	})

	port := ":8080"
	fmt.Println("Reverse Proxy running on http://localhost" + port)

}
