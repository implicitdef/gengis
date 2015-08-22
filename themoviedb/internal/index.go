package internal

import (
	"github.com/mtailor/gengis/config"
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
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

func DoGetAndJsonUnmarshall(urlCore string, dest interface{}) error {
	_url := BuildUrl(urlCore)
	fmt.Println(">>> GET", _url)
	response, err := http.Get(_url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}