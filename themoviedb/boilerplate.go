package themoviedb

import (
	"net/url"
	"github.com/mtailor/gengis/config"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)


func buildUrl(mainPath string) string {
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

func doGetAndJsonUnmarshall(urlCore string, dest interface{}) error {
	_url := buildUrl(urlCore)
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