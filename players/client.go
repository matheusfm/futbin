package players

import "net/http"

type Players interface {
	Get(opt PlayerOptions) ([]Player, error)
}

type players struct {
	client *http.Client
}

func (p players) Get(opt PlayerOptions) ([]Player, error) {
	panic("implement me")
}
