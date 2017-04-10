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

//EUW Values for the EUW Region
var EUW = Region{Name: "EUW", PlatformID: "EUW1", Host: "https://euw1.api.riotgames.com", DefaultLocale: "en_GB"}

//Region Struct containing information about a region
type Region struct {
	Name          string
	PlatformID    string
	Host          string
	DefaultLocale string
}

//NewAPI Returns an instance of the api you can use.
func NewAPI(region Region, APIKey string, rpers float64) (api GoLOLAPI) {
	api = GoLOLAPI{Region: region, APIKey: APIKey, limiter: rate.NewLimiter(rate.Limit(rpers), 9), cache: cache.New(1*time.Minute, 1*time.Minute)}
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
func getEndpointURI(endpointPath string, options map[string]string) (uri string, hasParameters bool) {
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
		for k, v := range options {
			uri = uri + "&" + k + "=" + v
		}
		return
	}
	return
}
func (api *GoLOLAPI) RequestEndpoint(path string, cacheDuration time.Duration) (r []byte, e error) {
	cacheHit, found := api.cache.Get(path)
	if found {
		fmt.Println("Cache hit")
		hit, _ := cacheHit.([]byte)
		return hit, nil
	}
	reservation := api.limiter.ReserveN(time.Now(), 1)
	time.Sleep(reservation.Delay())
	response, e := http.Get(api.Region.Host + path + "?api_key=" + api.APIKey)
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
