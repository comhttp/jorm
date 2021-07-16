package coins

import (
	nodes2 "github.com/comhttp/jorm/mod/nodes"
	jdb2 "github.com/comhttp/jorm/pkg/jdb"
	"github.com/comhttp/jorm/pkg/utl"
	"strings"
)

func GetCoin(j *jdb2.JDB, slug string) Coin {
	c := getCoin(j, slug)
	return c
}
func getCoin(j *jdb2.JDB, key string) Coin {
	c := Coin{}
	err := j.Read("coins", "coins_"+key, &c)
	utl.ErrorLog(err)
	return c
}

func GetCoins(j *jdb2.JDB) Coins {
	coins, err := j.ReadAll("coins", "coins_")
	utl.ErrorLog(err)
	allCoins := Coins{N: 0}
	for i, c := range coins {
		allCoins.C = append(allCoins.C, strings.TrimPrefix(c, "coins_"))
		allCoins.N = i
	}
	return allCoins
}
func GetAllCoins(j *jdb2.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "allcoins", &c)
	utl.ErrorLog(err)
	return c
}
func GetNodeCoins(j *jdb2.JDB) nodes2.NodeCoins {
	c := nodes2.NodeCoins{}
	err := j.Read("info", "nodecoins", &c)
	utl.ErrorLog(err)
	return c
}

func GetRestCoins(j *jdb2.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "restcoins", &c)
	utl.ErrorLog(err)
	return c
}
func GetCoinsWords(j *jdb2.JDB) CoinsWords {
	c := CoinsWords{}
	err := j.Read("info", "wordscoins", &c)
	utl.ErrorLog(err)
	return c
}
func GetAlgoCoins(j *jdb2.JDB) AlgoCoins {
	c := AlgoCoins{}
	err := j.Read("info", "algocoins", &c)
	utl.ErrorLog(err)
	return c
}
func GetUsableCoins(j *jdb2.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "usablecoins", &c)
	utl.ErrorLog(err)
	return c
}
func GetCoinsBin(j *jdb2.JDB) Coins {
	c := Coins{}
	err := j.Read("info", "bincoins", &c)
	utl.ErrorLog(err)
	return c
}
