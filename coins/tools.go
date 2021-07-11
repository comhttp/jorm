package coins

import (
	"fmt"
	"github.com/comhttp/jorm/cfg"
	"github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"path/filepath"
)

func ProcessCoins(j *jdb.JDB) {
	usableCoins := Coins{N: 0}
	nodeCoins := NodeCoins{N: 0}
	algoCoins := AlgoCoins{N: 0}
	coinsWords := CoinsWords{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	c := GetCoins(j)

	for i, slug := range c.C {
		coin := getCoin(j, slug)
		if utl.FileExists(filepath.FromSlash(cfg.Path + "nodes/" + slug)) {
			coin.BitNode = true
		}

		if coin.BitNode {
			nodeCoins.N++
			nodeCoins.C = append(nodeCoins.C, NodeCoin{
				Rank:   coin.Rank,
				Name:   coin.Name,
				Ticker: coin.Ticker,
				Slug:   coin.Slug,
				Algo:   coin.Algo,
			})
		} else {
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
		}
		allCoins.N = i
		allCoins.C = append(allCoins.C, coin.Slug)
	}

	//jdb.JDB.Write("jorm/info", "bitnodes", LoadCoinsBase(true, true))

	j.Write("coins", "info_nodecoins", nodeCoins)
	j.Write("coins", "info_rest", restCoins)
	j.Write("coins", "info_algo", algoCoins)
	j.Write("coins", "info_words", coinsWords)
	j.Write("coins", "info_usable", usableCoins)
	j.Write("coins", "info_all", allCoins)
	j.Write("coins", "info_bin", coinsBin)

}
