package coin

import (
	"encoding/base64"
	"time"

	"github.com/comhttp/jorm/mod/index"
	"github.com/comhttp/jorm/pkg/strapi"
	"github.com/comhttp/jorm/pkg/utl"
	"github.com/rs/zerolog/log"
)

func (cq *CoinsQueries) ProcessCoins(s strapi.StrapiRestClient) {
	log.Print("Start Process Coins")
	coins := s.GetAll("coins")
	subdomain := s.GetAll("subdomain")
	index.SetIndex(s, "coins", coins, nil)
	index.SetIndex(s, "allcoins", coins, SetCoinsIndex())

	logos := s.GetAll("logos")
	index.SetIndex(s, "logos", logos, nil)

	var logocoins []map[string]interface{}
	var algocoins []map[string]interface{}

	usableCoins := Coins{N: 0}
	algoCoins := AlgoCoins{N: 0}

	coinsWords := Coins{N: 0}
	restCoins := Coins{N: 0}

	coinsBin := Coins{N: 0}
	allCoins := Coins{N: 0}
	for i, c := range coins {

		cq.WriteCoin(c["slug"].(string), c)

		if c["subdomain"].(bool) {
			subdomain = append(subdomain, c)
		}
		if c["algo"].(string) != "" &&
			c["algo"].(string) != "N/A" &&
			c["symbol"].(string) != "" &&
			//coin.NetworkHashrate != 0 &&
			// coin.BlockHeight != 0 &&
			// coin.Difficulty != 0 &&
			c["name"].(string) != "" &&
			c["description"].(string) != "" {

			algoCoins.N++
			algoCoins.C = append(algoCoins.C, CoinShort{
				Rank:   int(c["rank"].(float64)),
				Name:   c["name"].(string),
				Symbol: c["symbol"].(string),
				Slug:   c["slug"].(string),
				Algo:   c["algo"].(string),
			})
			algoCoins.A = append(algoCoins.A, c["algo"].(string))
			algocoins = append(algocoins, c)
			// } else {
			// fmt.Println("cname else  :::", c["name"].(string))
			// if c["description"].(string) != "" {
			// 	//len(c[i].WebSite) > 0 &&
			// 	// len(coin.WebSite) > 0 &&
			// 	//if c[i].Platform != "token" &&
			// 	restCoins.N++
			// 	restCoins.C = append(restCoins.C, c["slug"].(string))
			// } else {

			// 	fmt.Println("descriptiondescriptiondescription ccc :   ", c)

			// 	coinsBin.N++
			// 	coinsBin.C = append(coinsBin.C, c["slug"].(string))
			// }
		}

		usableCoins.N = i
		usableCoins.C = append(usableCoins.C, c["slug"].(string))
		coinsWords.C = append(coinsWords.C, c["name"].(string))
		coinsWords.N = usableCoins.N
		allCoins.N = i
		allCoins.C = append(allCoins.C, c["slug"].(string))

		if index.CheckIndex(c["slug"].(string), logos[0]) {
			logocoins = append(logocoins, c)
		}

		time.Sleep(99 * time.Microsecond)
	}

	algoCoins.A = utl.RemoveDuplicateStr(algoCoins.A)

	var algoCoinsLogo AlgoCoinsLogo
	for _, ac := range algoCoins.C {
		logo := base64.StdEncoding.EncodeToString(cq.getLogo(ac.Slug, 32))
		if logo != "" {
			algoCoinsLogo.C = append(algoCoinsLogo.C, CoinShortLogo{
				Rank:   ac.Rank,
				Name:   ac.Name,
				Symbol: ac.Symbol,
				Slug:   ac.Slug,
				Algo:   ac.Algo,
				Logo:   "data:image/png;base64," + logo,
			})
		}
	}
	algoCoinsLogo.A = algoCoins.A
	algoCoinsLogo.N = algoCoins.N

	// fmt.Println("algoCoinsAAAA :   ", algoCoins.A)
	// fmt.Println("algoCoins :   ", algoCoins)
	// fmt.Println("coinsWords :   ", coinsWords)
	// fmt.Println("usableCoins :   ", usableCoins)
	// fmt.Println("allCoins :   ", allCoins)
	// fmt.Println("coinsBin :   ", coinsBin)
	index.SetIndex(s, "subdomain", subdomain, SetCoinsIndex())

	index.SetIndex(s, "allcoinslogo", logocoins, cq.SetCoinsLogoIndex(s))

	index.SetIndex(s, "algocoins", algocoins, SetCoinsIndex())
	index.SetIndex(s, "algocoinslogo", algocoins, cq.SetCoinsLogoIndex(s))

	cq.WriteInfo("restcoins", restCoins)
	cq.WriteInfo("algocoins", algoCoins)
	cq.WriteInfo("algocoinslogo", algoCoinsLogo)
	cq.WriteInfo("wordscoins", coinsWords)
	cq.WriteInfo("usablecoins", usableCoins)
	cq.WriteInfo("allcoins", allCoins)
	cq.WriteInfo("bincoins", coinsBin)

	log.Print("End ProcessCoins")
}