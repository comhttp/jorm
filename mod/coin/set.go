package coin

import (
	"strings"
	"time"

	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
)

func NewCoin(slug string) (c *Coin) {
	c = new(Coin)
	c.Slug = slug
	return c
}

func (c *Coin) SetSrcID(src, id string) {
	if c.SrcID == nil {
		c.SrcID = make(map[string]string)
	}
	c.SrcID[src] = id
}
func (c *Coin) SetName(name interface{}) {
	c.Name = utl.InsertString(c.Name, name)
	return
}

func (c *Coin) SetSymbol(ticker interface{}) {
	c.Symbol = strings.ToUpper(utl.InsertString(c.Symbol, ticker))
	return
}

func (c *Coin) SetAlgo(algo interface{}) {
	c.Algo = utl.InsertString(c.Algo, algo)
	return
}

func (c *Coin) SetDescription(description interface{}) {
	c.Description = utl.InsertString(c.Description, description)
	return
}

func (c *Coin) SetProof(proof interface{}) {
	c.Proof = utl.InsertString(c.Proof, proof)
	return
}

func (c *Coin) SetBuiltOn(proof interface{}) {
	c.Proof = utl.InsertString(c.Proof, proof)
	return
}

func (c *Coin) SetGenesisDate(start interface{}) {
	s, err := time.Parse("2017-07-01", start.(string))
	log.Log().Err(err)
	c.GenesisDate = s
	return
}

func (c *Coin) SetBitcoinTalk(bitcointalk interface{}) {
	//c.BitcoinTalk = utl.InsertString(c.BitcoinTalk, bitcointalk)
	return
}

func (c *Coin) SetTwitter(twitter interface{}) {
	//c.Twitter = utl.InsertString(c.Twitter, twitter)
	return
}

func (c *Coin) SetReddit(reddit interface{}) {
	//c.Reddit = utl.InsertString(c.Reddit, reddit)
	return
}

func (c *Coin) SetTelegram(telegram interface{}) {
	//c.Telegram = utl.InsertString(c.Telegram, telegram)
	return
}

func (c *Coin) SetFacebook(facebook interface{}) {
	//c.Facebook = utl.InsertString(c.Facebook, facebook)
	return
}

func (c *Coin) SetWebSite(website interface{}) {
	//c.WebSite = utl.InsertStringSlice(c.WebSite, website)
	return
}

func (c *Coin) SetExplorer(explorer interface{}) {
	//c.Explorer = utl.InsertStringSlice(c.Explorer, explorer)
	return
}

func (c *Coin) SetChat(chat interface{}) {
	//c.Chat = utl.InsertStringSlice(c.Chat, chat)
	return
}

func (c *Coin) SetNetworkHashrate(supply interface{}) {
	c.NetworkHashrate = utl.InsertFloat(supply)
	return
}
func (c *Coin) SetMaxSupply(supply interface{}) {
	c.MaxSupply = supply.(float64)
	return
}

func (c *Coin) SetTotalCoinsMined(supply interface{}) {
	c.TotalCoinsMined = supply.(float64)
	return
}

func (c *Coin) SetBlockHeight(supply interface{}) {
	c.BlockHeight = supply.(int)
	return
}

func (c *Coin) SetBlockTime(blockTime interface{}) {
	c.BlockTime = blockTime.(int)
	return
}

func (c *Coin) SetDifficulty(diff interface{}) {
	c.Difficulty = utl.InsertFloat(diff)
	return
}
func (c *Coin) SetDifficultyAdjustment(diffadj interface{}) {
	c.DifficultyAdjustment = utl.InsertString(c.DifficultyAdjustment, diffadj)
	return
}

func (c *Coin) SetBlockReward(rew interface{}) {
	c.BlockReward = rew.(float64)
	return
}
func (c *Coin) SetBlockRewardReduction(rew interface{}) {
	c.BlockRewardReduction = rew.(string)
	return
}
