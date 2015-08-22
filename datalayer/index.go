package datalayer
import (
	"github.com/mtailor/gengis/domain"
	"github.com/mtailor/gengis/themoviedb"
	"github.com/mtailor/gengis/myerrors"
	"log"
)

var seriesTitleById = map[int]string{
	1399:   "Game of Thrones",
	40008:  "Hannibal",
	46648:  "True detective",
	1402:   "Walking dead",
	62560:  "Mr Robot",
	61664:  "Sense 8",
	60948:  "12 Monkeys",
	60708:  "Gotham",
	1412:   "Arrow",
	62822:  "Humans",
	61889:  "Daredevil",
	62823:  "Scream",
	47640:  "The Strain",
	1413:   "American Horror Story",
	1421:   "Modern Family"}

type SeasonDisplay struct {
	SerieId int
	SerieTitle string
	SeasonNumber int
	SeasonTimeRange domain.TimeRange
}


func isSeasonDisplayInYear(year int, sd SeasonDisplay) bool {
	return sd.SeasonTimeRange.Start.Year() <= year &&
		sd.SeasonTimeRange.End.Year() >= year
}

func GetSeasonsDisplayForYear(year int) ([]SeasonDisplay, error) {
	seasonsDisplays := []SeasonDisplay{}
	// TODO parallelize
	for id, title := range seriesTitleById {
		log.Printf("- %s :", title)
		seasonsNumbers, err := themoviedb.GetSeasonsNumbers(id)
		if err != nil {
			return nil, err
		}
		for _, seasonNumber := range seasonsNumbers {
			log.Printf("  - %s S%02d", title, seasonNumber)
			tr, err := themoviedb.GetSeasonTimeRange(id, seasonNumber)
			if err != nil {
				if _, ok := err.(*myerrors.UnprocessableSeasonError); ok {
					log.Printf("    - ignored : " + err.Error())
				} else {
					return nil, err
				}
			} else {
				seasonDisplay := SeasonDisplay{id, title, seasonNumber,	tr}
				if isSeasonDisplayInYear(year, seasonDisplay){
					seasonsDisplays = append(seasonsDisplays, seasonDisplay)
					log.Printf("    - OK !")
				} else {
					log.Printf("    - not in the year")
				}
			}
		}
	}
	return seasonsDisplays, nil
}