package xsrc

// GetCoinSources updates the available coin information sources
func GetExchangeSources() {
	go getPoloniexExchange()
	go getDigiFinexExchange()
	go getBitTrexExchange()
	go getBinanceExchange()
	go getHuobiExchange()
	go getGateExchange()
	go getGeminiExchange()
	go getDexTradeExchange()
	go getBitZExchange()

	//go getCoinBeneExchange()
	//go getBitMartExchange()
	//go getHitBTCExchange()
	//go getNovaExchange()
	//go getIDAXExchange()
	//go getKuCoinExchange()
	//go getLBankExchange()
	return
}
