package main

import (
	"encoding/json"
	"fmt"
	"github.com/mtailor/gengis/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/mtailor/gengis/datalayer"
	"net/http"
	"strconv"
	"github.com/mtailor/gengis/config"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Gengis salutes you")
	})
	r.HandleFunc("/seasons/{year}", func(writer http.ResponseWriter, request *http.Request) {
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
	const staticPrefix = "/static/"
	r.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir("public"))))
	http.Handle("/", r)
	port := config.Get("PORT")
	fmt.Printf("Running on port %s...", port)
	http.ListenAndServe(":" + port, nil)
}
