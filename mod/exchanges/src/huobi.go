package xsrc

func getHuobiExchange() {
	//fmt.Println("Get Huobi Exchange Start")
	//exchangeRaw := make(map[string]interface{})
	//
	//slug := "huobi"
	//var e exchange.Exchange
	//e.Name = "Huobi"
	//e.Slug = slug
	//resps, err := http.Get("https://api.huobi.pro/v1/common/symbols")
	//if err != nil {
	//}
	//defer resps.Body.Close()
	//mapBodyS, err := ioutil.ReadAll(resps.Body)
	//json.Unmarshal(mapBodyS, &exchangeRaw)
	//
	//if exchangeRaw["data"] != nil {
	//	markets := exchangeRaw["data"].([]interface{})
	//	tickersRaw := make(map[string]interface{})
	//	respt, err := http.Get("https://api.huobi.pro/market/tickers")
	//	if err != nil {
	//	}
	//	defer respt.Body.Close()
	//	mapBodyT, err := ioutil.ReadAll(respt.Body)
	//	json.Unmarshal(mapBodyT, &tickersRaw)
	//
	//	tickers := make(map[string]map[string]interface{})
	//	if tickersRaw["Data"] != nil {
	//
	//		for _, ticker := range tickersRaw["Data"].([]interface{}) {
	//			if ticker != nil {
	//				t := ticker.(map[string]interface{})
	//				if t["symbol"] != nil {
	//					tickers[strings.ToUpper(t["symbol"].(string))] = t
	//				}
	//			}
	//		}
	//	}
	//	e.Markets = make(map[string]exchange.Market)
	//	for _, marketSrc := range markets {
	//		m := marketSrc.(map[string]interface{})
	//		if q := m["quote-currency"]; q != nil {
	//			if nq := strings.ToUpper(q.(string)); nq != e.Markets[nq].Symbol {
	//				e.Markets[nq] = exchange.Market{
	//					Symbol:     nq,
	//					Currencies: make(map[string]exchange.Currency),
	//				}
	//			}
	//			if m["symbol"] != nil {
	//				s := m["symbol"].(string)
	//				q := strings.ToUpper(m["quote-currency"].(string))
	//
	//				e.SetCurrencyMarket(
	//					q,
	//					strings.ToUpper(m["base-currency"].(string)),
	//					tickers[s]["ask"],
	//					tickers[s]["bid"],
	//					tickers[s]["high"],
	//					tickers[s]["close"],
	//					tickers[s]["low"],
	//					tickers[s]["vol"])
	//				//Symbol: strings.ToUpper(marketSrc.BaseCurrency),
	//				// Ask:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Ask),
	//				// Bid:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Bid),
	//				//	High: fmt.Sprintf("%f", tickers[marketSrc.Symbol].High),
	//				// Last:   fmt.Sprintf("%f", tickers[marketSrc.Symbol].Last),
	//				//	Low:    fmt.Sprintf("%f", tickers[marketSrc.Symbol].Low),
	//				//Volume: fmt.Sprintf("%f", tickers[marketSrc.Symbol].Vol),
	//
	//			}
	//		}
	//	}
	//	jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//	fmt.Println("Get Huobi Exchange Done")
	//}
}
