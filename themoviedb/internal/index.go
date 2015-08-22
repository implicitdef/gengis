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
	"time"
	"github.com/mtailor/gengis/vendor/_nuts/github.com/cenkalti/backoff"
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


func getBackOff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 10 * time.Second
	return b
}

func DoGetAndJsonUnmarshall(urlCore string, dest interface{}) error {
	_url := BuildUrl(urlCore)
	operation := func() error {
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
	return backoff.Retry(operation, getBackOff())
}



