package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var link = "https://adventofcode.com/2017/day/1/input"

/*
After trying for many hours to get the input via http request I have come to the realization
that this was won't work. Using these functions the resp comes back blank. Using curl,
the resp come back along the lines of: " Pleasee login, the input is different for every user."
Looking at the page source reveals only the input, so it must be ajax or something.
Since I am unsure of how to continue, I am just going to skip this part, and hard-code
the input to the file.

func main() {
	var respdata []byte
	err := getDataFromHTTP(link, respdata)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(respdata))

	if err := writeToFile(deriveFileNameFromLink(link), respdata); err != nil {
		log.Println(err)
	}
}
*/

var input, _ = ioutil.ReadFile("2017day1input.txt")

func main() {
}
