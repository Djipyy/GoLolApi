package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//League  This object contains league information.
type League struct {
	Queue         string
	Tier          string
	ParticipantID string
	Name          string
	Entries       []LeagueEntry
}

//LeagueEntry This object contains league participant information representing a summoner or team.
type LeagueEntry struct {
	IsFreshBlood     bool
	Division         string
	Playstyle        string
	MiniSeries       MiniSeries
	Wins             int
	Losses           int
	PlayerOrTeamID   string
	PlayerOrTeamName string
	IsHotStreak      bool
	IsInactive       bool
	IsVeteran        bool
	LeaguePoints     int
}

//MiniSeries This object contains mini series information.
type MiniSeries struct {
	Progress string
	Losses   int
	Wins     int
	Target   int
}

//GetMasterLeague Get master tier leagues.
func (api *GoLOLAPI) GetMasterLeague(queue string) (league League) {
	options := map[string]string{"type": queue}
	uri, hasParameters := GetEndpointURI("/api/lol/"+api.Region.Name+"/v2.5/league/master", options)
	response, e := api.RequestLegacyEndpoint(uri, time.Hour, hasParameters)
	if e != nil {
		panic(e)
	}
	league = League{}
	err := json.Unmarshal(response, &league)
	if err != nil {
		panic(err)
	}
	return
}

//GetChallengerLeague Get challenger tier leagues.
func (api *GoLOLAPI) GetChallengerLeague(queue string) (league League) {
	options := map[string]string{"type": queue}
	uri, hasParameters := GetEndpointURI("/api/lol/"+api.Region.Name+"/v2.5/league/challenger", options)
	response, e := api.RequestLegacyEndpoint(uri, time.Hour, hasParameters)
	if e != nil {
		panic(e)
	}
	league = League{}
	err := json.Unmarshal(response, &league)
	if err != nil {
		panic(err)
	}
	return
}

//GetLeagues Get leagues mapped by summoner ID for a given list of summoner ID.
//Returns all leagues for specified summoners and his teams. Entries for each requested participant (i.e., each summoner and related teams) will be included in the returned leagues data, whether or not the participant is inactive. However, no entries for other inactive summoners or teams in the leagues will be included.
func (s *Summoner) GetLeagues() (leagues map[string][]League) {
	response, e := s.API.RequestLegacyEndpoint("/api/lol/"+s.API.Region.Name+"/v2.5/league/by-summoner/"+strconv.FormatFloat(s.AccountID, 'f', -1, 64), time.Hour, false)
	if e != nil {
		panic(e)
	}
	leagues = map[string][]League{}
	err := json.Unmarshal(response, &leagues)
	if err != nil {
		panic(err)
	}
	return
}

//GetLeagueEntries Get league entries mapped by summoner ID for a given list of summoner ID.
//Returns all league entries for specified summoners and summoners&apos; teams.
func (s *Summoner) GetLeagueEntries() (leagues map[string][]League) {
	response, e := s.API.RequestLegacyEndpoint("/api/lol/"+s.API.Region.Name+"/v2.5/league/by-summoner/"+strconv.FormatFloat(s.AccountID, 'f', -1, 64)+"/entry", time.Hour, false)
	if e != nil {
		panic(e)
	}
	leagues = map[string][]League{}
	err := json.Unmarshal(response, &leagues)
	if err != nil {
		panic(err)
	}
	return
}
