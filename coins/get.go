package coins

import (
	"github.com/comhttp/jorm/app/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

func (c *Coin) GetxXXXX(proof interface{}) {
	c.Proof = utl.InsertString(c.Proof, proof)
	return
}

func GetCoin(j *jdb.JDB, slug string) *Coin {
	c := &Coin{}
	err := j.Read("coins", "coins_"+slug, &c)
	utl.ErrorLog(err)
	return c
}

func GetCoins(j *jdb.JDB) Coins {
	coins, err := j.ReadAll("coins", "coins_")
	utl.ErrorLog(err)
	allCoins := Coins{N: 0}
	for i, c := range coins {
		allCoins.C = append(allCoins.C, c)
		allCoins.N = i
	}
	return allCoins
}

//func LoadCoinsBase(j *jdb.JDB) Coins {
//	usableCoins := Coins{N: 0}
//	nodeCoins := NodeCoins{N: 0}
//	algoCoins := AlgoCoins{N: 0}
//	coinsWords := CoinsWords{N: 0}
//	restCoins := Coins{N: 0}
//
//	coinsBin := Coins{N: 0}
//	allCoins := Coins{N: 0}
//	cIN := GetCoins(j)
//	fmt.Println("cccccccccc::", cIN)
//
//	//var bitNodes []string
//	//if err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/info"), "bitnoded", &bitNodes); err != nil {
//	//	fmt.Println("Error", err)
//	//}
//	//for _, coiNn := range c {
//	//	for _, bitNode := range bitNodes {
//	//		if bitNode == coiNn.Slug {
//	//			coiNn.BitNode = true
//	//			jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/coins"), coiNn.Slug, coiNn)
//	//		}
//	//	}
//	//}
//
//	for i, coin := range cIN.C {
//		cOUT := Coin{}.GetCoin()
//		//if coin.BitNode {
//		//	nodeCoins.N++
//		//	nodeCoins.C = append(nodeCoins.C, NodeCoin{
//		//		Rank:   coin.Rank,
//		//		Name:   coin.Name,
//		//		Ticker: coin.Ticker,
//		//		Slug:   coin.Slug,
//		//		Algo:   coin.Algo,
//		//	})
//		//} else {
//			if coin.Algo != "N/A" && coin.Algo != "" {
//				algoCoins.N++
//				algoCoins.C = append(algoCoins.C, AlgoCoin{
//					Rank:   coin.Rank,
//					Name:   coin.Name,
//					Ticker: coin.Ticker,
//					Slug:   coin.Slug,
//					Algo:   coin.Algo,
//				})
//				for _, a := range algoCoins.A {
//					if a != coin.Algo {
//						algoCoins.A = append(algoCoins.A, coin.Algo)
//						fmt.Println("111aaa", a)
//						fmt.Println("111coin.Algo", coin.Algo)
//					}
//					fmt.Println("222aaa", a)
//					fmt.Println("2222coin.Algo", coin.Algo)
//				}
//			} else {
//				if coin.Description != "" {
//					//len(c[i].WebSite) > 0 &&
//					// len(coin.WebSite) > 0 &&
//					//if c[i].Platform != "token" &&
//					restCoins.N++
//					restCoins.C = append(restCoins.C, coin.Slug)
//				} else {
//					coinsBin.N++
//					coinsBin.C = append(coinsBin.C, coin.Slug)
//				}
//			}
//			usableCoins.N = i
//			usableCoins.C = append(usableCoins.C, coin.Slug)
//			coinsWords.C = coinsWords.C + " " + coin.Name
//			coinsWords.N = usableCoins.N
//		}
//		allCoins.N = i
//		allCoins.C = append(allCoins.C, coin.Slug)
//	//}
//	//jdb.JDB.Write(cfg.C.Out+"/info", "restcoins", restCoins)
//	//jdb.JDB.Write(cfg.C.Out+"/info", "algos", algoCoins)
//	//jdb.JDB.Write(cfg.C.Out+"/info", "coinswords", coinsWords)
//	//jdb.JDB.Write(cfg.C.Out+"/info", "usableinfo", usableCoins)
//	//jdb.JDB.Write(cfg.C.Out+"/info", "allcoins", allCoins)
//
//	//jdb.JDB.Write("jorm/info", "bitnodes", LoadCoinsBase(true, true))
//	//jdb.JDB.Write(cfg.C.Out+"/info", "nodecoins", nodeCoins)
//	//jdb.JDB.Write(cfg.C.Out+"/info", "coinsbin", coinsBin)
//	return c
//}
