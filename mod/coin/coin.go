package coin

//func LoadLogo(slug, size string) image.Image {
//	// Load logo image from database
//	logos := make(map[string]interface{})
//	log.Print("slug", slug)
//	err := jdb.JDB.Read(filepath.FromSlash(cfg.C.Out+"/data/"+slug), "logo", logos)
//	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(logos[size].(string)))
//	logo, _, err := image.Decode(reader)
//	utl.ErrorLog(err)
//	return logo
//}
//
//func LoadInfo(slug string) Coin {
//	// Load coin data from database
//	info := Coin{}
//	err := jdb.JDB.Read(filepath.FromSlash("data/"+slug), "info", info)
//	utl.ErrorLog(err)
//	//jsonString, _ := json.Marshal(info)
//
//	// convert json to struct
//	//s := CoinData{}
//	//json.Unmarshal(jsonString, &s)
//	return info
//}
