package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	var responseParsed Response
	url := "https://api.meetup.com/status?key=1c79452f27445bf6c365104a2b6132&sign=true"
	//signedURL := "https://api.meetup.com/2/cities?&sign=true&photo-host=public&query=tampa&page=20"
	netclient := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := netclient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	json.NewDecoder(resp.Body).Decode(&responseParsed)

	fmt.Println(responseParsed)
	fmt.Println(resp.Body)
}
