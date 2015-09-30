package main

import (
	"io/ioutil"
    "fmt"
    "net/http"
    "crypto/tls"
)

func main() {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/summoner/stats?region=TH&summoner_name=DuckReturn")
    if err != nil {
        fmt.Println(err)
    }
   body, _ := ioutil.ReadAll(resp.Body)
   fmt.Println(string(body))
}