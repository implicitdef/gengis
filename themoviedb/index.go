package themoviedb
import (
	"github.com/mtailor/gengis/themoviedb/internal/getseasonnumbers"
	"github.com/mtailor/gengis/themoviedb/internal/getseasontimerange"
)


func GetSeasonsNumbers(tvShowId string) ([]int, error) {
	return getseasonnumbers.Get(tvShowId)
}

func GetSeasonTimeRange(tvShowId string, seasonNumber int) (getseasontimerange.TimeRange, error) {
	return getseasontimerange.Get(tvShowId, seasonNumber)
}