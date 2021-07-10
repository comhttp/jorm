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
