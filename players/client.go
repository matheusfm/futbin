package players

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Players interface {
	Get(opt PlayerOptions) ([]Player, error)
}

type players struct {
	client *http.Client
}

func (p players) Get(opt PlayerOptions) ([]Player, error) {
	u, _ := url.Parse("https://www.futbin.org/futbin/api/getFilteredPlayers")
	q, err := query.Values(opt)
	if err != nil {
		return nil, err
	}
	u.RawQuery = q.Encode()
	r, err := http.Get(u.String())
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
