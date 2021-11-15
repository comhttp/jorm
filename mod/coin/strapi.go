package coin

import (
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/comhttp/jorm/pkg/utl/img"
	"github.com/rs/zerolog/log"
)

func SetCoin(s strapi.StrapiRestClient, coins, logos []string, src, slug string, get func(c *Coin, l *img.Logo)) {
	// var cc []*Coin
	// err := s.Get("coins", slug, &cc)
	// utl.ErrorLog(err)
	// var ll []*img.Logo
	// err = s.Get("logos", slug, &ll)
	// utl.ErrorLog(err)
	l := &img.Logo{
		Slug: slug,
	}
	if len(coins) != 0 {
		c := &Coin{}
		err := s.Get("coins", coins[0], c)
		utl.ErrorLog(err)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		if !c.Checked[src] {
			log.Print("Check Coin: ", c.Slug)
			get(c, l)
			if len(logos) > 0 {
				l := &img.Logo{}
				err := s.Get("logos", logos[0], l)
				utl.ErrorLog(err)
				if l.Data == "" {
					s.Post("logos", l)
				}
			}
			c.Checked[src] = true
		} else {
			get(c, l)
			log.Print("Already checked Coin: ", c.Slug)
		}
		s.Put("coins", c)
	} else {
		c := NewCoin(slug)
		log.Print("Insert Coin: ", slug)
		if c.Checked == nil {
			c.Checked = make(map[string]bool)
		}
		get(c, l)
		c.Checked[src] = true
		s.Post("coins", c)
		s.Post("logos", l)

		coins = append(coins, slug)
		logos = append(logos, slug)

		s.Post("indices", coins)
		s.Post("indices", logos)

	}
	return
}

// func SetCoinsIndex(s strapi.StrapiRestClient, col string) {
// 	var new bool
// 	collection := s.GetAll("coins")
// 	indexRaw := make(map[string]CoinShort)
// 	err := s.Get("indices", col, &indexRaw)
// 	utl.ErrorLog(err)
// 	if len(indexRaw) == 0 {
// 		new = true
// 	}
// 	for _, c := range collection {
// 		slug := c["slug"].(string)
// 		coinRaw := CoinShort{
// 			Rank:   c["rank"].(int),
// 			Name:   c["name"].(string),
// 			Symbol: c["symbol"].(string),
// 			Slug:   c["slug"].(string),
// 			Algo:   c["algo"].(string),
// 		}
// 		if len(indexRaw) > 0 {
// 			if !checkIndex(slug, indexRaw) {
// 				indexRaw[slug] = coinRaw
// 			}
// 		} else {
// 			indexRaw[slug] = coinRaw
// 		}
// 	}
// 	bytesIndex, err := json.Marshal(indexRaw)
// 	if err != nil {
// 		panic(err)
// 	}
// 	index := map[string]interface{}{
// 		"slug":  col,
// 		"index": string(bytesIndex),
// 	}
// 	if new {
// 		s.Post("indices", index)
// 	} else {
// 		s.Put("indices", index)
// 	}
// 	fmt.Println("Indexing done for: ", col)
// 	return
// }

// func checkIndex(slug string, index map[string]CoinShort) (c bool) {
// 	if _, found := index[slug]; !found {
// 		c = true
// 	}
// 	return
// }
