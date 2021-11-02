package coin

import (
	"time"

	"github.com/comhttp/jorm/pkg/jdb"
)

type CoinsQueries struct {
	j   *jdb.JDB
	col string
}

type Coin struct {
	Id                   int               `json:"id" form:"id"`
	Name                 string            `json:"name" form:"name"`
	Slug                 string            `json:"slug" form:"slug"`
	Description          string            `json:"description" form:"description"`
	Published            bool              `json:"published" form:"published"`
	Selected             bool              `json:"selected" form:"selected"`
	Favorite             bool              `json:"fav" form:"favorite"`
	Checked              map[string]bool   `json:"checked"`
	CreatedAt            time.Time         `json:"created"`
	UpdatedAt            time.Time         `json:"updated"`
	Order                int               `json:"order" form:"order"`
	SubDomain            []string          `json:"subdomain" form:"subdomain"`
	Symbol               string            `json:"symbol" form:"symbol"`
	Token                string            `json:"token" form:"token"`
	Algo                 string            `json:"algo" form:"algo"`
	Proof                string            `json:"proof" form:"proof"`
	Ico                  bool              `json:"ico" form:"ico"`
	BuiltOn              string            `json:"builton"`
	GenesisDate          time.Time         `json:"genesis"`
	NetworkHashrate      float64           `json:"hashrate"`
	MaxSupply            float64           `json:"supply"`
	TotalCoinsMined      float64           `json:"mined"`
	BlockHeight          int               `json:"blockheight"`
	BlockTime            int               `json:"blocktime"`
	Difficulty           float64           `json:"difficulty"`
	DifficultyAdjustment string            `json:"difficultyadjustment"`
	BlockReward          float64           `json:"blockreward"`
	BlockRewardReduction string            `json:"blockrewardreduction"`
	Rank                 int               `json:"rank" form:"rank"`
	Platform             string            `json:"platform" form:"platform"`
	BitNode              bool              `json:"bitnode" form:"bitnode"`
	SrcID                map[string]string `json:"srcid"`
	// Logo                 []byte            `json:"logo" form:"logo"`
	// LogoBase64           utl.Images        `json:"logobase64" form:"logobase64"`
	// Links Links `json:"links" form:"links"`
	//Nodes    nodes.BitNodes    `json:"bitnode"`
	//Markets  map[string]string `json:"markets"`
}

// General stores identifying information about item in the database
type Item struct {
	Name        string          `json:"name" form:"name"`
	Slug        string          `json:"slug" form:"slug"`
	Description string          `json:"description" form:"description"`
	Published   bool            `json:"published" form:"published"`
	Selected    bool            `json:"selected" form:"selected"`
	Favorite    bool            `json:"fav" form:"favorite"`
	Checked     map[string]bool `json:"checked"`
	CreatedAt   time.Time       `json:"created"`
	UpdatedAt   time.Time       `json:"updated"`
	Order       int             `json:"order" form:"order"`
	//SubDomain   []string        `json:"subdomain" form:"subdomain"`
}

// GeneralCoin stores identifying information about a coin in a database
type General struct {
	Symbol string `json:"symbol" form:"symbol"`
	Token  string `json:"token" form:"token"`
	Algo   string `json:"algo" form:"algo"`
	Proof  string `json:"proof" form:"proof"`

	Ico     bool   `json:"ico" form:"ico"`
	BuiltOn string `json:"builton"`
}

// GeneralCoin stores identifying information about a coin in a database
type Blockchain struct {
	GenesisDate          time.Time `json:"genesis"`
	NetworkHashrate      float64   `json:"hashrate"`
	MaxSupply            float64   `json:"supply"`
	TotalCoinsMined      float64   `json:"mined"`
	BlockHeight          int       `json:"block_height"`
	BlockTime            int       `json:"block_time"`
	Difficulty           float64   `json:"difficulty"`
	DifficultyAdjustment string    `json:"diffadj"`
	BlockReward          float64   `json:"block_rew"`
	BlockRewardReduction string    `json:"block_rewred"`
}

// GeneralCoin stores identifying information about a coin in a database
type Links struct {
	WebSite     []Link `json:"web"`
	Explorer    []Link `json:"explorer"`
	Chat        []Link `json:"chat"`
	Twitter     Link   `json:"tw"`
	Facebook    Link   `json:"facebook"`
	Telegram    Link   `json:"telegram"`
	Reddit      Link   `json:"reddit"`
	Github      []Link `json:"github"`
	BitcoinTalk Link   `json:"bitcointalk"`
	WhitePaper  Link   `json:"whitepaper"`
}

// GeneralCoin stores identifying information about a coin in a database
type Link struct {
	*Item
	URL string `json:"url"`
}

// Coin stores identifying information about coins in the database
type CoinShort struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Symbol string `json:"y"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
}

type AlgoCoins struct {
	N int         `json:"n"`
	C []CoinShort `json:"c"`
	A []string    `json:"a"`
}
type CoinsShort struct {
	N int         `json:"n"`
	C []CoinShort `json:"c"`
}

type Coins struct {
	N int      `json:"n"`
	C []string `json:"c"`
}

type CoinUser struct {
	Id                   int       `json:"id" form:"id"`
	Name                 string    `json:"name" form:"name"`
	Slug                 string    `json:"slug" form:"slug"`
	Description          string    `json:"description" form:"description"`
	Selected             bool      `json:"selected" form:"selected"`
	Favorite             bool      `json:"fav" form:"favorite"`
	UpdatedAt            time.Time `json:"updated"`
	Order                int       `json:"order" form:"order"`
	SubDomain            []string  `json:"subdomain" form:"subdomain"`
	Symbol               string    `json:"symbol" form:"symbol"`
	Token                string    `json:"token" form:"token"`
	Algo                 string    `json:"algo" form:"algo"`
	Proof                string    `json:"proof" form:"proof"`
	Ico                  bool      `json:"ico" form:"ico"`
	BuiltOn              string    `json:"builton"`
	GenesisDate          time.Time `json:"genesis"`
	NetworkHashrate      float64   `json:"hashrate"`
	MaxSupply            float64   `json:"supply"`
	TotalCoinsMined      float64   `json:"mined"`
	BlockHeight          int       `json:"blockheight"`
	BlockTime            int       `json:"blocktime"`
	Difficulty           float64   `json:"difficulty"`
	DifficultyAdjustment string    `json:"difficultyadjustment"`
	BlockReward          float64   `json:"blockreward"`
	BlockRewardReduction string    `json:"blockrewardreduction"`
	Rank                 int       `json:"rank" form:"rank"`
	Platform             string    `json:"platform" form:"platform"`
	BitNode              bool      `json:"bitnode" form:"bitnode"`
}

type Logo struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Extension string `json:"ext"`
	Data      []byte `json:"data"`
}
