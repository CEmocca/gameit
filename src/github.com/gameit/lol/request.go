package lol

import (
	"io/ioutil"
	"bytes"
	"net/http"
	)

func request(url string) string {
	var buffer bytes.Buffer
	buffer.WriteString(url)
	resp, err := http.Get(buffer.String())
	if err != nil {
		return "Invalid url"
	}
	defer resp.Body.Close()
	if(resp.StatusCode == 200) {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body)
	} else {
		return "Internal error"
	}
}