package coins

import (
	"fmt"
	"github.com/comhttp/jorm/pkg/jdb"
)

//func BitNodeCoins(c nodes.NodeCoins, j *jdb.JDB) {
//	fmt.Println("Start Process BitNodes Coins")
//	nodeCoins := nodes.NodeCoins{N: 0}
//	//c := GetAllCoins(j)
//
//	for _, nodeCoin := range c.C {
//		coin := getCoin(j, nodeCoin.Slug)
//		fmt.Println("Bitnode Coin: ", coin.Name)
//		bitNodes := nodes.BitNodes{}
//		if err := cfg.CFG.Read("nodes", coin.Slug, &bitNodes); err != nil {
//			fmt.Println("Error", err)
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

func ProcessCoins(j *jdb.JDB) {
	fmt.Println("Start ProcessCoins")

	usableCoins := Coins{N: 0}
	algoCoins := AlgoCoins{N: 0}
	coinsWords := CoinsWords{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	c := GetCoins(j)
	for i, slug := range c.C {
		coin := getCoin(j, slug)
		if coin.Algo != "N/A" && coin.Algo != "" {
			algoCoins.N++
			algoCoins.C = append(algoCoins.C, AlgoCoin{
				Rank:   coin.Rank,
				Name:   coin.Name,
				Ticker: coin.Ticker,
				Slug:   coin.Slug,
				Algo:   coin.Algo,
			})
			for _, a := range algoCoins.A {
				if a != coin.Algo {
					algoCoins.A = append(algoCoins.A, coin.Algo)
				}
			}
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
		coinsWords.C = coinsWords.C + " " + coin.Name
		coinsWords.N = usableCoins.N
		allCoins.N = i
		allCoins.C = append(allCoins.C, coin.Slug)
	}

	j.Write("info", "restcoins", restCoins)
	j.Write("info", "algocoins", algoCoins)
	j.Write("info", "wordscoins", coinsWords)
	j.Write("info", "usablecoins", usableCoins)
	j.Write("info", "allcoins", allCoins)
	j.Write("info", "bincoins", coinsBin)
}
