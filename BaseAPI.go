package gololapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
)

//Defines the different regions
var (
	EUW  = Region{Name: "EUW", PlatformID: "EUW1", Host: "https://euw1.api.riotgames.com"}
	BR   = Region{Name: "BR", PlatformID: "BR1", Host: "https://br1.api.riotgames.com"}
	EUNE = Region{Name: "EUNE", PlatformID: "EUN1", Host: "https://eun1.api.riotgames.com"}
	JP   = Region{Name: "JP", PlatformID: "JP1", Host: "https://jp1.api.riotgames.com"}
	KR   = Region{Name: "KR", PlatformID: "KR", Host: "https://kr.api.riotgames.com"}
	LAN  = Region{Name: "LAN", PlatformID: "LA1", Host: "https://la1.api.riotgames.com"}
	LAS  = Region{Name: "LAS", PlatformID: "LA2", Host: "https://la2.api.riotgames.com"}
	NA   = Region{Name: "NA", PlatformID: "NA1", Host: "https://na1.api.riotgames.com"}
	OCE  = Region{Name: "OCE", PlatformID: "OC1", Host: "https://oc1.api.riotgames.com"}
	TR   = Region{Name: "TR", PlatformID: "TR1", Host: "https://tr1.api.riotgames.com"}
	RU   = Region{Name: "RU", PlatformID: "RU", Host: "https://ru.api.riotgames.com"}
	PBE  = Region{Name: "PBE", PlatformID: "PBE1", Host: "https://pbe1.api.riotgames.com"}
)

//Region Struct containing information about a region
type Region struct {
	Name       string
	PlatformID string
	Host       string
}

//NewAPI Returns an instance of the api you can use.
func NewAPI(region Region, APIKey string, rpers float64) (api GoLOLAPI) {
	api = GoLOLAPI{Region: region, APIKey: APIKey, limiter: rate.NewLimiter(rate.Limit(rpers), 9), cache: cache.New(1*time.Minute, 1*time.Minute)}
	return
}

//GoLOLAPI This holds the api.It has the receiver for many of the API Endpoints.
type GoLOLAPI struct {
	Region  Region
	APIKey  string
	limiter *rate.Limiter
	cache   *cache.Cache
}

var httpClient = &http.Client{}

func GetEndpointURI(endpointPath string, options map[string]string) (uri string, hasParameters bool) {
	hasParameters = false
	if len(options) == 0 {
		uri = endpointPath
		return
	}
	if len(options) == 1 {
		for k, v := range options {
			uri = endpointPath + "?" + k + "=" + v
			hasParameters = true
			return
		}
	}
	if len(options) > 1 {
		uri = endpointPath + "?"
		hasParameters = true
		first := true
		for k, v := range options {
			if first {
				uri = uri + k + "=" + v
				first = false
			} else {
				uri = uri + "&" + k + "=" + v
			}
		}
		return
	}
	return
}
func (api *GoLOLAPI) RequestEndpoint(path string, cacheDuration time.Duration) (r []byte, e error) {
	cacheHit, found := api.cache.Get(path)
	if found {
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
	r, e2 := ioutil.ReadAll(response.Body)
	if e2 != nil {
		panic(e2)
	}

	if cacheDuration != 0 {
		api.cache.Set(path, r, cacheDuration)
	}
	return
}
func (api *GoLOLAPI) RequestStaticData(path string, cacheDuration time.Duration, withparameters bool) (r []byte, e error) {
	cacheHit, found := api.cache.Get(path)
	if found {
		fmt.Println("Cache hit")
		hit, _ := cacheHit.([]byte)
		return hit, nil
	}
	var uri string
	if withparameters {
		uri = api.Region.Host + path + "&api_key=" + api.APIKey
	} else {
		uri = api.Region.Host + path + "?api_key=" + api.APIKey
	}
	response, e := http.Get(uri)
	fmt.Println("get")
	if e != nil {
		panic(e)
	}
	if response.StatusCode == 429 {
		fmt.Println("RATE LIMIT EXCEEDED 429")
	}
	r, e2 := ioutil.ReadAll(response.Body)
	if e2 != nil {
		panic(e2)
	}

	api.cache.Set(path, r, cacheDuration)
	return
}
func MinifySummonerName(name string) (r string) {
	r = strings.Replace(name, " ", "", -1)
	r = strings.ToLower(r)
	return
}
