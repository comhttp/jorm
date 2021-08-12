package xsrc

func getDigiFinexExchange() {
	//log.Println("Get DigiFinex Exchange Start")
	//marketsRaw := make(map[string]interface{})
	//slug := "digifinex"
	//var e exchange.Exchange
	//e.Name = "DigiFinex"
	//e.Slug = slug
	//respcs, err := http.Get("https://openapi.digifinex.com/v3/ticker")
	//utl.ErrorLog(err)
	//defer respcs.Body.Close()
	//mapBody, err := ioutil.ReadAll(respcs.Body)
	//if mapBody != nil {
	//	json.Unmarshal(mapBody, &marketsRaw)
	//	e.Markets = make(map[string]exchange.Market)
	//	if marketsRaw["ticker"] != nil {
	//		for _, marketSrc := range marketsRaw["ticker"].([]interface{}) {
	//			mSrc := marketSrc.(map[string]interface{})
	//			if mSrc["symbol"] != nil {
	//				m := strings.Split(mSrc["symbol"].(string), "_")
	//				if nq := strings.ToUpper(m[1]); nq != e.Markets[nq].Symbol {
	//					e.Markets[nq] = exchange.Market{
	//						Symbol:     nq,
	//						Currencies: make(map[string]exchange.Currency),
	//					}
	//				}
	//				e.SetCurrencyMarket(
	//					strings.ToUpper(m[1]),
	//					strings.ToUpper(m[0]),
	//					mSrc["sell"],
	//					mSrc["buy"],
	//					mSrc["high"],
	//					mSrc["last"],
	//					mSrc["low"],
	//					mSrc["vol"])
	//			}
	//		}
	//		jdb.JDB.Write(cfg.C.Out+"/exchanges", e.Slug, e)
	//		log.Println("Get DigiFinex Exchange Done")
	//	}
	//} else {
	//	log.Println("Get Gate Exchange Fail")
	//}
}
