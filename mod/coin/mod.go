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

type Coin struct {
	General
	GeneralCoin
	LinksCoin
	BlockchainCoin
	Rank     int               `json:"rank" form:"rank"`
	Platform string            `json:"platform" form:"platform"`
	BitNode  bool              `json:"bitnode" form:"bitnode"`
	Markets  map[string]string `json:"markets"`
	src      map[string]string `json:"markets"`
	isLogo   bool              `json:"islogo" form:"islogo"`
	Logo     utl.Images        `json:"logo" form:"logo"`
	Nodes    nodes.BitNodes    `json:"bitnode"`
}

// General stores identifying information about item in the database
type General struct {
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
	SubDomain   []string        `json:"subdomain" form:"subdomain"`
}

// GeneralCoin stores identifying information about a coin in a database
type GeneralCoin struct {
	Ticker string    `json:"ticker" form:"ticker"`
	Token  string    `json:"token" form:"token"`
	Algo   string    `json:"algo" form:"algo"`
	Proof  string    `json:"proof" form:"proof"`
	Start  time.Time `json:"start"`

	Ico     bool   `json:"ico" form:"ico"`
	BuiltOn string `json:"builton"`
}

// GeneralCoin stores identifying information about a coin in a database
type BlockchainCoin struct {
	NetHashesPerSecond   float64 `json:"nhps"`
	MaxSupply            float64 `json:"supply"`
	TotalCoinsMined      float64 `json:"mined"`
	BlockHeight          int     `json:"block_height"`
	BlockTime            float64 `json:"block_time"`
	DifficultyAdjustment string  `json:"diff"`
	BlockReward          int     `json:"block_rew"`
	BlockRewardReduction string  `json:"block_rewred"`
}

// GeneralCoin stores identifying information about a coin in a database
type LinksCoin struct {
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
	*General
	URL string `json:"url"`
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
