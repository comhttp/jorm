package coingecko

import (
	"fmt"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/rs/zerolog/log"
	gecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
	"time"
)

func GetAllCoins(coinQueries *coin.CoinsQueries) {
	cg := gecko.NewClient(nil)
	cgCoins, err := cg.CoinsList()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("::::::::::::::::::::::::::::::::END Coin Gecko COINS:::::::::::::::::::::::::::::: ")
	for _, cgCoin := range *cgCoins {
		if cgCoin.Name != "" {
			cgCoinID, err := cg.CoinsID(cgCoin.ID, false, false, false, false, false, false)
			if err != nil {
				log.Print(err)
			}
			coinQueries.SetCoin("coingecko", cgCoin.ID, getCoinGeckoCoin(cgCoinID))
			time.Sleep(1200 * time.Millisecond)
		}
	}
	fmt.Println("Available coins:", len(*cgCoins))
	fmt.Println("::::::::::::::::::::::::::::::::END Coin Gecko COINS:::::::::::::::::::::::::::::: ")
}

func getCoinGeckoCoin(cgCoin *types.CoinsID) func(c *coin.Coin) {
	return func(c *coin.Coin) {
		//if cgCoin.Image.Large != "" {
		//	c.SetLogo(cgCoin.Image.Large)
		//}
		c.SetSrcID("coingecko", cgCoin.ID)

		c.SetName(cgCoin.Name)
		c.SetSymbol(cgCoin.Symbol)

		//c.SetDescription(cgCoin.Description["en"])
		//c.SetAlgo(cgCoin.Algorithm)
		//c.SetProof(cgCoin.ProofType)
		//c.SetStart(cgCoin.GenesisDate)
		//c.SetTotalCoinsMined(cgCoin.TotalCoinsMined)

		//c.SetBlockTime(int(cgCoin.BlockTimeInMin))
	}
}
