package exchange

import (
	"fmt"
	"github.com/comhttp/jorm/app/cfg"
	"github.com/comhttp/jorm/pkg/utl"
	"os"
	"path/filepath"
)

func GetExchange(src, slug string, get func(e *Exchange)) {
	e := Exchange{}
	_, err := os.Stat(filepath.FromSlash(cfg.Path + "/exchanges/" + slug))
	if err != nil {
		e.Slug = slug
		fmt.Println("Insert Exchange: ", slug)
		//if c.Checked == nil {
		//	c.Checked = make(map[string]bool)
		//}
		//get(&e)
		//e.Checked[src] = true
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, e)
	} else {
		//err = jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, &e)
		utl.ErrorLog(err)
		fmt.Println("Ima Coin: ", e.Name)
		//if c.Checked == nil {
		//	c.Checked = make(map[string]bool)
		//}
		//if !c.Checked[src] {
		//	fmt.Println("Check Coin: ", c.Name)
		//	get(&c)
		//	c.Checked[src] = true
		//}
		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/exchanges"), slug, e)
	}
	return
}

func (e *Exchange) SetName(name interface{}) {
	//c.Name = insertString(c.Name, name)
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
