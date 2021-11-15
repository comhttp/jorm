package strapi

import (
	"encoding/json"
	"fmt"

	"github.com/comhttp/jorm/pkg/utl"
)

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

func (s StrapiRestClient) SetIndex(slug string, collection []map[string]interface{}, typeFunc func(c map[string]interface{}) interface{}) {
	var new bool
	// collection := s.GetAll(col)
	indexRaw := s.GetIndex(slug)
	if indexRaw == nil {
		indexRaw = make(map[string]interface{})
	}
	if len(indexRaw) == 0 {
		new = true
	}
	ss := indexRaw
	for _, c := range collection {
		itemSlug := c["slug"].(string)
		if len(ss) > 0 {
			if CheckIndex(itemSlug, ss) {
				if typeFunc != nil {
					indexRaw[itemSlug] = typeFunc(c)
				} else {
					indexRaw[itemSlug] = true
				}
			}
		} else {
			if typeFunc != nil {
				indexRaw[itemSlug] = typeFunc(c)
			} else {
				indexRaw[itemSlug] = true
			}
		}
	}
	bytesIndex, err := json.Marshal(indexRaw)
	if err != nil {
		panic(err)
	}
	index := map[string]interface{}{
		"slug":  slug,
		"index": string(bytesIndex),
	}
	if new {
		s.Post("indices", index)
	} else {
		s.Put("indices", index)
	}
	fmt.Println("Indexing done for: ", slug)
	return
}

func (s StrapiRestClient) SetIndexItem(slug string, item map[string]interface{}) {
	var new bool
	indexRaw := s.GetIndex(slug)
	if indexRaw == nil {
		indexRaw = make(map[string]interface{})
	}
	if len(indexRaw) == 0 {
		new = true
	}
	ss := indexRaw
	itemSlug := item["slug"].(string)
	if len(ss) > 0 {
		if CheckIndex(itemSlug, ss) {
			indexRaw[itemSlug] = item
		}
	} else {
		indexRaw[itemSlug] = item
	}
	bytesIndex, err := json.Marshal(indexRaw)
	if err != nil {
		panic(err)
	}
	index := map[string]interface{}{
		"slug":  slug,
		"index": string(bytesIndex),
	}
	if new {
		s.Post("indices", index)
	} else {
		s.Put("indices", index)
	}
	fmt.Println("Indexed: ", itemSlug)
	return
}

func (s StrapiRestClient) GetIndex(col string) (ix map[string]interface{}) {
	indexRaw := []interface{}{}
	err := s.Get("indices", col, &indexRaw)
	utl.ErrorLog(err)
	index := make(map[string]interface{})
	if len(indexRaw) != 0 {
		index = indexRaw[0].(map[string]interface{})
		if index["slugs"] != "" {
			err = json.Unmarshal([]byte(index["index"].(string)), &ix)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

func CheckIndex(slug string, index map[string]interface{}) (c bool) {
	if _, found := index[slug]; !found {
		c = true
		return
	}
	return
}
