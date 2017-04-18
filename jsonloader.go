package jsonloader

import (
	"io"
	"bufio"
	"io/ioutil"
	"fmt"
	"log"
	"os"
)

type Language struct {
	Key, Value, Description string
}

func LoadJson(fn string) string, error {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	defer f.Close()

	rdr := bufio.NewReader(f)
	
	jsonStr, err := bufio.ReadAll(rdr)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return jsonStr, nil
}

