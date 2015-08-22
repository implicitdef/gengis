package datalayer
import (
	"github.com/mtailor/gengis/domain"
	"github.com/mtailor/gengis/themoviedb"
	"log"
)

type serieBaseInfo struct {
	id int
	title string
}


var seriesTitleById = []serieBaseInfo{
	serieBaseInfo{1399, "Game of Thrones"},
	serieBaseInfo{40008, "Hannibal"},
	serieBaseInfo{46648, "True detective"},
	serieBaseInfo{1402, "Walking dead"},
	serieBaseInfo{62560, "Mr Robot"},
	serieBaseInfo{61664, "Sense 8"},
	serieBaseInfo{60948, "12 Monkeys"},
	serieBaseInfo{60708, "Gotham"},
	serieBaseInfo{1412, "Arrow"},
	serieBaseInfo{62822, "Humans"},
	serieBaseInfo{61889, "Daredevil"},
	serieBaseInfo{62823, "Scream"},
	serieBaseInfo{47640, "The Strain"},
	serieBaseInfo{1413, "American Horror Story"},
	serieBaseInfo{1421, "Modern Family"}}

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

func isSeasonNumberToKeep(seasonNumber int) bool {
	// some series have a some kind of season 0
	// that is irrelevant
	return seasonNumber > 0
}


// This method swallows all error (just log them)
// to be fault tolerant
func fetchAndAppendSeasonDisplay(
	seasonsDisplays []SeasonDisplay,
	year int,
	serie serieBaseInfo,
	seasonNumber int) []SeasonDisplay {
	tr, err := themoviedb.GetSeasonTimeRange(serie.id, seasonNumber)
	// log but keep on
	if err != nil {
		log.Printf("Error %v", err)
		return seasonsDisplays
	}
	seasonDisplay := SeasonDisplay{serie.id, serie.title, seasonNumber, tr}
	if isSeasonDisplayInYear(year, seasonDisplay){
		seasonsDisplays = append(seasonsDisplays, seasonDisplay)
	}
	return seasonsDisplays
}


func GetSeasonsDisplayForYear(year int) []SeasonDisplay {
	seasonsDisplays := []SeasonDisplay{}
	for _, serie := range seriesTitleById {
		seasonsNumbers, err := themoviedb.GetSeasonsNumbers(serie.id)
		if err != nil {
			// log but keep on
			log.Printf("Error %v", err)
		} else {
			for _, seasonNumber := range seasonsNumbers {
				if isSeasonNumberToKeep(seasonNumber) {
					seasonsDisplays = fetchAndAppendSeasonDisplay(seasonsDisplays, year, serie, seasonNumber)
				}
			}
		}
	}
	return seasonsDisplays
}