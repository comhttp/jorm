package coins

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"log"
	"strings"
)

func GetCoin(j *jdb.JDB, slug string) Coin {
	c := getCoin(j, slug)
	return c
}
func getCoin(j *jdb.JDB, key string) Coin {
	c := Coin{}
	err := j.Read("coins", "coin_"+key, &c)
	log.Println(err)
	return c
}

func GetCoins(j *jdb.JDB) Coins {
	coins, err := j.ReadAll("coins", "coin_")
	log.Println(err)
	allCoins := Coins{N: 0}
	for i, c := range coins {
		allCoins.C = append(allCoins.C, strings.TrimPrefix(c, "coin_"))
		allCoins.N = i
	}
	return allCoins
}
func GetAllCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "allcoins", &c)
	log.Println(err)
	return c
}

//func GetNodeCoins(j *jdb.JDB) nodes.NodeCoins {
//	c := nodes.NodeCoins{}
//	err := j.Read("info", "nodecoins", &c)
//	log.Println(err)
//	return c
//}

func GetRestCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "restcoins", &c)
	log.Println(err)
	return c
}
func GetCoinsWords(j *jdb.JDB) CoinsWords {
	c := CoinsWords{}
	err := j.Read("info", "wordscoins", &c)
	log.Println(err)
	return c
}
func GetAlgoCoins(j *jdb.JDB) AlgoCoins {
	c := AlgoCoins{}
	err := j.Read("info", "algocoins", &c)
	log.Println(err)
	return c
}
func GetUsableCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "usablecoins", &c)
	log.Println(err)
	return c
}
func GetCoinsBin(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "bincoins", &c)
	log.Println(err)
	return c
}
