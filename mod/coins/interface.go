package coins

import (
	jdb2 "github.com/comhttp/jorm/pkg/jdb"
)

type JormCoins interface {
	SetCoin(j *jdb2.JDB, src, slug string, get func(c *Coin))
	SetName(name interface{})
	SetTicker(ticker interface{})
	SetAlgo(algo interface{})
	SetDescription(description interface{})
	SetProof(proof interface{})
	SetStart(start interface{})
	SetBitcoinTalk(bitcointalk interface{})
	SetTwitter(twitter interface{})
	SetReddit(reddit interface{})
	SetTelegram(telegram interface{})
	SetFacebook(facebook interface{})
	SetWebSite(website interface{})
	SetExplorer(explorer interface{})
	SetChat(chat interface{})
	SetLogo(logo interface{})
}
