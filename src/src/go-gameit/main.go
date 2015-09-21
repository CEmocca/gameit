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
	// "reflect"
	// "regexp"
	"github.com/go-martini/martini"
	// "github.com/jeffail/gabs"
	)

const MAX_TOURNAMENT = 10
const NEWS string = "http://na.lolesports.com:80/api/news.json?"
const TOURNAMENT string = "http://na.lolesports.com:80/api/tournament.json?published=1%2C0"

type Team struct {
  Id    string  `json:"id"`
  Name    string  `json:"name"`
  Acronym   string  `json:"acronym"`
}

type Contest struct {
	Contestants Team `json:"contestants"`
}

// type TeamContestants map[string]Interface{}

type Tournament struct {
  TournamentId    int `json:"tournamentId"`
  TournamentName    string  `json:"tournamentName"`
  NamePublic    string  `json:"namePublic"`
  Contestants    []Contest
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
  m := martini.Classic()
  m.Get("/news", func() string {
    return getNews()
  })
  m.Get("/tournament/:status", func(params martini.Params) string {
    return getTournaments(params["status"])
  })
  m.Run()
}

func jsonToType_getTournament(input []byte) string{
	dec := json.NewDecoder(strings.NewReader(string(input)))
	//fmt.Println(string(input))
	for {
		var info Tour
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info)
	}
	
	// type response []Tournament
	// reg, _ := regexp.Compile("\"tourney\\d+\"")
	// match := reg.FindAllString(string(input), MAX_TOURNAMENT)
	// fmt.Println(match)
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

// const mocca string = "{
//   \\"tourney261\": {
//     \"tournamentName\": \"2015 LPL Regional Qualifiers\",
//     \"namePublic\": \"Regional Qualifiers\",
//     \"contestants\": {
//       \"contestant1\": {
//         \"id\": \"632\",
//         \"name\": \"Invictus Gaming\",
//         \"acronym\": \"iG\"
//       },
//       \"contestant2\": {
//         \"id\": \"1847\",
//         \"name\": \"Edward Gaming\",
//         \"acronym\": \"EDG\"
//       },
//       \"contestant3\": {
//         \"id\": \"3748\",
//         \"name\": \"Snake\",
//         \"acronym\": \"SS\"
//       },
//       \"contestant4\": {
//         \"id\": \"4380\",
//         \"name\": \"Qiao Gu\",
//         \"acronym\": \"QG\"
//       }
//     },
//     \"isFinished\": false,
//     \"dateBegin\": \"2015-09-04T05:00Z\",
//     \"dateEnd\": \"2015-09-04T05:00Z\",
//     \"noVods\": 0,
//     \"season\": \"2015\",
//     \"published\": true,
//     \"winner": \"\"
//   },
//   \"tourney241\": {
//     \"tournamentName\": \"2015 ACM Regional Qualifiers\",
//     \"namePublic\": \"Regional Qualifiers\",
//     \"contestants\": {
//       \"contestant1\": {
//         \"id\": \"632\",
//         \"name\": \"Invictus Gaming\",
//         \"acronym\": \"iG\"
//       },
//       \"contestant2\": {
//         \"id\": \"1847\",
//         \"name\": \"Edward Gaming\",
//         \"acronym\": \"EDG\"
//       },
//       \"contestant3\": {
//         \"id\": \"3748\",
//         \"name\": \"Snake\",
//         \"acronym\": \"SS\"
//       },
//       \"contestant4\": {
//         \"id\": \"4380\",
//         \"name\": \"Qiao Gu\",
//         \"acronym\": \"QG\"
//       }
//     },
//     \"isFinished\": false,
//     \"dateBegin\": \"2015-09-04T05:00Z\",
//     \"dateEnd\": \"2015-09-04T05:00Z\",
//     \"noVods\": 0,
//     \"season\": \"2015\",
//     \"published\": true,
//     \"winner\": \"\"
//   }
// }"