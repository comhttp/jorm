package utl

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

// FileExists returns whether a file exists
func GetSource(url string, srcRaw interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		log.Print("Error: ", err)
	}
	defer res.Body.Close()
	mapBody, err := ioutil.ReadAll(res.Body)
	if mapBody != nil {
		json.Unmarshal(mapBody, &srcRaw)
	}
	return err
}

// FileExists returns whether a file exists
func GetSourceHeadersAPIkey(apiKey, url string, srcRaw interface{}) {
	apiKeyHeader := "Apikey " + apiKey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Get Source Headers API key, creating HTTP request, %s", err)
	}
	req.Header.Set("authorization", apiKeyHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Get Source Headers API key, querying url, %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Get Source Headers API key, reading body %s", err)
	}
	if err := json.Unmarshal(body, &srcRaw); err != nil {
		log.Printf("Get Source Headers API key, unmarshalling response, %s", err)
	}
}
