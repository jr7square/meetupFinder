package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	router *httprouter.Router
}

func NewHandler() *Handler {
	h := &Handler{
		router: httprouter.New(),
	}

	h.router.GET("/status", h.meetUpStatus)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) meetUpStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//url := "https://api.meetup.com/status?key=1c79452f27445bf6c365104a2b6132&sign=true"
	url := "https://api.meetup.com/2/cities?&sign=true&photo-host=public&query=tampa&page=20"
	netclient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, _ := http.NewRequest("GET", url, nil)
	resp, err := netclient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	w.Write(data)

}
