package main

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"io"
	"log"
	"strings"
	"container/list"
	// "reflect"
	// "regexp"
	// "github.com/go-martini/martini"
	// "github.com/jeffail/gabs"
	)

const MAX_TOURNAMENT = 10
const NEWS string = "http://na.lolesports.com:80/api/news.json?"
const TOURNAMENT string = "http://na.lolesports.com:80/api/tournament.json?published=1%2C0"
const TEAM string = "http://na.lolesports.com:80/api/tournament/258"

type Team struct {
  	Id    string  `json:"id"`
  	Name    string  `json:"name"`
  	Acronym   string  `json:"acronym"`
}

type TeamMap map[string]Team

type TournamentDetail struct {
	NamePublic string `json:"namePublic"`
	Contestants TeamMap
}

type Contest struct {
	Contestants Team
}

type ContestMap map[string]Team

type Tournament struct {
  TournamentId    int `json:"tournamentId"`
  TournamentName    string  `json:"tournamentName"`
  NamePublic    string  `json:"namePublic"`
  Contestants    ContestMap
  IsFinished    bool  `json:"isFinised"`
  DateBegin     string  `json:"dateBegin"`
  DateEnd     string  `json:"dateEnd"`
  NoVods      float64 `json:"noVods"`
  Season      string  `json:"season"`
  Published   bool  `json:"published"`
  Winner      string  `json:"winner"`
}

type Tour map[string]Tournament

func main() {
	// getTeam("258")
  getTournaments("true")
}

func jsonToType_getTournament(input []byte) string{
	dec := json.NewDecoder(strings.NewReader(string(input)))
	//fmt.Println(string(input))
	allTour := list.New()
	for {
		var info Tour
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info["tourney247"].Contestants["contestant1"].Name)
		allTour.PushBack(info)
	}

	// fmt.Println(allTour.Front())	
	// type response []Tournament
	// reg, _ := regexp.Compile("\"tourney\\d+\"")
	// match := reg.FindAllString(string(input), MAX_TOURNAMENT)
	// fmt.Println(match)
	return "true"
}

func jsonToType_getTeam(input []byte) string{
	dec := json.NewDecoder(strings.NewReader(string(input)))
	// fmt.Println(string(input))
	for {
		var info TournamentDetail
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info)
	}
	return "true"
}

func getTournaments(isFinished string) string {
	var buffer bytes.Buffer
	buffer.WriteString(TOURNAMENT)
	resp, err := http.Get(buffer.String())
	if err != nil {
		return "Invalid url"
	}
	defer resp.Body.Close()
	if(resp.StatusCode == 200) {
		body, _ := ioutil.ReadAll(resp.Body)
		jsonToType_getTournament(body)
		// res := &Tournament{}
		// fmt.Println(json.Unmarshal(body, &res))
		return string(body)
	} else {
		return "Internal error"
	}
}

func getTeam(tournamentId string) string {
	var buffer bytes.Buffer
	buffer.WriteString(TEAM)
	resp, err := http.Get(buffer.String())
	if err != nil {
		return "Invalid url"
	}
	defer resp.Body.Close()
	if(resp.StatusCode == 200) {
		body, _ := ioutil.ReadAll(resp.Body)
		jsonToType_getTeam(body)
		// res := &Tournament{}
		// fmt.Println(json.Unmarshal(body, &res))
		return string(body)
	} else {
		return "Internal error"
	}
}

func getNews() string {
	var buffer bytes.Buffer
	buffer.WriteString(NEWS)
	buffer.WriteString("limit=5")
	resp, err := http.Get(buffer.String())
	if err != nil {
		return "Invalid url"
	}
	defer resp.Body.Close()
	if(resp.StatusCode == 200) {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("body")
		return string(body)
	} else {
		return "Internal error"
	}
	
}

const mocca string = "{\"tourney261\": {\"tournamentName\": \"2015 LPL Regional Qualifiers\",\"namePublic\": \"Regional Qualifiers\",\"contestants\": {\"contestant1\": {\"id\": \"632\",\"name\": \"Invictus Gaming\",\"acronym\": \"iG\"},\"contestant2\": {\"id\": \"1847\",\"name\": \"Edward Gaming\",\"acronym\": \"EDG\"},\"contestant3\": {\"id\": \"3748\",\"name\": \"Snake\",\"acronym\": \"SS\"},\"contestant4\": {\"id\": \"4380\",\"name\": \"Qiao Gu\",\"acronym\": \"QG\"}},\"isFinished\": false,\"dateBegin\": \"2015-09-04T05:00Z\",\"dateEnd\": \"2015-09-04T05:00Z\",\"noVods\": 0,\"season\": \"2015\",\"published\": true,\"winner\": \"\"}}"