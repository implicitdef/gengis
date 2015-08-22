package getseasontimerange
import (
	"time"
	"fmt"
	"github.com/mtailor/gengis/themoviedb/internal"
	"errors"
)


type season struct {
	Episodes []episode
}
type episode struct {
	AirDate string `json:"air_date"`
}

type TimeRange struct {
	start time.Time
	end time.Time
}

func (tr TimeRange) String() string {
	return fmt.Sprintf("{%s, %s}", tr.start.String(), tr.end.String())
}

func parseAirDate(s string) (time.Time, error) {
	const format = "2006-01-02"
	return time.Parse(format, s)
}

func Get(tvShowId string, seasonNumber int) (TimeRange, error) {
	_season := season{}
	_url := fmt.Sprintf("/tv/%s/season/%d", tvShowId, seasonNumber)
	err := internal.DoGetAndJsonUnmarshall(_url, &_season)
	tr := TimeRange{}
	if err != nil {
		return tr, err
	}
	eps := _season.Episodes
	if len(eps) == 0 {
		return tr, errors.New("No episodes for that season")
	}
	start, err := parseAirDate(eps[0].AirDate)
	if err != nil {
		return tr, err
	}
	end, err := parseAirDate(eps[len(eps) - 1].AirDate)
	if err != nil {
		return tr, err
	}
	tr = TimeRange{start, end}
	return tr, nil

}