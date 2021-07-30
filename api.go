package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	IP          string `json:"Query"`
	Status      string
	Country     string
	CountryCode string
	Region      string
	RegionName  string
	City        string
	Timezone    string
	Isp         string
	Org         string
	As          string
	Asname      string
	Proxy       bool
}

func get_from_api(ip string) Response {

	url := "http://ip-api.com/json/" + ip + "?fields=status,message,country,countryCode,region,regionName,city,timezone,isp,org,as,asname,proxy,query"

	spaceClient := http.Client{
		Timeout: time.Second * 20, // Timeout after 2 seconds
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)

	if reqErr != nil {
		return Response{IP: ip}
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return Response{IP: ip}
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return Response{IP: ip}
	}

	apiResponse := Response{}
	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		return Response{IP: ip}
	}

	return apiResponse

}
