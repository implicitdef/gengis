package main

import (
	"fmt"
	"github.com/mtailor/gengis/datalayer"
)



func main() {
	//TODO fault tolerance on bad error codes from themoviedb
	//TODO find a way to be more tolerant with themoviedb limits : caching ? throttling ? retry ?
	//TODO parallelize calls to the movie db ?
	fmt.Println(datalayer.GetSeasonsDisplayForYear(2015))
}

