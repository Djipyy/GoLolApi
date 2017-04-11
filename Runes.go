package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//Rune Holds the data of a rune .To get more information about the rune, you need to query the Static Data.
type Rune struct {
	Count int
	ID    int `json:"runeId"`
}

//RunePages Holds the runes pages of a summoner and his summoner ID.
type RunePages struct {
	Pages []RunePage
	ID    int
}

//RunePage Holds the data of a rune page.The runes are in Slots.
type RunePage struct {
	Current bool
	Slots   []RuneSlot
	Name    string
	ID      int
}

//RuneSlot Holds the data of a rune slot.To get more information about the rune, you need to query the Static Data.
type RuneSlot struct {
	RuneSlotID int
	ID         int `json:"runeId"`
}

//GetRunePages Makes a request to the Runes-V3 endpoint and returns a RunePages struct
func (s *Summoner) GetRunePages() (pages RunePages) {
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
