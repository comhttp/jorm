package coin

import (
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/rs/zerolog/log"
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
func SetCoinsIndex() func(c map[string]interface{}) interface{} {
	return func(c map[string]interface{}) interface{} {
		return &CoinShort{
			Rank:   int(c["rank"].(float64)),
			Name:   c["name"].(string),
			Symbol: c["symbol"].(string),
			Slug:   c["slug"].(string),
			Algo:   c["algo"].(string),
		}
	}
}
func (cq *CoinsQueries) SetCoinsLogoIndex(s strapi.StrapiRestClient) func(c map[string]interface{}) interface{} {
	return func(c map[string]interface{}) interface{} {
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
func (cq *CoinsQueries) ProcessCoins(s strapi.StrapiRestClient) {
	log.Print("Start Process Coins")
	coins := s.GetAll("coins")

	s.SetIndex("coins", coins, nil)
	s.SetIndex("allcoins", coins, SetCoinsIndex())

	logos := s.GetAll("logos")
	s.SetIndex("logos", logos, nil)

	var logocoins []map[string]interface{}

	usableCoins := Coins{N: 0}
	algoCoins := AlgoCoins{N: 0}

	coinsWords := Coins{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	for i, c := range coins {

		cq.WriteCoin(c["slug"].(string), c)

		if c["algo"].(string) != "" &&
			c["algo"].(string) != "N/A" &&
			c["symbol"].(string) != "" &&
			//coin.NetworkHashrate != 0 &&
			// coin.BlockHeight != 0 &&
			// coin.Difficulty != 0 &&
			c["name"].(string) != "" &&
			c["description"].(string) != "" {

			algoCoins.N++
			algoCoins.C = append(algoCoins.C, CoinShort{
				Rank:   int(c["rank"].(float64)),
				Name:   c["name"].(string),
				Symbol: c["symbol"].(string),
				Slug:   c["slug"].(string),
				Algo:   c["algo"].(string),
			})
			algoCoins.A = append(algoCoins.A, c["algo"].(string))

			// } else {
			// fmt.Println("cname else  :::", c["name"].(string))
			// if c["description"].(string) != "" {
			// 	//len(c[i].WebSite) > 0 &&
			// 	// len(coin.WebSite) > 0 &&
			// 	//if c[i].Platform != "token" &&
			// 	restCoins.N++
			// 	restCoins.C = append(restCoins.C, c["slug"].(string))
			// } else {

			// 	fmt.Println("descriptiondescriptiondescription ccc :   ", c)

			// 	coinsBin.N++
			// 	coinsBin.C = append(coinsBin.C, c["slug"].(string))
			// }
		}

		usableCoins.N = i
		usableCoins.C = append(usableCoins.C, c["slug"].(string))
		coinsWords.C = append(coinsWords.C, c["name"].(string))
		coinsWords.N = usableCoins.N
		allCoins.N = i
		allCoins.C = append(allCoins.C, c["slug"].(string))

		if strapi.CheckIndex(c["slug"].(string), logos[0]) {
			logocoins = append(logocoins, c)
		}

		time.Sleep(99 * time.Microsecond)
	}

	algoCoins.A = utl.RemoveDuplicateStr(algoCoins.A)

	var algoCoinsLogo AlgoCoinsLogo
	for _, ac := range algoCoins.C {
		logo := base64.StdEncoding.EncodeToString(cq.getLogo(ac.Slug, 32))
		if logo != "" {
			algoCoinsLogo.C = append(algoCoinsLogo.C, CoinShortLogo{
				Rank:   ac.Rank,
				Name:   ac.Name,
				Symbol: ac.Symbol,
				Slug:   ac.Slug,
				Algo:   ac.Algo,
				Logo:   "data:image/png;base64," + logo,
			})
		}
	}
	algoCoinsLogo.A = algoCoins.A
	algoCoinsLogo.N = algoCoins.N

	// fmt.Println("algoCoinsAAAA :   ", algoCoins.A)
	// fmt.Println("algoCoins :   ", algoCoins)
	// fmt.Println("coinsWords :   ", coinsWords)
	// fmt.Println("usableCoins :   ", usableCoins)
	// fmt.Println("allCoins :   ", allCoins)
	// fmt.Println("coinsBin :   ", coinsBin)

	s.SetIndex("allcoinslogo", logocoins, cq.SetCoinsLogoIndex(s))

	cq.WriteInfo("restcoins", restCoins)
	cq.WriteInfo("algocoins", algoCoins)
	cq.WriteInfo("algocoinslogo", algoCoinsLogo)
	cq.WriteInfo("wordscoins", coinsWords)
	cq.WriteInfo("usablecoins", usableCoins)
	cq.WriteInfo("allcoins", allCoins)
	cq.WriteInfo("bincoins", coinsBin)

	log.Print("End ProcessCoins")
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
