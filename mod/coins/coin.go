package coins

import (
	"github.com/comhttp/jorm/mod/nodes"
	"github.com/comhttp/jorm/pkg/utl"
	"time"
)

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
}

type Coins struct {
	N int      `json:"n"`
	C []string `json:"c"`
}

type AlgoCoins struct {
	N int        `json:"n"`
	A []string   `json:"a"`
	C []AlgoCoin `json:"c"`
}
type CoinsWords struct {
	N int    `json:"n"`
	C string `json:"c"`
}
type BaseCoins struct {
	N int        `json:"n"`
	C []BaseCoin `json:"c"`
}

// Coin stores identifying information about coins in the database
type BaseCoin struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
}

// Coin stores identifying information about coins in the database
type AlgoCoin struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
}

//func LoadLogo(slug, size string) image.Image {
//	// Load logo image from database
//	logos := make(map[string]interface{})
//	log.Println("slug", slug)
//	err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/data/"+slug), "logo", logos)
//	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
//	logo, _, err := image.Decode(reader)
//	log.Println(err)
//	return logo
//}
//
//func LoadInfo(slug string) Coin {
//	// Load coin data from database
//	info := Coin{}
//	err := jdb.JDB.Read(filepath.FromSlash("data/"+slug), "info", info)
//	log.Println(err)
//	//jsonString, _ := json.Marshal(info)
//
//	// convert json to struct
//	//s := CoinData{}
//	//json.Unmarshal(jsonString, &s)
//	return info
//}
