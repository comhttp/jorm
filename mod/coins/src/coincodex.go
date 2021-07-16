package csrc

import (
	"encoding/json"
	"fmt"
	coins2 "github.com/comhttp/jorm/mod/coins"
	jdb2 "github.com/comhttp/jorm/pkg/jdb"
	"io/ioutil"
	"net/http"

	"github.com/comhttp/jorm/pkg/utl"
)

func getCoinCodex(j *jdb2.JDB) {
	fmt.Println("GetCoinCodexStart")
	var coinsRaw []interface{}
	respcs, err := http.Get("https://coincodex.com/apps/coincodex/cache/all_coins.json")
	if err != nil {
	}
	defer respcs.Body.Close()
	mapBody, err := ioutil.ReadAll(respcs.Body)
	json.Unmarshal(mapBody, &coinsRaw)
	for _, cSrc := range coinsRaw {
		if cSrc != nil {
			coinSrc := cSrc.(map[string]interface{})
			if coinSrc["name"] != nil {
				slug := utl.MakeSlug(coinSrc["name"].(string))
				var coin coins2.Coin
				if j.Read("coins", slug, &coin) != nil {
					if coin.Checked == nil {
						coin.Checked = make(map[string]bool)
						if !coin.Checked["cx"] {
							//coin.Name = coinSrc["name"].(string)
							coin.Ticker = coinSrc["symbol"].(string)
							//coin.Slug = slug
							coinDetails := make(map[string]interface{})
							respcCoin, err := http.Get("https://coincodex.com/api/coincodex/get_coin/" + coin.Ticker)
							utl.ErrorLog(err)
							defer respcCoin.Body.Close()
							mapBodyCoin, err := ioutil.ReadAll(respcCoin.Body)
							utl.ErrorLog(err)
							json.Unmarshal(mapBodyCoin, &coinDetails)

							if coinDetails["description"] != nil {
								coin.Description = insertString(coinDetails["description"].(string), coin.Description)
							}
							//coin.WebSite = insertStringSlice(coinDetails["Website"], coin.WebSite)
							if coinDetails["totalsupply"] != nil {
								coin.TotalCoinSupply = insertFloat(coinDetails["totalsupply"].(float64), coin.TotalCoinSupply)
							}
							if coinDetails["whitepaper"] != nil {
								coin.WhitePaper = insertString(coinDetails["whitepaper"].(string), coin.WhitePaper)
							}

							if coinDetails["ico_price"] != nil {
								coin.Ico = true
								// jdb.WriteCoinData(slug, "ico", coinDetails.ICO)
								fmt.Println("Insert ICO Coin: ", coinDetails["ico_price"])
							}
							coin.Checked["cx"] = true
							//coin.SetLogo("https://coincodex.com/en/resources/images/admin/coins/" + slug + ".png")
							//coin.Logo = true
							j.Write("coins", slug, coin)
						}
					}
				}
			}
		}
	}
	fmt.Println("GetCoinCodexDone")
}
