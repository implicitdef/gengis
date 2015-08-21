package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}



func bindHelloEndpoint() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Gengis salutes you")
	})
}

func bindJsonSampleEndpoint(){
	http.HandleFunc("/jsonSample", func(writer http.ResponseWriter, request *http.Request){
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.Write(getJsonBytes())
	})
}

func serve(){
	http.ListenAndServe(":8080", nil)
}

func getJsonBytes() []byte {
	data := []todo{
		todo{Name: "Write presentation"},
		todo{Name: "Host meetup"},
	}
	jsonBytes, _ := json.Marshal(data)
	return jsonBytes
}


func main() {


	fmt.Println("We got the following json", string(getJsonBytes()))


	bindHelloEndpoint()
	bindJsonSampleEndpoint()
	fmt.Println("Listening...")
	serve()
}
