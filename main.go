package main

import (
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	// var responseParsed Response
	// url := "https://api.meetup.com/status?key=1c79452f27445bf6c365104a2b6132&sign=true"
	// netclient := &http.Client{
	// 	Timeout: time.Second * 10,
	// }
	// req, _ := http.NewRequest("GET", url, nil)
	// resp, err := netclient.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// //body, err := ioutil.ReadAll(resp.Body)
	// json.NewDecoder(resp.Body).Decode(&responseParsed)
	//
	// fmt.Printf("%+v\n", responseParsed)
	h := NewHandler()
	http.ListenAndServe(":8080", h)
}
