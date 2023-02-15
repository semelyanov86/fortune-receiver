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

type LiveGetWebRequest struct {
}

func (g LiveGetWebRequest) FetchBytes(typeFort int) ([]byte, error) {
	url := "http://rzhunemogu.ru/RandJSON.aspx?CType=" + strconv.Itoa(typeFort)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "Conky-Fortune-Ext")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Wrong status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func getFortune(getWebRequest GetWebRequest, typeFort int) (FortuneResult, error) {
	var fortune FortuneResult
	body, err := getWebRequest.FetchBytes(typeFort)
	if err != nil {
		return fortune, err
	}
	data := bytes.Replace(body, []byte("\r"), []byte(""), -1)
	data = bytes.Replace(data, []byte("\n"), []byte("\\n"), -1)
	dec := charmap.Windows1251.NewDecoder()
	utf8Bytes, _, err := transform.Bytes(dec, data)
	if err != nil {
		return fortune, err
	}
	err = json.Unmarshal(utf8Bytes, &fortune)
	if err != nil {
		return fortune, err
	}
	return fortune, err
}

func main() {
	var key int
	var maxLength int
	flag.IntVar(&key, "type", 4, "Type of fortune")
	flag.IntVar(&maxLength, "length", 85, "Max length of string")
	flag.Parse()
	liveClient := LiveGetWebRequest{}
	content, err := getFortune(liveClient, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content.splitString(maxLength))
}
