package main

import (
	"fmt"
	"log"
)

var link = "https://adventofcode.com/2017/day/1/input"

func main() {
	data, err := getDataFromHTTP(link, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
