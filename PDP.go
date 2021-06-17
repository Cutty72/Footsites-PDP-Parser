package PDP

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	RequestCreationErr = errors.New("error creating request")
	RequestSendErr = errors.New("error sending request")
)

func Parse(site, sku, stockStatus string) (map[string]string, error){
	client := &http.Client{}


	url := fmt.Sprintf("https://www.%v/api/products/pdp/%s", site, sku)
	PDP, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, RequestCreationErr
	}

	PDP.Header.Add("content-type", "application/json")
	PDP.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	resp, err := client.Do(PDP)
	if err != nil {
		return nil, RequestSendErr
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ReadBodyErr
	}

	if resp.StatusCode != http.StatusOK {
		return nil, RequestSendErr
	}

	mapSize, err := GetSizes(body, sku, stockStatus)
	if err != nil {
		return nil, err
	}
	return mapSize, nil
}
