package storage

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/pkg/utl"
	"image"
	"strings"

	"github.com/comhttp/jorm/app/jdb"
)

// Coin stores identifying information about coins in the database
type CoinBase struct {
	Rank   int    `json:"r"`
	Name   string `json:"n"`
	Ticker string `json:"t"`
	Slug   string `json:"s"`
	Algo   string `json:"a"`
	Img    string `json:"i"`
}

type Storage struct {
	Name   string `json:"name" form:"name"`
	Ticker string `json:"ticker" form:"ticker"`
	Slug   string `json:"slug" form:"slug"`

	Rank        int    `json:"rank" form:"rank"`
	Token       bool   `json:"token" form:"token"`
	Platform    string `json:"platform" form:"platform"`
	Algo        string `json:"algo" form:"algo"`
	Proof       string `json:"proof" form:"proof"`
	Description string `json:"description" form:"description"`
	BuiltOn     string `json:"builton"`
	Start       string `json:"start"`

	WebSite     []string   `json:"web"`
	Explorer    []string   `json:"explorer"`
	Chat        []string   `json:"chat"`
	Twitter     string     `json:"tw"`
	Facebook    string     `json:"facebook"`
	Telegram    string     `json:"telegram"`
	Reddit      string     `json:"reddit"`
	Github      []string   `json:"github"`
	BitcoinTalk string     `json:"bitcointalk"`
	WhitePaper  string     `json:"whitepaper"`
	Logo        utl.Images `json:"logo" form:"logo"`

	Published bool `json:"published" form:"published"`
	BitNode   bool `json:"bitnode" form:"bitnode"`
	Selected  bool `json:"selected" form:"selected"`
	Favorite  bool `json:"fav" form:"favorite"`

	Checked map[string]bool `json:"checked"`
}

//
//// ReadAllCoins reads in all of the data about all coins in the database
//func ReadAllCoins() Coins {
//	csb := LoadCoinsBase(false, false)
//	cns := Coins{
//		N: csb.N,
//		C: getCoins(),
//	}
//
//	jdb.JDB.Write("jorm/info", "coinsbase", csb)
//
//	jdb.JDB.Write("jorm/info", "bitnodes", LoadCoinsBase(true, true))
//	jdb.JDB.Write("jorm/info", "coins", LoadCoinsBase(false, true))
//	return cns
//}
//
//func (coin *Coin) SelectCoin() *Coin {
//	//coin.LogoBig = LoadLogo(coin.Slug, "img128")
//	//coin.Data = LoadInfo(coin.Slug)
//	return coin
//}
//func LoadLogo(slug, size string) image.Image {
//	// Load logo image from database
//	logos := make(map[string]interface{})
//	fmt.Println("slug", slug)
//	err := jdb.JDB.Read("jorm/data/"+slug, "logo", logos)
//	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
//	logo, _, err := image.Decode(reader)
//	utl.ErrorLog(err)
//	return logo
//}
//
//func LoadInfo(slug string) Coin {
//	// Load coin data from database
//	info := Coin{}
//	err := jdb.JDB.Read("data/"+slug, "info", info)
//	utl.ErrorLog(err)
//	//jsonString, _ := json.Marshal(info)
//
//	// convert json to struct
//	//s := CoinData{}
//	//json.Unmarshal(jsonString, &s)
//	return info
//}
//
//func LoadCoinsBase(bitnode, filter bool) CoinsBase {
//	coins := getCoins()
//	csb := CoinsBase{}
//	csb.N = 0
//	for i, coin := range coins {
//		ccb := CoinBase{
//			Rank:   csb.N,
//			Name:   coin.Name,
//			Ticker: coin.Ticker,
//			Slug:   coin.Slug,
//			Algo:   coin.Algo,
//		}
//		if filter {
//			if coins[i].Platform != "token" &&
//				coins[i].Algo != "" &&
//				coins[i].Algo != "N/A" &&
//				coins[i].Twitter != "" &&
//				coins[i].Description != "" &&
//				len(coins[i].WebSite) > 0 &&
//				len(coins[i].Explorer) > 0 {
//				if bitnode {
//					if coins[i].BitNode {
//						ccb.Img = coin.Logo.Img64
//						csb.N++
//						csb.C = append(csb.C, ccb)
//					}
//				} else {
//					ccb.Img = coin.Logo.Img16
//					csb.N++
//					csb.C = append(csb.C, ccb)
//				}
//			}
//		} else {
//			csb.N++
//			csb.C = append(csb.C, ccb)
//		}
//
//	}
//	return csb
//}
//
//func getCoins() []Coin {
//	data, err := jdb.JDB.ReadAll("jorm/coins")
//	utl.ErrorLog(err)
//	coins := make([][]byte, len(data))
//	for i := range data {
//		coins[i] = []byte(data[i])
//	}
//	cs := make([]Coin, len(coins))
//	for i := range coins {
//		if err := json.Unmarshal(coins[i], &cs[i]); err != nil {
//			fmt.Println("Error", err)
//		}
//	}
//	return cs
//}
