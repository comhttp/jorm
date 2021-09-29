package coin

import (
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func NewCoin(slug string) (c *Coin) {
	c = new(Coin)

	c.Slug = slug
	return c
}


func SetCoin(src, slug string, get func(c *Coin)) {
	s:= strapi.New("http://127.0.0.1:1337")
	var cc []*Coin
	err := s.Get("coins",slug, &cc)
	utl.ErrorLog(err)
	if len(cc) != 0 {
		c := cc[0]
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		if !c.Checked[src] {
			log.Print("Check Coin: ", c.Slug)
			get(c)
			c.Checked[src] = true
		} else {
			get(c)
			log.Print("Already checked Coin: ", c.Slug)
		}
		s.Put("coins", c)
	} else {
		c := NewCoin(slug)
		log.Print("Insert Coin: ", slug)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		get(c)
		c.Checked[src] = true
		s.Post("coins", c)
	}
	return
}

//func (cq *CoinsQueries) SetCoin(src, slug string, get func(c *Coin)) {
//	c, err := cq.getCoin(slug)
//	if err != nil {
//		c = NewCoin(slug)
//		log.Print("Insert Coin: ", slug)
//		if c.Checked == nil {
//			c.Checked = make(map[string]bool)
//		}
//		get(c)
//		c.Checked[src] = true
//		//cq.WriteCoin(slug, c)
//		//utl.ErrorLog(err)
//	} else {
//		if c.Checked == nil {
//			c.Checked = make(map[string]bool)
//		}
//		if !c.Checked[src] {
//			log.Print("Check Coin: ", c.Slug)
//			get(c)
//			c.Checked[src] = true
//		} else {
//			get(c)
//			log.Print("Already checked Coin: ", c.Slug)
//		}
//		//cq.WriteCoin(slug, c)
//
//	}
//	return
//}

func (cq *CoinsQueries) WriteCoin(slug string, c interface{}) error {
	return cq.j.Write("coin", slug, c)
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

func (c *Coin) SetLogo(logo interface{}) {
	if logo.(string) != "" && logo.(string) != "missing_large.png" {
		//imgs := utl.GetIMG(logo.(string), cfg.Path+cfg.C.Out+"/imgs/", c.Slug)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "all", imgs)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "16", imgs.Img16)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "32", imgs.Img32)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "64", imgs.Img64)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "128", imgs.Img128)
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "256", imgs.Img256)
		//Create a empty file
		c.LogoBase64 = utl.GetIMG(logo.(string),c.Slug)
	}
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
