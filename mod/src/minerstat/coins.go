package minerstat

import (
	"fmt"
	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/utl"
)

type Coin struct {
	ID              string  `json:"id"`
	Coin            string  `json:"coin"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Algorithm       string  `json:"algorithm"`
	NetworkHashrate int     `json:"network_hashrate"`
	Difficulty      float64 `json:"difficulty"`
	Reward          float64 `json:"reward"`
	RewardUnit      string  `json:"reward_unit"`
	RewardBlock     int     `json:"reward_block"`
	Price           float64 `json:"price"`
	Volume          float64 `json:"volume"`
	Updated         int     `json:"updated"`
}

func GetAllCoins(coinQueries *coin.CoinQueries) {
	allCoins := []Coin{}
	utl.GetSource("https://api.minerstat.com/v2/coins", &allCoins)
	fmt.Println("::::::::::::::::::::::::::::::::START minerstat COINS:::::::::::::::::::::::::::::: ")
	for _, msCoin := range allCoins {
		if msCoin.Name != "" {
			slug := utl.MakeSlug(msCoin.Name)
			switch msCoin.Type {
			case "coin":
				coinQueries.SetCoin("minerstat", slug, getMinerStatCoin(msCoin))
			case "pool":
			default:
			}
		}
	}
	fmt.Println("::::::::::::::::::::::::::::::::END minerstat COINS:::::::::::::::::::::::::::::: ")
	return
}

func getMinerStatCoin(msCoin Coin) func(c *coin.Coin) {
	return func(c *coin.Coin) {
		//if msCoin.ImageURL != "" && msCoin.ImageURL != "<nil>" {
		//	c.SetLogo("https://cryptocompare.com" + msCoin.ImageURL)
		//}
		c.SetSrcID("minerstat", msCoin.ID)
		c.SetName(msCoin.Name)
		c.SetSymbol(msCoin.Coin)
		c.SetAlgo(msCoin.Algorithm)
		c.SetNetworkHashrate(msCoin.NetworkHashrate)
		c.SetDifficulty(msCoin.Difficulty)

		//c.SetProof(msCoin.ProofType)
		//c.SetStart(msCoin.AssetLaunchDate)

		//c.SetMaxSupply(msCoin.MaxSupply)
		//c.SetTotalCoinsMined(msCoin.TotalCoinsMined)
		//c.SetBlockHeight(msCoin.BlockNumber)
		//c.SetBlockTime(int(msCoin.BlockTime))
		//c.SetBlockReward(msCoin.BlockReward)
	}
}
