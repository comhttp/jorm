package coin

import (
	"github.com/comhttp/jorm/pkg/utl"
)

func (cq *CoinsQueries) GetCoin(slug string) (Coin, error) {
	c, err := cq.getCoin(slug)
	return *c, err
}
func (cq *CoinsQueries) GetCoinShort(slug string) (CoinShort, error) {
	c, err := cq.getCoin(slug)
	cs := &CoinShort{
		Rank:   c.Rank,
		Name:   c.Name,
		Symbol: c.Symbol,
		Slug:   c.Slug,
		Algo:   c.Algo,
	}
	return *cs, err
}

func (cq *CoinsQueries) getCoin(key string) (c *Coin, err error) {
	//err = cq.j.Read("coin", key, &c)
	//utl.ErrorLog(err)
	return c, err
}

func (cq *CoinsQueries) GetCoins() Coins {
	coins, err := cq.j.ReadAll("coin")
	utl.ErrorLog(err)
	allCoins := Coins{N: 0}
	for i, c := range coins {
		allCoins.C = append(allCoins.C, c)
		allCoins.N = i
	}
	return allCoins
}

func (cq *CoinsQueries) GetAllCoins() (c Coins) {
	err := cq.j.Read("info", "allcoins", &c)
	utl.ErrorLog(err)
	return c
}

func (cq *CoinsQueries) GetNodeCoins() (n Coins) {
	err := cq.j.Read("info", "nodecoins", &n)
	utl.ErrorLog(err)
	return n
}

func (cq *CoinsQueries) GetRestCoins() (c Coins) {
	err := cq.j.Read("info", "restcoins", &c)
	utl.ErrorLog(err)
	return c
}
func (cq *CoinsQueries) GetCoinsWords() (c Coins) {
	err := cq.j.Read("info", "wordscoins", &c)
	utl.ErrorLog(err)
	return c
}
func (cq *CoinsQueries) GetAlgoCoins() CoinsShort {
	c := CoinsShort{}
	err := cq.j.Read("info", "algocoins", &c)
	utl.ErrorLog(err)
	return c
}
func (cq *CoinsQueries) GetUsableCoins() Coins {
	c := Coins{}
	err := cq.j.Read("info", "usablecoins", &c)
	utl.ErrorLog(err)
	return c
}
func (cq *CoinsQueries) GetCoinsBin() Coins {
	c := Coins{}
	err := cq.j.Read("info", "bincoins", &c)
	utl.ErrorLog(err)
	return c
}
