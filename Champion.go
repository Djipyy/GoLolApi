package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//ChampionList Holds a list of Champion in Champions
type ChampionList struct {
	Champions []Champion
}

//Champion Hold a champion data that was retrived by the Champion-V3 endpoint
type Champion struct {
	ID                int
	FreeToPlay        bool
	Active            bool
	BotMMEnabled      bool
	BotEnabled        bool
	RankedPlayEnabled bool
}

//GetChampîons Make a request to the Champion-V3 endpoint and returns a ChampionList
func (api *GoLOLAPI) GetChampîons(onlyFreeToPlay bool) (champions ChampionList) {
	var endpoint string
	if onlyFreeToPlay {
		endpoint = "/lol/platform/v3/champions?freeToPlay=true"
	} else {
		endpoint = "/lol/platform/v3/champions?freeToPlay=true"

	}
	response, e := api.RequestEndpoint(endpoint, time.Hour)

	if e != nil {
		panic(e)
	}
	champions = ChampionList{}
	err := json.Unmarshal(response, &champions)
	if err != nil {
		panic(err)
	}
	return
}

//GetChampîon Make a request to the Champion-V3 endpoint and returns a Champion
func (api *GoLOLAPI) GetChampîon(ID int) (champion Champion) {
	IDString := strconv.Itoa(ID)
	response, e := api.RequestEndpoint("/lol/platform/v3/champions/"+IDString, time.Hour)
	if e != nil {
		panic(e)
	}
	champion = Champion{}
	err := json.Unmarshal(response, &champion)
	if err != nil {
		panic(err)
	}
	return
}
