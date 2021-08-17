package coin

import (
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

func (cq *CoinQueries) ProcessCoins() {
	log.Print("Start ProcessCoins")

	usableCoins := Coins{N: 0}
	algoCoins := CoinsShort{N: 0}
	coinsWords := Coins{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	c := cq.GetCoins()
	for i, slug := range c.C {
		coin := cq.getCoin(slug)
		if coin.Algo != "N/A" && coin.Algo != "" {
			algoCoins.N++
			algoCoins.C = append(algoCoins.C, CoinShort{
				Rank:   coin.Rank,
				Name:   coin.Name,
				Ticker: coin.Ticker,
				Slug:   coin.Slug,
				Algo:   coin.Algo,
			})
			//for _, a := range algoCoins.A {
			//	if a != coin.Algo {
			//		algoCoins.A = append(algoCoins.A, coin.Algo)
			//	}
			//}
		} else {
			if coin.Description != "" {
				//len(c[i].WebSite) > 0 &&
				// len(coin.WebSite) > 0 &&
				//if c[i].Platform != "token" &&
				restCoins.N++
				restCoins.C = append(restCoins.C, coin.Slug)
			} else {
				coinsBin.N++
				coinsBin.C = append(coinsBin.C, coin.Slug)
			}
		}
		usableCoins.N = i
		usableCoins.C = append(usableCoins.C, coin.Slug)
		coinsWords.C = append(coinsWords.C, coin.Name)
		coinsWords.N = usableCoins.N
		allCoins.N = i
		allCoins.C = append(allCoins.C, coin.Slug)
	}

	cq.j.Write("info", "restcoins", restCoins)
	cq.j.Write("info", "algocoins", algoCoins)
	cq.j.Write("info", "wordscoins", coinsWords)
	cq.j.Write("info", "usablecoins", usableCoins)
	cq.j.Write("info", "allcoins", allCoins)
	cq.j.Write("info", "bincoins", coinsBin)
}
