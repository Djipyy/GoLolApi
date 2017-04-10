package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type Mastery struct {
	ID   int `json:"masteryId"`
	Rank int
}

type MasteryPages struct {
	Pages []MasteryPage
	ID    int
}
type MasteryPage struct {
	Current   bool
	Masteries []Mastery
	Name      string
	ID        float64
}

func (s *Summoner) GetMasteryPages() (pages MasteryPages, e error) {
	response, e := s.API.RequestEndpoint("/lol/platform/v3/masteries/by-summoner/"+strconv.Itoa(s.ID), time.Hour)
	if e != nil {
		panic(e)
	}
	pages = MasteryPages{}
	err := json.Unmarshal(response, &pages)
	if err != nil {
		panic(err)
	}

	return
}
