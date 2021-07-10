package coin

import (
	"github.com/comhttp/jorm/mod/rav"
	"time"
)

type Coin struct {
	*rav.Item
	Ticker               string            `json:"ticker" form:"ticker"`
	Rank                 int               `json:"rank" form:"rank"`
	Token                string            `json:"token" form:"token"`
	Platform             string            `json:"platform" form:"platform"`
	Algo                 string            `json:"algo" form:"algo"`
	Proof                string            `json:"proof" form:"proof"`
	Start                time.Time         `json:"start"`
	Ico                  bool              `json:"ico" form:"ico"`
	TotalCoinSupply      float64           `json:"total"`
	BuiltOn              string            `json:"builton"`
	BlockTime            float64           `json:"blocktime"`
	DifficultyAdjustment string            `json:"diff"`
	BlockRewardReduction string            `json:"rew"`
	BitNode              bool              `json:"bitnode" form:"bitnode"`
	Markets              map[string]string `json:"markets"`
	*rav.Media           `json:"media"`
}
