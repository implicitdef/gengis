package main

import (
	"fmt"
	"github.com/mtailor/gengis/themoviedb"
)



func main() {

	fmt.Println(themoviedb.GetSeasonsNumbers("1399"))
	fmt.Println(themoviedb.GetSeasonTimeRange("1399", 5))
}

