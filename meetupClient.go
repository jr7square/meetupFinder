package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

/*MeetupClient is a go client for the MeetUp Api */
type MeetupClient struct {
	client *http.Client
	key    string
}

/*NewMeetupClient constructs MeetupClinet*/
func NewMeetupClient(apiKey string) *MeetupClient {

	mc := &MeetupClient{
		client: &http.Client{Timeout: time.Second * 10},
		key:    apiKey,
	}
	return mc
}

/*Status executes a status request */
func (mc *MeetupClient) Status() ([]byte, error) {
	url := "https://api.meetup.com/status?key" + mc.key + "&sign=true"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := mc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
