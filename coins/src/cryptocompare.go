package csrc

import (
	"encoding/json"
	"fmt"
	"github.com/comhttp/jorm/coins"
	jdb2 "github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"io/ioutil"
	"net/http"
)

func getCryptoCompare(j *jdb2.JDB) {
	fmt.Println("Get Crypto Compare Start")
	respcs, err := http.Get("https://min-api.cryptocompare.com/data/all/coinlist")
	utl.ErrorLog(err)
	defer respcs.Body.Close()
	if respcs != nil {
		coinsRaw := make(map[string]interface{})
		mapBody, err := ioutil.ReadAll(respcs.Body)
		utl.ErrorLog(err)
		json.Unmarshal(mapBody, &coinsRaw)
		if coinsRaw["Data"] != nil {
			for _, coinSrc := range coinsRaw["Data"].(map[string]interface{}) {
				if coinSrc != nil {
					cs := coinSrc.(map[string]interface{})
					if cs["CoinName"] != nil {
						slug := utl.MakeSlug(cs["CoinName"].(string))
						coins.SetCoin(j, "cc", slug, getCryptoCompareCoin(slug, cs))
					}
				}
			}
		}
	}
	fmt.Println("Get Crypto Compare Done")
}

func getCryptoCompareCoin(slug string, coinSrc map[string]interface{}) func(c *coins.Coin) {
	return func(c *coins.Coin) {
		if coinSrc["ImageUrl"] != nil {
			imgurl := fmt.Sprint(coinSrc["ImageUrl"].(string))
			if imgurl != "<nil>" {
				c.SetLogo("https://cryptocompare.com" + imgurl)
			}
		}
		c.SetName(coinSrc["CoinName"])
		c.SetTicker(coinSrc["Symbol"])
		c.SetDescription(coinSrc["Description"])
		c.SetAlgo(coinSrc["Algorithm"])
		c.SetProof(coinSrc["ProofType"])
		c.SetStart(coinSrc["AssetLaunchDate"])

	}
}
