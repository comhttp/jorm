package coin

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/comhttp/jorm/pkg/utl/img"
)

//func BitNodeCoins(c nodes.NodeCoins, j *jdb.JDB) {
//	log.Print("Start Process BitNodes Coins")
//	nodeCoins := nodes.NodeCoins{N: 0}
//	//c := GetAllCoins(j)
//
//	for _, nodeCoin := range c.C {
//		coin := getCoin(j, nodeCoin.Slug)
//		log.Print("Bitnode Coin: ", coin.Name)
//		bitNodes := nodes.BitNodes{}
//		if err := cfg.CFG.Read("nodes", coin.Slug, &bitNodes); err != nil {
//			log.Print("Error", err)
//		}
//		nodeCoins.N++
//		nodeCoins.C = append(nodeCoins.C, nodes.NodeCoin{
//			Rank:   coin.Rank,
//			Name:   coin.Name,
//			Ticker: coin.Ticker,
//			Slug:   coin.Slug,
//			Algo:   coin.Algo,
//			Nodes:  bitNodes,
//		})
//		coin.BitNode = true
//		j.Write("coins", "coin_"+coin.Slug, coin)
//	}
//	j.Write("info", "nodecoins", nodeCoins)
//}
func SetCoinsIndex() func(c map[string]interface{}) *CoinShort {
	return func(c map[string]interface{}) *CoinShort {
		return &CoinShort{
			Rank:   int(c["rank"].(float64)),
			Name:   c["name"].(string),
			Symbol: c["symbol"].(string),
			Slug:   c["slug"].(string),
			Algo:   c["algo"].(string),
		}
	}
}
func (cq *CoinsQueries) SetCoinsLogoIndex(s strapi.StrapiRestClient) func(c map[string]interface{}) *CoinShortLogo {
	return func(c map[string]interface{}) *CoinShortLogo {
		var size float64 = 32
		l := &img.Logo{}
		err := s.Get("logos", c["slug"].(string), l)
		utl.ErrorLog(err)
		logoRaw, err := hex.DecodeString(l.Data)
		logoRawBytes, _ := img.ImageResize(logoRaw, img.Options{Width: size, Height: size})

		logo := base64.StdEncoding.EncodeToString(logoRawBytes)
		return &CoinShortLogo{
			Rank:   int(c["rank"].(float64)),
			Name:   c["name"].(string),
			Symbol: c["symbol"].(string),
			Slug:   c["slug"].(string),
			Algo:   c["algo"].(string),
			Logo:   "data:image/png;base64," + logo,
		}

	}
}

func coinUser(coin *Coin) CoinUser {
	return CoinUser{
		Id:                   coin.Id,
		Name:                 coin.Name,
		Slug:                 coin.Slug,
		Description:          coin.Description,
		Selected:             coin.Selected,
		Favorite:             coin.Favorite,
		UpdatedAt:            coin.UpdatedAt,
		Order:                coin.Order,
		Symbol:               coin.Symbol,
		Token:                coin.Token,
		Algo:                 coin.Algo,
		Proof:                coin.Proof,
		Ico:                  coin.Ico,
		BuiltOn:              coin.BuiltOn,
		GenesisDate:          coin.GenesisDate,
		NetworkHashrate:      coin.NetworkHashrate,
		MaxSupply:            coin.MaxSupply,
		TotalCoinsMined:      coin.TotalCoinsMined,
		BlockHeight:          coin.BlockHeight,
		BlockTime:            coin.BlockTime,
		Difficulty:           coin.Difficulty,
		DifficultyAdjustment: coin.DifficultyAdjustment,
		BlockReward:          coin.BlockReward,
		BlockRewardReduction: coin.BlockRewardReduction,
		Rank:                 coin.Rank,
		Platform:             coin.Platform,
		BitNode:              coin.BitNode,
	}
}
