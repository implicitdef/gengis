package main

import (
	"fmt"
	"github.com/mtailor/gengis/themoviedb/getseasonnumbers"
	"github.com/mtailor/gengis/themoviedb/getseasontimerange"
)



func main() {
	//fmt.Println(themoviedb.GetSeasonNumbers("1399"))
	fmt.Println(getseasonnumbers.Get("1399"))
	fmt.Println(getseasontimerange.Get("1399", 5))
}

