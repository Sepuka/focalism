package message

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/vkbotserver/message"
)

const (
	ConfirmationDef = `def.message.confirmation`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: ConfirmationDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx di.Container) (interface{}, error) {
				return message.NewConfirmation(cfg.Server), nil
			},
		})
	})
}
