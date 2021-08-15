package coin

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"time"
)

type CoinQueries struct {
	j   *jdb.JDB
	col string
}

//type CoinQueries struct {
//	j *jdb.JDB
//	col string
//}

type Coin struct {
	Name                 string          `json:"name" form:"name"`
	Slug                 string          `json:"slug" form:"slug"`
	Rank                 int             `json:"rank" form:"rank"`
	Platform             string          `json:"platform" form:"platform"`
	Description          string          `json:"description" form:"description"`
	Published            bool            `json:"published" form:"published"`
	Selected             bool            `json:"selected" form:"selected"`
	Favorite             bool            `json:"fav" form:"favorite"`
	Checked              map[string]bool `json:"checked"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Ticker               string            `json:"ticker" form:"ticker"`
	Token                string            `json:"token" form:"token"`
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
	WebSite              []string          `json:"web"`
	Explorer             []string          `json:"explorer"`
	Chat                 []string          `json:"chat"`
	Twitter              string            `json:"tw"`
	Facebook             string            `json:"facebook"`
	Telegram             string            `json:"telegram"`
	Reddit               string            `json:"reddit"`
	Github               []string          `json:"github"`
	BitcoinTalk          string            `json:"bitcointalk"`
	WhitePaper           string            `json:"whitepaper"`
	isLogo               bool              `json:"islogo" form:"islogo"`
	Logo                 utl.Images        `json:"logo" form:"logo"`
	Nodes                nodes.BitNodes    `json:"bitnode"`
	SubDomain            bool              `json:"subdomain" form:"subdomain"`
	Order                int               `json:"order" form:"order"`
}

// Coin stores identifying information about coins in the database
type CoinShort struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
}

type CoinsShort struct {
	N int         `json:"n"`
	C []CoinShort `json:"c"`
}

type Coins struct {
	N int      `json:"n"`
	C []string `json:"c"`
}

type AlgoCoins struct {
	N int         `json:"n"`
	A []string    `json:"a"`
	C []CoinShort `json:"c"`
}
type CoinsWords struct {
	N int    `json:"n"`
	C string `json:"c"`
}
