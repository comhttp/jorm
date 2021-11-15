package coin

import (
	"github.com/comhttp/jorm/pkg/jdb"
	"github.com/rs/zerolog/log"
)

func Queries(j *jdb.JDB, col string) *CoinsQueries {
	return &CoinsQueries{
		j,
		col,
	}
}

func (cq *CoinsQueries) SetCoin(src, slug string, get func(c Coin)) {
	c, err := cq.getCoin(slug)
	if err != nil {
		c = NewCoin(slug)
		log.Print("Insert Coin: ", slug)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		get(*c)
		c.Checked[src] = true
		//cq.WriteCoin(slug, c)
		//utl.ErrorLog(err)
	} else {
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		if !c.Checked[src] {
			log.Print("Check Coin: ", c.Slug)
			get(*c)
			c.Checked[src] = true
		} else {
			get(*c)
			log.Print("Already checked Coin: ", c.Slug)
		}
		//cq.WriteCoin(slug, c)

	}
	return
}

func (cq *CoinsQueries) WriteCoin(slug string, c interface{}) error {
	return cq.j.Write("coin", slug, c)
}

func (cq *CoinsQueries) WriteLogo(slug string, c interface{}) error {
	return cq.j.Write("logo", slug, c)
}

func (cq *CoinsQueries) WriteInfo(slug string, c interface{}) error {
	return cq.j.Write("info", slug, c)
}
