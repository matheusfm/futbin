package leagueclubs

import (
	"encoding/json"
	"net/http"
)

type Client interface {
	Get() ([]League, error)
}

func DefaultClient() Client {
	return NewClient(http.DefaultClient)
}

func NewClient(c *http.Client) Client {
	return &client{client: c}
}

type client struct {
	client *http.Client
}

func (c client) Get() ([]League, error) {
	r, err := c.client.Get("https://www.futbin.org/futbin/api/getLeaguesAndClubsAndroid")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data leaguesAndClubs
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Leagues, nil
}
