package httpreq

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type RestyConfig struct {
	Client *http.Client
}

type Resty struct {
	resty.Client
}

func NewResty(cfg RestyConfig) *Resty {
	client := &Resty{}
	if cfg.Client == nil {
		client.Client = *resty.New()
	} else {
		client.Client = *resty.NewWithClient(cfg.Client)
	}
	return client
}
