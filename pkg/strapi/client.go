package strapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/comhttp/jorm/pkg/utl"
)

type StrapiRestClient struct {
	BaseUrl string
}

func New(url string) StrapiRestClient {
	return NewWithUlr(url)
}

func NewWithUlr(baseUrl string) (src StrapiRestClient) {
	if baseUrl == "" {
		panic("STRAPI BASE URL IS MANDATORY")
	}
	src.BaseUrl = baseUrl
	return
}

func call(method, url, contentType string, post []byte, response interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(post))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	fmt.Println("cccccc", response)
	return nil
}

func (s StrapiRestClient) GetAll(col string, data interface{}) error {
	return call(http.MethodGet, s.BaseUrl+"/"+col+"?_limit=99999", "application/json", nil, &data)
}

func (s StrapiRestClient) Get(col, slug string, data interface{}) error {
	return call(http.MethodGet, s.BaseUrl+"/"+col+"?slug="+slug, "application/json", nil, &data)
}

func (s StrapiRestClient) Put(col string, data interface{}) error {
	var res interface{}
	putRest, _ := json.Marshal(&data)
	log.Println("Reponse: ", &res)
	return call(http.MethodPut, "application/json", s.BaseUrl+"/"+col, putRest, &res)
}

func (s StrapiRestClient) Post(col string, data interface{}) error {
	var res interface{}

	postRest, _ := json.Marshal(data)
	log.Println("Reponse: ", &res)

	fmt.Println("s.BaseUrl", s.BaseUrl)

	return call(http.MethodPost, "application/json", s.BaseUrl+"/"+col, postRest, &res)
}

func (s StrapiRestClient) DelAll(col string) error {
	var all []map[string]interface{}
	err := call(http.MethodGet, s.BaseUrl+"/"+col+"?_limit=9999", "application/json", nil, &all)
	utl.ErrorLog(err)
	for _, entry := range all {
		go s.Del(col, fmt.Sprint(entry["id"]))
	}
	return nil
}

func (s StrapiRestClient) Del(col, id string) error {
	var res interface{}
	log.Println("Reponse: ", &res)
	return call(http.MethodDelete, s.BaseUrl+"/"+col+"/"+id, "application/json", nil, &res)
}
