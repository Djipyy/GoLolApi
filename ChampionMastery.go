package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type ChampionMasteries struct {
	a []ChampionMastery `json:""`
}
type ChampionMastery struct {
	ChestGranted         bool
	Level                int `json:"championLevel"`
	Points               int `json:"championPoints"`
	ID                   int `json:"championID"`
	PID                  int `json:"playerId"`
	PointsUntilNextLevel int `json:"championPointsUntilNextLevel"`
	PointsSinceLastLevel int `json:"championPointsSinceLastLevel"`
	LastPlayTime         int `json:"lastPlayTime"`
}

func (s *Summoner) GetChampionMasteries() (masteries []ChampionMastery) {
	response, e := s.API.RequestEndpoint("/lol/champion-mastery/v3/champion-masteries/by-summoner/"+strconv.Itoa(s.ID), time.Hour)
	if e != nil {
		panic(e)
	}
	masteries = []ChampionMastery{}
	err := json.Unmarshal(response, &masteries)
	if err != nil {
		panic(err)
	}

	return
}
func (s *Summoner) GetMasteryOfChampion(championID int) (mastery ChampionMastery) {
	response, e := s.API.RequestEndpoint("/lol/champion-mastery/v3/champion-masteries/by-summoner/"+strconv.Itoa(s.ID)+"/by-champion/"+strconv.Itoa(championID), time.Hour)
	if e != nil {
		panic(e)
	}
	mastery = ChampionMastery{}
	err := json.Unmarshal(response, &mastery)
	if err != nil {
		panic(err)
	}

	return
}
func (s *Summoner) GetTotalChampionMastery() (score int) {
	response, e := s.API.RequestEndpoint("/lol/champion-mastery/v3/scores/by-summoner/"+strconv.Itoa(s.ID), time.Hour)
	if e != nil {
		panic(e)
	}
	scoreString := string(response)
	score, err := strconv.Atoi(scoreString)
	if err != nil {
		panic(err)
	}

	return
}
