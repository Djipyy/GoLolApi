package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type Rune struct {
	Count int
	ID    int `json:"runeId"`
}
type RunePages struct {
	Pages []RunePage
	ID    int
}
type RunePage struct {
	Current bool
	Slots   []RuneSlot
	Name    string
	ID      int
}
type RuneSlot struct {
	RuneSlotID int
	RuneID     int
}

func (s *Summoner) GetRunePages() (pages RunePages, e error) {
	response, e := s.API.RequestEndpoint("/lol/platform/v3/runes/by-summoner/"+strconv.Itoa(s.ID), time.Hour)
	if e != nil {
		panic(e)
	}
	pages = RunePages{}
	err := json.Unmarshal(response, &pages)
	if err != nil {
		panic(err)
	}

	return
}
