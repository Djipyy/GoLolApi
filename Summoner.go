package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//Summoner Contains a SummonerDto Object
type Summoner struct {
	Name          string  `json:"name"`
	ID            float64 `json:"id"`
	AccountID     float64 `json:"accountId"`
	Summonerlevel int     `json:"summonerLevel"`
	ProfileIconID int     `json:"profileIconId"`
	RevisionDate  int     `json:"revisionDate"`
	Region        *Region
	API           *GoLOLAPI
}

func createSummonerFromResponse(response []byte, api *GoLOLAPI) (s Summoner) {
	s = Summoner{}
	err := json.Unmarshal(response, &s)
	if err != nil {
		panic(err)
	}
	s.Region = &api.Region
	s.API = api
	return

}

//GetSummonerByName Get a summoner by summoner name.
func (api *GoLOLAPI) GetSummonerByName(name string) (s Summoner) {
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-name/"+name, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}

//GetSummonerByID Get a summoner by summoner ID.
func (api *GoLOLAPI) GetSummonerByID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/"+IDString, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}

//GetSummonerByAccountID Get a summoner by account ID.
func (api *GoLOLAPI) GetSummonerByAccountID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-account/"+IDString, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
