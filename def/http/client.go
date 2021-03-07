package http

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/internal/config"
	"net/http"
)

const ClientDef = `http.client`

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ClientDef,
			Build: func(ctx di.Container) (interface{}, error) {
				return &http.Client{}, nil
			},
		})
	})
}
