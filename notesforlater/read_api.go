package notesforlater
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


///// How to read an api and unmarshal the json


func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

type post struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

func call() {
	response, err := http.Get("http://........")
	panicIfErr(err)
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	target := []post{}
	json.Unmarshal(bytes, &target)
	fmt.Println(target)

}
