package main

import (
	"net/http"
	"fmt"
)

func main() {

	publicEndpoint := "/static"
	staticFolderPath := "public"

	fs := http.FileServer(http.Dir(staticFolderPath))
	http.Handle(publicEndpoint + "/", http.StripPrefix(publicEndpoint + "/", fs))

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello KKKKK.K.")
	})

	http.ListenAndServe(":8080", nil)
}
