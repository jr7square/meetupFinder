package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {

	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
}
