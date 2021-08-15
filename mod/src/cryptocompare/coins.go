package cryptocompare

import (
	"github.com/comhttp/jorm/pkg/utl"
)

type rawAllCoins struct {
	Data map[string]Coin `json:"Data"`
}

type Coin struct {
	ID               string `json:"Id"`
	URL              string `json:"Url"`
	ImageURL         string `json:"ImageUrl"`
	ContentCreatedOn int    `json:"ContentCreatedOn"`
	Name             string `json:"Name"`
	Symbol           string `json:"Symbol"`
	CoinName         string `json:"CoinName"`
	FullName         string `json:"FullName"`
	Description      string `json:"Description"`
	AssetTokenStatus string `json:"AssetTokenStatus"`
	Algorithm        string `json:"Algorithm"`
	ProofType        string `json:"ProofType"`
	SortOrder        string `json:"SortOrder"`
	Sponsored        bool   `json:"Sponsored"`
	Taxonomy         struct {
		Access                  string `json:"Access"`
		FCA                     string `json:"FCA"`
		FINMA                   string `json:"FINMA"`
		Industry                string `json:"Industry"`
		CollateralizedAsset     string `json:"CollateralizedAsset"`
		CollateralizedAssetType string `json:"CollateralizedAssetType"`
		CollateralType          string `json:"CollateralType"`
		CollateralInfo          string `json:"CollateralInfo"`
	} `json:"Taxonomy"`
	Rating struct {
		Weiss struct {
			Rating                   string `json:"Rating"`
			TechnologyAdoptionRating string `json:"TechnologyAdoptionRating"`
			MarketPerformanceRating  string `json:"MarketPerformanceRating"`
		} `json:"Weiss"`
	} `json:"Rating"`
	IsTrading          bool   `json:"IsTrading"`
	TotalCoinsMined    int    `json:"TotalCoinsMined"`
	BlockNumber        int    `json:"BlockNumber"`
	NetHashesPerSecond int    `json:"NetHashesPerSecond"`
	BlockReward        int    `json:"BlockReward"`
	BlockTime          int    `json:"BlockTime"`
	AssetLaunchDate    string `json:"AssetLaunchDate"`
	MaxSupply          int    `json:"MaxSupply"`
	MktCapPenalty      int    `json:"MktCapPenalty"`
	IsUsedInDefi       int    `json:"IsUsedInDefi"`
	IsUsedInNft        int    `json:"IsUsedInNft"`
}

func (c *cryptocompare) GetAllCoins() map[string]Coin {
	allCoins := &rawAllCoins{}
	utl.GetSourceHeadersAPIkey(c.apiKey, c.apiEndpoint+"data/all/coinlist", allCoins)
	return allCoins.Data
}
