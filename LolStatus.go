package gololapi

import (
	"encoding/json"
	"time"
)

//ShardStatus Holds the response from the Lol-Status-V3 endpoint.
type ShardStatus struct {
	Name      string
	RegionTag string `json:"region_tag"`
	Hostname  string
	Services  []ShardService
	Slug      string
	Locales   []string
}

//ShardService Data from a service
type ShardService struct {
	Status    string
	Incidents []StatusIncident
	Name      string
	Slug      string
}

//StatusIncident Data from an incident of a service
type StatusIncident struct {
	Active    bool
	CreatedAt string `json:"created_at"`
	ID        float64
	Updates   []StatusMessage
}

//StatusMessage A message of an accident.Contains the translations in Translations
type StatusMessage struct {
	StatusMessage string
	Author        string
	CreatedAt     string `json:"created_at"`
	Translations  []StatusMessageTranslation
	UpdatedAt     string `json:"updated_at"`
	Content       string
	ID            string
}

//StatusMessageTranslation Translation of a Message
type StatusMessageTranslation struct {
	Locale    string
	Content   string
	UpdatedAt string `json:"updated_at"`
}

//GetShardStatus Make a request to the  Lol-Status-V3 endpoint and returns a ShardStatus
func (api *GoLOLAPI) GetShardStatus() (status ShardStatus) {
	response, e := api.RequestEndpoint("/lol/status/v3/shard-data", time.Minute)
	if e != nil {
		panic(e)
	}
	status = ShardStatus{}
	err := json.Unmarshal(response, &status)
	if err != nil {
		panic(err)
	}
	return
}
