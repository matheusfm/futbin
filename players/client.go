package players

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Client interface {
	Get(opt *Options) ([]Player, error)
	TOTW() ([]Player, error)
	Popular() ([]Player, error)
	Latest() ([]Player, error)
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

func (c client) Get(opt *Options) ([]Player, error) {
	u, _ := url.Parse("https://www.futbin.org/futbin/api/getFilteredPlayers")
	q, err := query.Values(opt)
	if err != nil {
		return nil, err
	}
	u.RawQuery = q.Encode()
	return c.get(u.String())
}

func (c client) TOTW() ([]Player, error) {
	return c.get("https://www.futbin.org/futbin/api/currentTOTW")
}

func (c client) Popular() ([]Player, error) {
	return c.get("https://www.futbin.org/futbin/api/getPopularPlayers")
}

func (c client) Latest() ([]Player, error) {
	return c.get("https://www.futbin.org/futbin/api/newPlayers")
}

func (c client) get(url string) ([]Player, error) {
	r, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	var data playersData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data.Players, err
}
