package getseasontimerange
import (
	"time"
	"fmt"
	"github.com/mtailor/gengis/themoviedb/internal"
	"github.com/mtailor/gengis/domain"
	"github.com/mtailor/gengis/myerrors"
)


type season struct {
	Episodes []episode
}
type episode struct {
	AirDate string `json:"air_date"`
}


func parseAirDate(s string) (time.Time, error) {
	const format = "2006-01-02"
	return time.Parse(format, s)
}

func Get(tvShowId int, seasonNumber int) (domain.TimeRange, error) {
	_season := season{}
	_url := fmt.Sprintf("/tv/%d/season/%d", tvShowId, seasonNumber)
	err := internal.DoGetAndJsonUnmarshall(_url, &_season)
	tr := domain.TimeRange{}
	if err != nil {
		return tr, err
	}
	eps := _season.Episodes
	if len(eps) == 0 {
		return tr, &myerrors.UnprocessableSeasonError{"No episodes for this season"}
	}
	start, err := parseAirDate(eps[0].AirDate)
	if err != nil {
		return tr, &myerrors.UnprocessableSeasonError{"Unparsable air date : " + eps[0].AirDate}
	}
	end, err := parseAirDate(eps[len(eps) - 1].AirDate)
	if err != nil {
		return tr, &myerrors.UnprocessableSeasonError{"Unparsable air date : " + eps[len(eps) - 1].AirDate}
	}
	tr = domain.TimeRange{start, end}
	return tr, nil

}