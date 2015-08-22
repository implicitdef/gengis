package getseasonnumbers
import (
	"fmt"
	"github.com/mtailor/gengis/themoviedb/internal"
)


type tvshow struct {
	Seasons []season
}
type season struct {
	SeasonNumber int `json:"season_number"`
}

func Get(tvShowId int) ([]int, error) {
	tvs := tvshow{}
	_url := fmt.Sprintf("/tv/%d", tvShowId)
	err := internal.DoGetAndJsonUnmarshall(_url, &tvs)
	if err != nil {
		return nil, err
	}
	numbers := make([]int, len(tvs.Seasons))
	for i, season := range tvs.Seasons {
		numbers[i] = season.SeasonNumber
	}
	return numbers, nil
}