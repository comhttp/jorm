package exchange

import (
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
	"strings"
)

func NewExchange(slug string) (e *Exchange) {
	e = new(Exchange)
	e.Slug = slug
	return e
}

func (eq *ExchangeQueries) SetExchange(src, slug string, get func(e *Exchange)) {
	e, err := eq.getExchange(slug)
	if err != nil {
		e = NewExchange(slug)
		log.Print("Insert Exchange: ", slug)
		//if e.Checked == nil {
		//	e.Checked = make(map[string]bool)
		//}
		get(e)
		//e.Checked[src] = true
		eq.WriteExchange(slug, e)
		//utl.ErrorLog(err)
	} else {
		//if e.Checked == nil {
		//	e.Checked = make(map[string]bool)
		//}
		//if !e.Checked[src] {
		log.Print("Check Exchange: ", e.Slug)
		get(e)
		//e.Checked[src] = true
		//}else{
		//	get(e)
		//	log.Print("Already checked Coin: ", e.Slug)
		//}
		eq.WriteExchange(slug, e)
	}
	return
}

func (cq *ExchangeQueries) WriteExchange(slug string, c interface{}) error {
	return cq.j.Write("exchange", slug, c)
}

func (cq *ExchangeQueries) WriteMarkets(slug string, c interface{}) error {
	return cq.j.Write("market", slug, c)
}

func GetExchange(slug string) {
	//e := Exchange{}
	//_, err := os.Stat(filepath.FromSlash(cfg.Path + "/exchanges/" + slug))
	//if err != nil {
	//	e.Slug = slug
	//	log.Print("Insert Exchange: ", slug)
	//	//if c.Checked == nil {
	//	//	c.Checked = make(map[string]bool)
	//	//}
	//	//get(&e)
	//	//e.Checked[src] = true
	//	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, e)
	//} else {
	//	//err = jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, &e)
	//	utl.ErrorLog(err)
	//	log.Print("Ima Coin: ", e.Name)
	//	//if c.Checked == nil {
	//	//	c.Checked = make(map[string]bool)
	//	//}
	//	//if !c.Checked[src] {
	//	//	log.Print("Check Coin: ", c.Name)
	//	//	get(&c)
	//	//	c.Checked[src] = true
	//	//}
	//	//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, e)
	//}
	return
}

func (e *Exchange) SetName(name interface{}) {
	e.Name = utl.InsertString(e.Name, name)
	return
}

func (e *Exchange) SetDescription(description interface{}) {
	e.Description = utl.InsertString(e.Description, description)
	return
}

func (e *Exchange) SetURL(url interface{}) {
	e.Url = utl.InsertString(e.Url, url)
	return
}

func (e *Exchange) SetLogo(logo interface{}) {
	//if c.Logo.Img256 == "" {
	//	var cImgs utl.Images
	//	if logo.(string) != "" && logo.(string) != "missing_large.png" {
	//		cImgs = utl.GetIMG(logo.(string), cfg.Path+"/static/coins/", c.Slug)
	//	}
	//	c.Logo = cImgs
	//}
	return
}

func (e *ExchangeSrc) SetCurrencyMarket(market, symbol string, ask, bid, high, last, low, vol interface{}) {
	e.Markets[market].Currencies[symbol] = Currency{
		Symbol: symbol,
		Ask:    utl.InsertFloat(ask),
		Bid:    utl.InsertFloat(bid),
		High:   utl.InsertFloat(high),
		Last:   utl.InsertFloat(last),
		Low:    utl.InsertFloat(low),
		Volume: utl.InsertFloat(vol),
	}
}

func (eq *ExchangeQueries) SetMarkets(slug string, marketsSrc map[string]MarketSrc) {
	markets := map[string]Market{}
	for _, exs := range marketsSrc {
		mSrc := Market{}
		for _, cur := range exs.Currencies {
			mSrc.Symbol = exs.Symbol
			mSrc.Currencies = append(mSrc.Currencies, cur)
		}
		markets[strings.ToLower(exs.Symbol)] = mSrc
	}
	eq.WriteMarkets(slug, markets)
}

//func (eq *ExchangeQueries) SetTicker(ex *ExchangeSrc) {
//	log.Print("Get " + ex.Name + " Exchange Start")
//	var e Exchange
//	e.Name = ex.Name
//	e.Slug = ex.Slug
//	marketsSrc := utl.GetSource(ex.APIUrl).(map[string]interface{})
//	if len(marketsSrc) > 0 {
//		es.Markets = make(map[string]MarketSrc)
//
//		for key, marketSrcRaw := range marketsSrc {
//			marketSrc := marketSrcRaw.(map[string]interface{})
//			m := strings.Split(key, "_")
//			if nq := m[0]; nq != es.Markets[nq].Symbol {
//				es.Markets[nq] = MarketSrc{
//					Symbol:     nq,
//					Currencies: make(map[string]Currency),
//				}
//			}
//			es.SetCurrencyMarket(
//				m[0],
//				m[1],
//				marketSrc[es.Ticker.Ask],
//				marketSrc[es.Ticker.Bid],
//				marketSrc[es.Ticker.High24],
//				marketSrc[es.Ticker.Last],
//				marketSrc[es.Ticker.Low24],
//				marketSrc[es.Ticker.Vol])
//		}
//		e.WriteExchange(j, *es)
//		log.Print("Get " + e.Name + " Exchange Done")
//	} else {
//		log.Print("Get " + e.Name + " Exchange Fail")
//	}
//}
