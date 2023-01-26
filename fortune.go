package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"strconv"
)

type FortuneResult struct {
	Content string
}

func main() {
	var key int
	var fortune FortuneResult
	flag.IntVar(&key, "type", 4, "Type of fortune")
	flag.Parse()
	result, err := http.Get("http://rzhunemogu.ru/RandJSON.aspx?CType=" + strconv.Itoa(key))
	if err != nil {
		log.Fatal(err)
	}
	if result.StatusCode != http.StatusOK {
		log.Fatalf("Wrong status code: %d", result.StatusCode)
	}
	data, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatalf("err - can not read json - %s", err)
	}
	data = bytes.Replace(data, []byte("\r"), []byte(""), -1)
	data = bytes.Replace(data, []byte("\n"), []byte("\\n"), -1)
	dec := charmap.Windows1251.NewDecoder()
	utf8Bytes, _, err := transform.Bytes(dec, data)
	if err != nil {
		log.Fatalf("err - can not change encoding - %s", err)
	}
	err = json.Unmarshal(utf8Bytes, &fortune)
	if err != nil {
		log.Fatalf("err - can not decode - %s", err)
	}
	content := fortune.Content

	fmt.Println(content)
}
