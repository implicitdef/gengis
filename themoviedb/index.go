package themoviedb
import (
	"github.com/mtailor/gengis/themoviedb/internal/getseasonnumbers"
	"github.com/mtailor/gengis/themoviedb/internal/getseasontimerange"
	"github.com/mtailor/gengis/domain"
)




func GetSeasonsNumbers(tvShowId int) ([]int, error) {
	return getseasonnumbers.Get(tvShowId)
}

func GetSeasonTimeRange(tvShowId int, seasonNumber int) (domain.TimeRange, error) {
	return getseasontimerange.Get(tvShowId, seasonNumber)
}