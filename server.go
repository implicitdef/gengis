package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/mtailor/gengis/vendor/_nuts/github.com/gorilla/mux"
	"strconv"
	"github.com/mtailor/gengis/datalayer"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Gengis salutes you")
	})
	r.HandleFunc("/seasons/{year}", func(writer http.ResponseWriter, request *http.Request){
		yearStr := mux.Vars(request)["year"]
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			fmt.Printf("Received invalid year %s", year)
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("Invalid year"))
		} else {
			data := datalayer.GetSeasonsDisplayForYear(year)
			bytes, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("Internal server error"))
			} else {
				writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
				writer.Write(bytes)
			}
		}

	})
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
