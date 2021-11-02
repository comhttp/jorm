package cryptocompare

import (
	"fmt"

	"github.com/comhttp/jorm/mod/coin"
	"github.com/comhttp/jorm/pkg/strapi"
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
	BuiltOn          string `json:"BuiltOn"`
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
	IsTrading          bool    `json:"IsTrading"`
	TotalCoinsMined    float64 `json:"TotalCoinsMined"`
	Difficulty         float64 `json:"Difficulty"`
	BlockNumber        int     `json:"BlockNumber"`
	NetHashesPerSecond float64 `json:"NetHashesPerSecond"`
	BlockReward        float64 `json:"BlockReward"`
	BlockTime          float64 `json:"BlockTime"`
	AssetLaunchDate    string  `json:"AssetLaunchDate"`
	MaxSupply          float64 `json:"MaxSupply"`
	MktCapPenalty      float64 `json:"MktCapPenalty"`
	IsUsedInDefi       int     `json:"IsUsedInDefi"`
	IsUsedInNft        int     `json:"IsUsedInNft"`
}

func (c *cryptocompare) GetAllCoins(s strapi.StrapiRestClient) {
	allCoins := &rawAllCoins{}
	utl.GetSourceHeadersAPIkey(c.apiKey, c.apiEndpoint+"data/all/coinlist", allCoins)
	fmt.Println("::::::::::::::::::::::::::::::::START cryptocompare COINS:::::::::::::::::::::::::::::: ")
	for _, ccCoin := range allCoins.Data {
		if ccCoin.CoinName != "" {
			slug := utl.MakeSlug(ccCoin.CoinName)
			go coin.SetCoin(s, "cryptocompare", slug, getCryptoCompareCoin(ccCoin))
		}
	}
	fmt.Println("::::::::::::::::::::::::::::::::END cryptocompare COINS:::::::::::::::::::::::::::::: ")
	return
}

func getCryptoCompareCoin(ccCoin Coin) func(c *coin.Coin) {
	return func(c *coin.Coin) {
		if ccCoin.ImageURL != "" && ccCoin.ImageURL != "<nil>" {
			c.SetLogo("https://cryptocompare.com" + ccCoin.ImageURL)
			fmt.Println("ImageURL:  ", "https://cryptocompare.com"+ccCoin.ImageURL)
		}
		c.SetSrcID("cryptocompare", ccCoin.ID)
		c.SetName(ccCoin.CoinName)
		c.SetSymbol(ccCoin.Symbol)
		c.SetDescription(ccCoin.Description)
		c.SetAlgo(ccCoin.Algorithm)
		c.SetProof(ccCoin.ProofType)
		c.SetGenesisDate(ccCoin.AssetLaunchDate)
		c.SetBlockHeight(ccCoin.BlockNumber)
		c.SetNetworkHashrate(ccCoin.NetHashesPerSecond)
		c.SetMaxSupply(ccCoin.MaxSupply)
		c.SetTotalCoinsMined(ccCoin.TotalCoinsMined)
		c.SetBlockHeight(ccCoin.BlockNumber)
		c.SetBlockTime(int(ccCoin.BlockTime))
		c.SetBlockReward(ccCoin.BlockReward)
		c.SetDifficulty(ccCoin.Difficulty)
		c.SetBuiltOn(ccCoin.BuiltOn)
	}
}
