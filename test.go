package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

//EUW Values for the EUW Region
var EUW = Region{Name: "EUW", PlatformID: "EUW1", Host: "https://euw1.api.riotgames.com"}

//Region Struct containing information about a region
type Region struct {
	Name       string
	PlatformID string
	Host       string
}

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
type CurrentGameInfo struct {
	GameID            int
	GameStartTime     int
	PlatformID        string
	GameMode          string
	MapID             int
	GameType          string
	BannedChampions   []BannedChampion
	Observers         Observer
	Participants      []CurrentGameParticipant
	GameLength        int
	GameQueueConfigID int
}
type FeaturedGames struct {
	ClientRefreshInterval int
	Games                 []FeaturedGameInfo `json:"gameList"`
}
type FeaturedGameInfo struct {
	GameID            int
	GameStartTime     int
	PlatformID        string
	GameMode          string
	MapID             int
	GameType          string
	BannedChampions   []BannedChampion
	Observers         Observer
	Participants      []CurrentGameParticipant
	GameLength        int
	GameQueueConfigID int
}
type Observer struct {
	EncryptionKey string
}
type BannedChampion struct {
	PickTurn   int
	ChampionID int
	TeamID     int
}
type CurrentGameParticipant struct {
	ProfileIconID    int
	ChampionID       int
	SummonerName     string
	Bot              bool
	SummonerSpell1ID int `json:"spell1Id"`
	SummonerSpell2ID int `json:"spell2Id"`
	TeamID           int
	SummonerID       int
	Runes            []Rune
	Masteries        []Mastery
}
type Rune struct {
	Count int
	ID    int `json:"runeId"`
}
type Mastery struct {
	ID   int `json:"masteryId"`
	Rank int
}
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

func MinifySummonerName(name string) (r string) {
	r = strings.Replace(name, " ", "", -1)
	r = strings.ToLower(r)
	return
}

func (s *Summoner) GetCurrentGame() (game CurrentGameInfo, e error) {
	response, e := s.API.RequestEndpoint("/lol/spectator/v3/active-games/by-summoner/" + strconv.Itoa(s.ID))
	if e != nil {
		panic(e)
	}
	if response.StatusCode == 404 {
		e = errors.New("No current game found")
		return
	}
	defer response.Body.Close()
	game = CurrentGameInfo{}
	err := json.NewDecoder(response.Body).Decode(&game)
	if err != nil {
		panic(err)
	}
	return
}
func (api *GoLOLAPI) GetFeaturedGames() (games FeaturedGames, e error) {
	response, e := api.RequestEndpoint("/lol/spectator/v3/featured-games")
	if e != nil {
		panic(e)
	}
	games = FeaturedGames{}
	err := json.Unmarshal(response, &games)
	if err != nil {
		panic(err)
	}

	return
}
func createSummonerFromResponse(response *http.Response, api *GoLOLAPI) (s Summoner) {
	defer response.Body.Close()
	s = Summoner{}
	err := json.NewDecoder(response.Body).Decode(&s)
	if err != nil {
		panic(err)
	}
	s.Region = &api.Region
	s.API = api
	return

}
func (api *GoLOLAPI) GetSummonerByName(name string) (s Summoner) {
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-name/" + name)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
func (api *GoLOLAPI) GetSummonerByID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/" + IDString)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
func (api *GoLOLAPI) GetSummonerByAccountID(ID int) (s Summoner) {
	IDString := strconv.Itoa(ID)
	result, e := api.RequestEndpoint("/lol/summoner/v3/summoners/by-account/" + IDString)
	if e != nil {
		panic(e)
	}
	s = createSummonerFromResponse(result, api)
	return
}
func (api *GoLOLAPI) GetChampîons(onlyFreeToPlay bool) (champions ChampionList) {
	var endpoint string
	if onlyFreeToPlay {
		endpoint = "/lol/platform/v3/champions?freeToPlay=true"
	} else {
		endpoint = "/lol/platform/v3/champions?freeToPlay=true"

	}
	response, e := api.RequestEndpoint(endpoint)

	if e != nil {
		panic(e)
	}
	defer response.Body.Close()
	champions = ChampionList{}
	err := json.NewDecoder(response.Body).Decode(&champions)
	if err != nil {
		panic(err)
	}
	return
}
func (api *GoLOLAPI) GetChampîon(ID int) (champion Champion) {
	IDString := strconv.Itoa(ID)
	response, e := api.RequestEndpoint("/lol/platform/v3/champions/" + IDString)
	if e != nil {
		panic(e)
	}
	defer response.Body.Close()
	champion = Champion{}
	err := json.NewDecoder(response.Body).Decode(&champion)
	if err != nil {
		panic(err)
	}
	return
}

type GoLOLAPI struct {
	Region  Region
	APIKey  string
	limiter *rate.Limiter
	cache   *cache.Cache
}

var httpClient = &http.Client{}

func (api *GoLOLAPI) RequestEndpoint_(path string) (r *http.Response, e error) {
	cacheHit, found := api.cache.Get(path)
	if found {
		fmt.Println("Cache hit")
		hit, _ := cacheHit.(http.Response)
		return &hit, nil
	}
	reservation := api.limiter.ReserveN(time.Now(), 1)
	time.Sleep(reservation.Delay())
	r, e = http.Get(api.Region.Host + path + "?api_key=" + api.APIKey)
	if e != nil {
		panic(e)
	}
	if r.StatusCode == 429 {
		fmt.Println("RATE LIMIT EXCEEDED 429")
	}
	api.cache.Set(path, *r, cache.DefaultExpiration)
	return
}
func (api *GoLOLAPI) RequestEndpoint(path string) (r []byte, e error) {
	cacheHit, found := api.cache.Get(path)
	if found {
		fmt.Println("Cache hit")
		hit, _ := cacheHit.([]byte)
		return hit, nil
	}
	reservation := api.limiter.ReserveN(time.Now(), 1)
	time.Sleep(reservation.Delay())
	response, e := http.Get(api.Region.Host + path + "?api_key=" + api.APIKey)
	if e != nil {
		panic(e)
	}
	if response.StatusCode == 429 {
		fmt.Println("RATE LIMIT EXCEEDED 429")
	}
	r, _ = ioutil.ReadAll(response.Body)
	api.cache.Set(path, r, cache.DefaultExpiration)

	return
}

func NewAPI(region Region, APIKey string, rpers float64) (api GoLOLAPI) {
	api = GoLOLAPI{Region: region, APIKey: APIKey, limiter: rate.NewLimiter(rate.Limit(rpers), 2), cache: cache.New(1*time.Minute, 1*time.Minute)}
	return
}
func main() {
	api := NewAPI(EUW, "RGAPI-1180b131-9820-4e9d-a91d-ec77c7629839", 0.8)
	//summ := api.GetSummonerByName("FNC Rekkles")
	//summ := GetSummonerByName("je suis kaas")
	//summ := GetSummonerByID(71248364)
	//fmt.Println(summ)
	//game, e := summ.GetCurrentGame()
	//if e != nil {
	//	fmt.Println(e)
	//} else {
	//	fmt.Println(game)
	//}
	fmt.Println(api.GetFeaturedGames())
	fmt.Println(api.GetFeaturedGames())
	//fmt.Println("")

}
