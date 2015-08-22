package internal

import (
	"github.com/mtailor/gengis/config"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"github.com/mtailor/gengis/myerrors"
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

func checkStatusCode(response *http.Response) error {
	c := response.StatusCode
	if c >= 200 && c < 300 {
		return nil
	}
	msg := fmt.Sprintf("Received %d from TheMovieDb", c)
	if c == 429 {
		return &myerrors.TooManyRequestsError{msg}
	}
	return &myerrors.OtherTheMovieDbError{msg}
}


func DoGetAndJsonUnmarshall(urlCore string, dest interface{}) error {
	_url := BuildUrl(urlCore)
	log.Println(">>> GET", _url)
	response, err := http.Get(_url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if err := checkStatusCode(response); err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}