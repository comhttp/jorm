package xsrc

func getLiquidExchange() {
	//fmt.Println("Get Liquid Exchange Start")
	//marketsRaw := []interface{}{}
	//slug := "liquid"
	//var e exchange.Exchange
	//e.Name = "Liquid"
	//e.Slug = slug
	//respcs, err := http.Get("https://api.liquid.com/products")
	//utl.ErrorLog(err)
	//
	//defer respcs.Body.Close()
	//mapBody, err := ioutil.ReadAll(respcs.Body)
	//json.Unmarshal(mapBody, &marketsRaw)
	//e.Markets = make(map[string]exchange.Market)
	//if marketsRaw != nil {
	//	for _, marketSrc := range marketsRaw {
	//		mSrc := marketSrc.(map[string]interface{})
	//		if mSrc != nil {
	//			if nq := mSrc["quoted_currency"].(string); nq != e.Markets[nq].Symbol {
	//				e.Markets[nq] = exchange.Market{
	//					Symbol:     mSrc["base_currency"].(string),
	//					Currencies: make(map[string]exchange.Currency),
	//				}
	//			}
	//			e.SetCurrencyMarket(
	//				mSrc["quoted_currency"].(string),
	//				mSrc["base_currency"].(string),
	//				mSrc["market_ask"],
	//				mSrc["market_bid"],
	//				mSrc["high_market_ask"],
	//				mSrc["last_price_24h"],
	//				mSrc["low_market_bid"],
	//				mSrc["volume_24h"])
	//		}
	//	}
	//	jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//	fmt.Println("Get Liquid Exchange Done")
	//}
}
