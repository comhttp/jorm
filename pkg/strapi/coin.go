package strapi



//type StrapiCoin struct {
//	coin.Coin
//}
//
//func (sc *StrapiCoin) New(params map[string]interface{}) {
//
//	sc.ID = int(params["id"].(float64))
//	sc.Price = params["price"].(float64)
//
//	if val, ok := params["title"]; ok && val != nil {
//		sc.Title = val.(string)
//	}
//
//	if val, ok := params["created_at"]; ok && val != nil {
//		createdAt, t1Err := time.Parse(time.RFC3339, val.(string))
//		if t1Err != nil {
//			sc.CreatedAt = createdAt
//		}
//	}
//
//
//	//sc.Image = make([]StrapiProductImage, 0)
//	//
//	//if images, ok := params["image"]; ok && images != nil {
//	//	for _, e := range images.([]interface{}) {
//	//		imageMap := e.(map[string]interface{})
//	//		var im StrapiProductImage
//	//		im.New(imageMap)
//	//		sc.Image = append(sc.Image, im)
//	//	}
//	//}
//
//}
