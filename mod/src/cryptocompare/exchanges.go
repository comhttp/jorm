package cryptocompare

import (
	"fmt"
	"github.com/comhttp/jorm/mod/exchange"
	"github.com/comhttp/jorm/pkg/utl"
)

type rawAllExchanges struct {
	Data map[string]Exchange `json:"Data"`
}

type Exchange struct {
	ID                 string   `json:"Id"`
	Name               string   `json:"Name"`
	URL                string   `json:"Url"`
	LogoURL            string   `json:"LogoUrl"`
	ItemType           []string `json:"ItemType"`
	CentralizationType string   `json:"CentralizationType"`
	InternalName       string   `json:"InternalName"`
	GradePoints        float64  `json:"GradePoints"`
	Grade              string   `json:"Grade"`
	GradePointsSplit   struct {
		Legal                    string `json:"Legal"`
		KYCAndTransactionRisk    string `json:"KYCAndTransactionRisk"`
		Team                     string `json:"Team"`
		DataProvision            string `json:"DataProvision"`
		AssetQualityAndDiversity string `json:"AssetQualityAndDiversity"`
		MarketQuality            string `json:"MarketQuality"`
		Security                 string `json:"Security"`
		NegativeReportsPenalty   string `json:"NegativeReportsPenalty"`
	} `json:"GradePointsSplit"`
	AffiliateURL      string `json:"AffiliateURL"`
	Country           string `json:"Country"`
	OrderBook         bool   `json:"OrderBook"`
	Trades            bool   `json:"Trades"`
	Description       string `json:"Description"`
	FullAddress       string `json:"FullAddress"`
	Fees              string `json:"Fees"`
	DepositMethods    string `json:"DepositMethods"`
	WithdrawalMethods string `json:"WithdrawalMethods"`
	Sponsored         bool   `json:"Sponsored"`
	Recommended       bool   `json:"Recommended"`
	Rating            struct {
		One        int     `json:"One"`
		Two        int     `json:"Two"`
		Three      int     `json:"Three"`
		Four       int     `json:"Four"`
		Five       int     `json:"Five"`
		Avg        float64 `json:"Avg"`
		TotalUsers int     `json:"TotalUsers"`
	} `json:"Rating"`
	SortOrder      string `json:"SortOrder"`
	TOTALVOLUME24H struct {
		BTC float64 `json:"BTC"`
	} `json:"TOTALVOLUME24H"`
	DISPLAYTOTALVOLUME24H struct {
		BTC string `json:"BTC"`
	} `json:"DISPLAYTOTALVOLUME24H"`
}

func (c *cryptocompare) GetAllExchanges(exchangeQueries *exchange.ExchangeQueries) {
	allExchanges := &rawAllExchanges{}
	utl.GetSourceHeadersAPIkey(c.apiKey, c.apiEndpoint+"data/exchanges/general", allExchanges)
	fmt.Println("::::::::::::::::::::::::::::::::START cryptocompare EXCHANGES:::::::::::::::::::::::::::::: ")
	for _, ccExchange := range allExchanges.Data {
		if ccExchange.Name != "" {
			slug := utl.MakeSlug(ccExchange.Name)
			exchangeQueries.SetExchange("cryptocompare", slug, getCryptoCompareExchange(ccExchange))
		}
	}
	fmt.Println("::::::::::::::::::::::::::::::::END cryptocompare EXCHANGES:::::::::::::::::::::::::::::: ")
	return
}

func getCryptoCompareExchange(ccExchange Exchange) func(e *exchange.Exchange) {
	return func(e *exchange.Exchange) {
		//if ccCoin.ImageURL != "" && ccCoin.ImageURL != "<nil>" {
		//	c.SetLogo("https://cryptocompare.com" + ccCoin.ImageURL)
		//}
		e.SetName(ccExchange.Name)
		e.SetDescription(ccExchange.Description)
		e.SetURL(ccExchange.URL)

	}
}
