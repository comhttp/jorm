package xsrc

import "github.com/comhttp/jorm/app/jdb"

// GetCoinSources updates the available coin information sources
func GetExchangeSources(j *jdb.JDB) {
	getPoloniexExchange(j)
	//getDigiFinexExchange()
	//getBitTrexExchange()
	getBinanceExchange(j)
	//getHuobiExchange()
	//getGateExchange()
	//getGeminiExchange()
	getDexTradeExchange(j)
	//getBitZExchange()
	//getLiquidExchange()
	//go getCoinBeneExchange()
	//go getBitMartExchange()
	//go getHitBTCExchange()
	//go getNovaExchange()
	//go getIDAXExchange()
	//go getKuCoinExchange()
	//go getLBankExchange()
	return
}
