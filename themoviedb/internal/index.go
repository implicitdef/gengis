package internal

import (
	"github.com/mtailor/gengis/config"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"errors"
	"fmt"
)

func BuildUrl(mainPath string) string {
	u, err := url.Parse("https://api.themoviedb.org/3")
	if err != nil {
		panic(err)
	}
	u.Path += mainPath
	params := url.Values{}
	params.Set("api_key", config.GetTheMovieDbApiKey())
	u.RawQuery = params.Encode()
	return u.String()
}

func isGoodStatusCode(response *http.Response) bool {
	return response.StatusCode >= 200 &&
		response.StatusCode < 300
}


func DoGetAndJsonUnmarshall(urlCore string, dest interface{}) error {
	_url := BuildUrl(urlCore)
	log.Println(">>> GET", _url)
	response, err := http.Get(_url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if ! isGoodStatusCode(response) {
		return errors.New(fmt.Sprintf("Received status code %d", response.StatusCode))
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}