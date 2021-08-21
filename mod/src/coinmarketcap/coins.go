package coinmarketcap

import (
	"fmt"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/utl"
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

type coinmarketcap struct {
	apiKey string
	*cmc.Client
}

func NewCoinMarketCapAPI(apiKey string) *coinmarketcap {
	return &coinmarketcap{apiKey, cmc.NewClient(&cmc.Config{
		ProAPIKey: apiKey,
	}),
	}
}

func (cm *coinmarketcap) GetAllCoins(coinQueries *coin.CoinsQueries) {
	listings, err := cm.Cryptocurrency.Map(&cmc.MapOptions{
		//Limit: 999,
	})
	utl.ErrorLog(err)
	for i, listing := range listings {
		fmt.Println("COIN:", listing.Name)
		fmt.Println("COIN No:", i)
	}
}
