package main

import "net/http"

func main() {
	route := http.NewServeMux()

	route.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {

	})

	server := http.Serve({
		
	})
}
