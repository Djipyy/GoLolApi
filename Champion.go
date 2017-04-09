package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

type ChampionList struct {
	Champions []Champion
}
type Champion struct {
	ID                int
	FreeToPlay        bool
	Active            bool
	BotMMEnabled      bool
	BotEnabled        bool
	RankedPlayEnabled bool
}

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
