package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getDataFromHTTP(link string, to ...io.Writer) error {
	var tmpbuf bytes.Buffer
	cwd, _ := os.Getwd()
	file, err := os.Create(cwd + link + ".txt")
	if err != nil {
		log.Printf("failed to create file %s, error: %s \n", link, err)
		return err
	}
	defer file.Close()

	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	if read, _ := tmpbuf.Write(data); read != len(data) {
		copyerr := fmt.Errorf("ineffectual copy to buffer; copied: %d out of %d bytes", read, len(data))
		log.Print(copyerr)

		return copyerr
	}

	for i := 0; i < len(to); i++ {
		_, err := to[i].Write(tmpbuf.Bytes())
		if err != nil {
			log.Println(err)
			return err
		}

	}

	if _, err := file.Write(tmpbuf.Bytes()); err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("data from HTTP request written to file: %s", cwd+link+".txt")
	return nil
}
