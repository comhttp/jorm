package minerstat

import "github.com/comhttp/jorm/pkg/utl"

type Coin struct {
	ID              string  `json:"id"`
	Coin            string  `json:"coin"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Algorithm       string  `json:"algorithm"`
	NetworkHashrate int     `json:"network_hashrate"`
	Difficulty      int     `json:"difficulty"`
	Reward          float64 `json:"reward"`
	RewardUnit      string  `json:"reward_unit"`
	RewardBlock     int     `json:"reward_block"`
	Price           float64 `json:"price"`
	Volume          float64 `json:"volume"`
	Updated         int     `json:"updated"`
}

func GetAllCoins() []Coin {
	allCoins := &[]Coin{}
	utl.GetSource("https://api.minerstat.com/v2/coins", allCoins)
	return *allCoins
}
