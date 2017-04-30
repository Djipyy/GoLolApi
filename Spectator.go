package gololapi

import (
	"encoding/json"
	"strconv"
)

type CurrentGameInfo struct {
	GameID            int
	GameStartTime     int
	PlatformID        string
	GameMode          string
	MapID             int
	GameType          string
	BannedChampions   []BannedChampion
	Observers         Observer
	Participants      []CurrentGameParticipant
	GameLength        int
	GameQueueConfigID int
}
type FeaturedGames struct {
	ClientRefreshInterval int
	Games                 []FeaturedGameInfo `json:"gameList"`
}
type FeaturedGameInfo struct {
	GameID            int
	GameStartTime     int
	PlatformID        string
	GameMode          string
	MapID             int
	GameType          string
	BannedChampions   []BannedChampion
	Observers         Observer
	Participants      []CurrentGameParticipant
	GameLength        int
	GameQueueConfigID int
}
type Observer struct {
	EncryptionKey string
}
type BannedChampion struct {
	PickTurn   int
	ChampionID int
	TeamID     int
}
type CurrentGameParticipant struct {
	ProfileIconID    int
	ChampionID       int
	SummonerName     string
	Bot              bool
	SummonerSpell1ID int `json:"spell1Id"`
	SummonerSpell2ID int `json:"spell2Id"`
	TeamID           int
	SummonerID       int
	Runes            []Rune
	Masteries        []Mastery
}

func (s *Summoner) GetCurrentGame() (game CurrentGameInfo, e error) {
	response, e := s.API.RequestEndpoint("/lol/spectator/v3/active-games/by-summoner/"+strconv.FormatFloat(s.ID, 'f', -1, 64), 0)
	if e != nil {
		panic(e)
	}
	game = CurrentGameInfo{}
	err := json.Unmarshal(response, &game)
	if err != nil {
		panic(err)
	}
	return
}
func (api *GoLOLAPI) GetFeaturedGames() (games FeaturedGames, e error) {
	response, e := api.RequestEndpoint("/lol/spectator/v3/featured-games", 0)
	if e != nil {
		panic(e)
	}
	games = FeaturedGames{}
	err := json.Unmarshal(response, &games)
	if err != nil {
		panic(err)
	}

	return
}
