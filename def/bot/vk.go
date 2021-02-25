package bot

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	messageDef "github.com/sepuka/focalism/def/message"
	"github.com/sepuka/focalism/def/middleware"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/vkbotserver/message"
	middleware2 "github.com/sepuka/vkbotserver/middleware"
	"github.com/sepuka/vkbotserver/server"
)

const (
	Bot = `def.bot.vk`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					handlerMap = container.Get(messageDef.HandlerMapDef).(message.HandlerMap)
					handler    = container.Get(middleware.BotMiddlewareDef).(middleware2.HandlerFunc)
				)

				return server.NewSocketServer(cfg.Server, handlerMap, handler), nil
			},
			Close:    nil,
			Name:     Bot,
			Scope:    "",
			Tags:     nil,
			Unshared: false,
		})
	})
}
