package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//Mastery Holds the ID and Rank from a mastery.Is obtained via the Masteries-V3 endpoint
type Mastery struct {
	ID   int `json:"masteryId"`
	Rank int
}

//MasteryPages Holds the mastery pages of a summoner in Pages and his summoner ID in ID.
type MasteryPages struct {
	Pages []MasteryPage
	ID    int
}

//MasteryPage Holds the data of a mastery page.
type MasteryPage struct {
	Current   bool
	Masteries []Mastery
	Name      string
	ID        float64
}

//GetMasteryPages Makes a request to the Masteries-V3 endpoint and returns a MasteryPages struct
func (s *Summoner) GetMasteryPages() (pages MasteryPages) {
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
