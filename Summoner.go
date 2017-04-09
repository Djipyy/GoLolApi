package gololapi

import (
	"encoding/json"
	"strconv"
	"time"
)

//Summoner Contains a SummonerDto Object
type Summoner struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
	Summonerlevel int    `json:"summonerLevel"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
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
func (api *GoLOLAPI) GetSummonerByName(name string) (s Summoner) {
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-name/"+name, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
func (api *GoLOLAPI) GetSummonerByID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/"+IDString, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
func (api *GoLOLAPI) GetSummonerByAccountID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-account/"+IDString, time.Hour)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
