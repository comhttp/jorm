package coin

import (
	"strings"
	"time"

	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/rs/zerolog/log"
)

func NewCoin(slug string) (c *Coin) {
	c = new(Coin)
	c.Slug = slug
	return c
}

func SetCoin(s strapi.StrapiRestClient, src, slug string, get func(c *Coin, l *img.Logo)) {
	var cc []*Coin
	err := s.Get("coins", slug, &cc)
	utl.ErrorLog(err)
	var ll []*img.Logo
	err = s.Get("logos", slug, &ll)
	utl.ErrorLog(err)
	l := &img.Logo{
		Slug: slug,
	}
	if len(cc) != 0 {
		c := cc[0]
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		if !c.Checked[src] {
			log.Print("Check Coin: ", c.Slug)
			get(c, l)
			if ll[0].Data == "" {
				s.Post("logos", l)
			}
			c.Checked[src] = true
		} else {
			get(c, l)
			log.Print("Already checked Coin: ", c.Slug)
		}
		s.Put("coins", c)
	} else {
		c := NewCoin(slug)
		log.Print("Insert Coin: ", slug)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		get(c, l)
		c.Checked[src] = true
		s.Post("coins", c)
		s.Post("logos", l)
	}
	return
}

func (cq *CoinsQueries) SetCoin(src, slug string, get func(c Coin)) {
	c, err := cq.getCoin(slug)
	if err != nil {
		c = NewCoin(slug)
		log.Print("Insert Coin: ", slug)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		get(*c)
		c.Checked[src] = true
		//cq.WriteCoin(slug, c)
		//utl.ErrorLog(err)
	} else {
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		if !c.Checked[src] {
			log.Print("Check Coin: ", c.Slug)
			get(*c)
			c.Checked[src] = true
		} else {
			get(*c)
			log.Print("Already checked Coin: ", c.Slug)
		}
		//cq.WriteCoin(slug, c)

	}
	return
}

func (cq *CoinsQueries) WriteCoin(slug string, c interface{}) error {
	return cq.j.Write("coin", slug, c)
}

func (cq *CoinsQueries) WriteLogo(slug string, c interface{}) error {
	return cq.j.Write("logo", slug, c)
}

func (cq *CoinsQueries) WriteInfo(slug string, c interface{}) error {
	return cq.j.Write("info", slug, c)
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
