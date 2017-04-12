# GoLolApi

A wrapper around the League Of Legends API written in Go.

## Getting Started

### Installing
```
go get github.com/Djipyy/GoLolApi
```

### Exemple
```Go 
package main

import (
	"fmt"

	gololapi "github.com/Djipyy/GoLolApi"
)

func main() {
	api := gololapi.NewAPI(gololapi.EUW, "YOUR_API_KEY", 0.8) //0.8 is the rate limit for the developpement key, you should change it if you have a production key
	summ := api.GetSummonerByID(71248364)
	fmt.Printf("His summoner name is %s", summ.Name)

	fmt.Println(api.GetFeaturedGames()) //This will return the current featured games of the region
}
```

## Built With

* [go-cache](https://github.com/patrickmn/go-cache/) - Library used for caching
* [rate](https:/golang.org/x/time/rate/) - Library for the rate limiting

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
