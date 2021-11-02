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
	return nil
}

func (s StrapiRestClient) GetAll(col string, data interface{}) error {
	var count int
	call(http.MethodGet, s.BaseUrl+"/"+col+"/count", "application/json", nil, &count)
	times := count / 300
	start := 0
	var allDataRaw []interface{}
	for i := 0; i < times; i++ {
		var dataRaw interface{}
		call(http.MethodGet, s.BaseUrl+"/"+col+"?_start="+fmt.Sprint(start)+"&_limit=300", "application/json", nil, &dataRaw)
		start = start + 300
		fmt.Println("Times: ", i)
		allDataRaw = append(allDataRaw, dataRaw)
	}
	return nil
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
	postRest, _ := json.Marshal(data)
	responseBody := bytes.NewBuffer(postRest)
	resp, err := http.Post(s.BaseUrl+"/"+col, "application/json", responseBody)
	if err != nil {
		log.Print(err)
	}
	log.Print(resp)
	return nil
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
