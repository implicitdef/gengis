package main

import (
	"github.com/mtailor/gengis/datalayer"
	"log"
)



func main() {
	//TODO parallelize calls to the movie db ?
	seasonDisplays := datalayer.GetSeasonsDisplayForYear(2015)
	for _, sd := range seasonDisplays {
		log.Printf("%d %s %02d", sd.SerieId, sd.SerieTitle, sd.SeasonNumber)
	}
}

