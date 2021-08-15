package cryptocompare

type cryptocompare struct {
	apiEndpoint, apiKey string
}

func NewCryptoCompareAPI(apiKey string) *cryptocompare {
	return &cryptocompare{"https://min-api.cryptocompare.com/", apiKey}
	//fsyms := "BTC"
	//tsyms := "USD"
	//apiKey = "30054c457ea44e183bb614813c325675d53f4eef2151406b4745f8baaeeaa381"
}
