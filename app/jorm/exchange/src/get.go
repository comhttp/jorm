package xsrc

// GetCoinSources updates the available coin information sources
func GetExchangeSources() {
	getPoloniexExchange()
	getDigiFinexExchange()
	getBitTrexExchange()
	getBinanceExchange()
	getHuobiExchange()
	getGateExchange()
	getGeminiExchange()
	getDexTradeExchange()
	getBitZExchange()
	getLiquidExchange()
	//go getCoinBeneExchange()
	//go getBitMartExchange()
	//go getHitBTCExchange()
	//go getNovaExchange()
	//go getIDAXExchange()
	//go getKuCoinExchange()
	//go getLBankExchange()
	return
}
