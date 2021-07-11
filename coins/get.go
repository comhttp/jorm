package coins

import (
	"github.com/comhttp/jorm/jdb"
	"github.com/comhttp/jorm/pkg/utl"
)

func GetCoin(j *jdb.JDB, slug string) Coin {
	c := getCoin(j, "coins_"+slug)
	return c
}
func getCoin(j *jdb.JDB, key string) Coin {
	c := Coin{}
	err := j.Read("coins", key, &c)
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
func GetAllCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("coins", "info_all", &c)
	utl.ErrorLog(err)
	return c
}
func GetNodeCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("coins", "info_nodes", &c)
	utl.ErrorLog(err)
	return c
}

func GetRestCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("coins", "info_rest", &c)
	utl.ErrorLog(err)
	return c
}
func GetCoinsWords(j *jdb.JDB) CoinsWords {
	c := CoinsWords{}
	err := j.Read("coins", "info_words", &c)
	utl.ErrorLog(err)
	return c
}
func GetAlgoCoins(j *jdb.JDB) AlgoCoins {
	c := AlgoCoins{}
	err := j.Read("coins", "info_algo", &c)
	utl.ErrorLog(err)
	return c
}
func GetUsableCoins(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("coins", "info_usable", &c)
	utl.ErrorLog(err)
	return c
}
func GetCoinsBin(j *jdb.JDB) Coins {
	c := Coins{}
	err := j.Read("coins", "info_bin", &c)
	utl.ErrorLog(err)
	return c
}
