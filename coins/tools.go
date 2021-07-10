package coins

import (
	"fmt"
	"github.com/comhttp/jorm/app/jdb"
)

func ProcessCoins(j *jdb.JDB) {
	usableCoins := Coins{N: 0}
	//nodeCoins := NodeCoins{N: 0}
	algoCoins := AlgoCoins{N: 0}
	coinsWords := CoinsWords{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	c := GetCoins(j)
	fmt.Println("cccccccccc::", c)

	//var bitNodes []string
	//if err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", &bitNodes); err != nil {
	//	fmt.Println("Error", err)
	//}
	//for _, coiNn := range c {
	//	for _, bitNode := range bitNodes {
	//		if bitNode == coiNn.Slug {
	//			coiNn.BitNode = true
	//			jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/coins"), coiNn.Slug, coiNn)
	//		}
	//	}
	//}

	for i, slug := range c.C {
		coin := GetCoin(j, slug)
		//if coin.BitNode {
		//	nodeCoins.N++
		//	nodeCoins.C = append(nodeCoins.C, NodeCoin{
		//		Rank:   coin.Rank,
		//		Name:   coin.Name,
		//		Ticker: coin.Ticker,
		//		Slug:   coin.Slug,
		//		Algo:   coin.Algo,
		//	})
		//} else {
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
					fmt.Println("111aaa", a)
					fmt.Println("111coin.Algo", coin.Algo)
				}
				fmt.Println("222aaa", a)
				fmt.Println("2222coin.Algo", coin.Algo)
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
		//}
		allCoins.N = i
		allCoins.C = append(allCoins.C, coin.Slug)
	}

	//jdb.JDB.Write("jorm/info", "bitnodes", LoadCoinsBase(true, true))
	//jdb.JDB.Write(cfg.C.Out+"/info", "nodecoins", nodeCoins)
	//jdb.JDB.Write(cfg.C.Out+"/info", "coinsbin", coinsBin)

	j.Write("coins", "info_rest", restCoins)
	j.Write("coins", "info_algo", algoCoins)
	j.Write("coins", "info_words", coinsWords)
	j.Write("coins", "info_usable", usableCoins)
	j.Write("coins", "info_all", allCoins)

}
