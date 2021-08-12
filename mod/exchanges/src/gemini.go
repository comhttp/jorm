package xsrc

func getGeminiExchange() {
	//log.Println("Get Gemini Exchange Start")
	//var symbolsRaw []string
	//slug := "gemini"
	//var e exchange.Exchange
	//e.Name = "Gemini"
	//e.Slug = slug
	//resps, err := http.Get("https://api.gemini.com/v1/symbols")
	//log.Println(err)
	//defer resps.Body.Close()
	//mapBodyS, err := ioutil.ReadAll(resps.Body)
	//if mapBodyS != nil {
	//	json.Unmarshal(mapBodyS, &symbolsRaw)
	//	e.Markets = make(map[string]exchange.Market)
	//	if symbolsRaw != nil {
	//		for _, symbol := range symbolsRaw {
	//			detailsRaw := make(map[string]interface{})
	//			respDetails, err := http.Get("https://api.gemini.com/v1/symbols/details/" + symbol)
	//			log.Println(err)
	//			defer respDetails.Body.Close()
	//			mapBodyDetails, err := ioutil.ReadAll(respDetails.Body)
	//			if mapBodyDetails != nil {
	//				json.Unmarshal(mapBodyDetails, &detailsRaw)
	//				if q := detailsRaw["quote_currency"]; q != nil {
	//					if nq := q.(string); nq != e.Markets[nq].Symbol {
	//						e.Markets[nq] = exchange.Market{
	//							Symbol:     nq,
	//							Currencies: make(map[string]exchange.Currency),
	//						}
	//					}
	//				}
	//				tickerRaw := make(map[string]interface{})
	//				respTicker, err := http.Get("https://api.gemini.com/v2/ticker/" + symbol)
	//				log.Println(err)
	//				defer respTicker.Body.Close()
	//				mapBodyTicker, err := ioutil.ReadAll(respTicker.Body)
	//				if mapBodyTicker != nil {
	//					json.Unmarshal(mapBodyTicker, &tickerRaw)
	//					tickerVolRaw := make(map[string]interface{})
	//					respVolTicker, err := http.Get("https://api.gemini.com/v1/pubticker/" + symbol)
	//					log.Println(err)
	//					defer respVolTicker.Body.Close()
	//					mapBodyVolTicker, err := ioutil.ReadAll(respVolTicker.Body)
	//					if mapBodyVolTicker != nil {
	//						json.Unmarshal(mapBodyVolTicker, &tickerVolRaw)
	//						if detailsRaw["quote_currency"] != nil && detailsRaw["base_currency"] != nil && tickerVolRaw["volume"] != nil {
	//							volume := tickerVolRaw["volume"].(map[string]interface{})
	//							e.SetCurrencyMarket(
	//								detailsRaw["quote_currency"].(string),
	//								detailsRaw["base_currency"].(string),
	//								tickerRaw["ask"],
	//								tickerRaw["bid"],
	//								tickerRaw["high"],
	//								tickerVolRaw["last"],
	//								tickerRaw["low"],
	//								volume[detailsRaw["quote_currency"].(string)])
	//						}
	//					}
	//				}
	//			}
	//		}
	//		jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//		log.Println("Get Gemini Exchange Done")
	//	}
	//} else {
	//	log.Println("Get Gemini Exchange Fail")
	//}
}
