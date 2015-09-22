package lol

import (
	// "net/http"
	"fmt"
	// "bytes"
	// "io/ioutil"
	"encoding/json"
	"io"
	"log"
	"strings"
	"regexp"
	)

func parseTournament(input string) []Tournament{
	dec := json.NewDecoder(strings.NewReader(input))
	var info TournamentMap
	for {
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	tournamentPattern, _ := regexp.Compile("tourney\\d+")
	matchTournament := tournamentPattern.FindAllString(string(input), MAX_TOURNAMENT)
	teamPattern, _ := regexp.Compile("contestant\\d+")

	tournaments := make([]Tournament, 0)
	for _, tournamentName := range matchTournament {
		var t = info[tournamentName]
		matchTeam := teamPattern.FindAllString(fmt.Sprintf(`"%s"`, t.Contestants), MAX_TEAM_CONTESTANT)
		contestants := make([]Team, 0)
		for _, teamName := range matchTeam {
			var c = t.Contestants[teamName]
			contestant := Team{Id: c.Id, Name: c.Name, Acronym: c.Acronym}
			contestants = append(contestants, contestant)
		}
		tournamentDetail := Tournament{TournamentId: t.TournamentId, TournamentName : t.TournamentName, NamePublic: t.NamePublic, IsFinished: t.IsFinished, DateBegin: t.DateBegin, DateEnd: t.DateEnd, NoVods: t.NoVods, Season: t.Season, Published: t.Published, Winner: t.Winner, Contestants: contestants}
		tournaments = append(tournaments, tournamentDetail)
	}
	return tournaments
}

func GetTournaments() []Tournament {
	response := request(TOURNAMENT_URL)
	ret := parseTournament(response)
	return ret
}