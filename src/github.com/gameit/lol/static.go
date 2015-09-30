package lol

/*  --------------------------------------------------------------------------------------------------------
            FOR request section
    --------------------------------------------------------------------------------------------------------  */
const MAX_TOURNAMENT      int = 2
const NEWS_URL            string = "http://na.lolesports.com:80/api/news.json?"
const TOURNAMENT_URL      string = "http://na.lolesports.com:80/api/tournament.json?published=1,0"
const TEAM_URL            string = "http://na.lolesports.com:80/api/tournament/258"

    /* ---------------- About Garena region ---------------- */

const REAL_TIME           string = "https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/summoner/realtime?region=%s&summoner_name=%s"
const SERVER_STATUS       string = "https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/status/servers"
const SUMMONER_STAT       string = "https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/summoner/stats?region=%s&summoner_name=%s";
const SUMMONER_NAME       string = "https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/summoner/get_names?region=%s&summoner_ids=%s";
const VERSION             string = "https://ws.leagueofasia.com:8443/LeagueOfAsia/jaxrs/version/get"
const MMR                 string = "http://%s:%d/observer-mode/rest/consumer/getGameMetaData/%s/%d/1/token"


/*
  Static type that multiple method share usage
*/
type RawTournamentRequest struct {
  TournamentId    int     `json:"tournamentId"`
  TournamentName  string  `json:"tournamentName"`
  NamePublic      string  `json:"namePublic"`
  Contestants     ContestantMap
  IsFinished      bool    `json:"isFinised"`
  DateBegin       string  `json:"dateBegin"`
  DateEnd         string  `json:"dateEnd"`
  NoVods          float64 `json:"noVods"`
  Season          string  `json:"season"`
  Published       bool    `json:"published"`
  Winner          string  `json:"winner"`
}

type TournamentDetail struct {
  NamePublic  string  `json:"namePublic"`
  Contestants TeamMap
}

/*
  Map Type to parse json with static type
*/
type TeamMap        map[string]Team
type TournamentMap  map[string]RawTournamentRequest
type ContestantMap  map[string]Team

/*  --------------------------------------------------------------------------------------------------------
            FOR our structure section
    --------------------------------------------------------------------------------------------------------  */

const MAX_TEAM_CONTESTANT     int = 32

type Tournament struct {
    TournamentId    int     `json:"tournamentId"`
    TournamentName  string  `json:"tournamentName"`
    NamePublic      string  `json:"namePublic"`
    Contestants     []Team
    IsFinished      bool    `json:"isFinised"`
    DateBegin       string  `json:"dateBegin"`
    DateEnd         string  `json:"dateEnd"`
    NoVods          float64 `json:"noVods"`
    Season          string  `json:"season"`
    Published       bool    `json:"published"`
    Winner          string  `json:"winner"`
}

type Team struct {
    Id        string  `json:"id"`
    Name      string  `json:"name"`
    Acronym   string  `json:"acronym"`
}