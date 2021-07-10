package coin

//
//func GetCoin(j *jdb.JDB, src, slug string, get func(c *Coin)) {
//	c := Coin{}
//	//_, err := os.Stat(filepath.FromSlash(cfg.Path + "/" + cfg.C.Out + "/coins/" + slug))
//	err := j.Read("coins", "coins_"+slug, &c)
//	if err != nil {
//		c.media.Slug = slug
//		fmt.Println("Insert Coin: ", slug)
//		if c.Checked == nil {
//			c.Checked = make(map[string]bool)
//		}
//		get(&c)
//		c.Checked[src] = true
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/coins"), slug, c)
//		j.Write("coins", "coins_"+slug, c)
//	} else {
//		//err = jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/coins"), slug, &c)
//		utl.ErrorLog(err)
//		fmt.Println("Ima Coin: ", c.Name)
//		if c.Checked == nil {
//			c.Checked = make(map[string]bool)
//		}
//		if !c.Checked[src] {
//			fmt.Println("Check Coin: ", c.Name)
//			get(&c)
//			c.Checked[src] = true
//		}
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/coins"), slug, c)
//		j.Write("coins", "coins_"+slug, c)
//	}
//	return
//}
//
//func (c *Coin) SetName(name interface{}) {
//	c.Name = utl.InsertString(c.Name, name)
//	return
//}
//
//func (c *Coin) SetTicker(ticker interface{}) {
//	c.Ticker = strings.ToUpper(utl.InsertString(c.Ticker, ticker))
//	return
//}
//
//func (c *Coin) SetAlgo(algo interface{}) {
//	c.Algo = utl.InsertString(c.Algo, algo)
//	return
//}
//
//func (c *Coin) SetDescription(description interface{}) {
//	c.Description = utl.InsertString(c.Description, description)
//	return
//}
//
//func (c *Coin) SetProof(proof interface{}) {
//	c.Proof = utl.InsertString(c.Proof, proof)
//	return
//}
//
//func (c *Coin) SetStart(start interface{}) {
//	c.Start = utl.InsertString(c.Start, start)
//	return
//}
//
//func (c *Coin) SetBitcoinTalk(bitcointalk interface{}) {
//	c.BitcoinTalk = utl.InsertString(c.BitcoinTalk, bitcointalk)
//	return
//}
//
//func (c *Coin) SetTwitter(twitter interface{}) {
//	c.Twitter = utl.InsertString(c.Twitter, twitter)
//	return
//}
//
//func (c *Coin) SetReddit(reddit interface{}) {
//	c.Reddit = utl.InsertString(c.Reddit, reddit)
//	return
//}
//
//func (c *Coin) SetTelegram(telegram interface{}) {
//	c.Telegram = utl.InsertString(c.Telegram, telegram)
//	return
//}
//func (c *Coin) SetFacebook(facebook interface{}) {
//	c.Facebook = utl.InsertString(c.Facebook, facebook)
//	return
//}
//
//func (c *Coin) SetWebSite(website interface{}) {
//	c.WebSite = utl.InsertStringSlice(c.WebSite, website)
//	return
//}
//
//func (c *Coin) SetExplorer(explorer interface{}) {
//	c.Explorer = utl.InsertStringSlice(c.Explorer, explorer)
//	return
//}
//
//func (c *Coin) SetChat(chat interface{}) {
//	c.Chat = utl.InsertStringSlice(c.Chat, chat)
//	return
//}
//
//func (c *Coin) SetLogo(logo interface{}) {
//	if logo.(string) != "" && logo.(string) != "missing_large.png" {
//		//imgs := utl.GetIMG(logo.(string), cfg.Path+cfg.C.Out+"/imgs/", c.Slug)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "all", imgs)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "16", imgs.Img16)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "32", imgs.Img32)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "64", imgs.Img64)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "128", imgs.Img128)
//		//jdb.JDB.Write(filepath.FromSlash(cfg.C.Out+"/imgs/"+c.Slug+"/base64/"), "256", imgs.Img256)
//		//Create a empty file
//		c.Logo = true
//	}
//	return
//}
