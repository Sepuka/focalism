package message

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/def/api/method"
	"github.com/sepuka/focalism/def/log"
	"github.com/sepuka/focalism/def/repository"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/focalism/internal/domain"
	message2 "github.com/sepuka/focalism/internal/message"
	"github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/focalism/internal/message/handler"
	api2 "github.com/sepuka/vkbotserver/api"
	"go.uber.org/zap"
)

const (
	StartDef = `def.message.start`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Name: StartDef,
			Tags: []di.Tag{
				{
					Name: ExecutorDef,
					Args: nil,
				},
			},
			Build: func(ctx di.Container) (interface{}, error) {
				var (
					api            = ctx.Get(method.ApiDef).(*api2.Api)
					logger         = ctx.Get(log.LoggerDef).(*zap.SugaredLogger)
					vocabularyRepo = ctx.Get(repository.VocabularyRepoDef).(domain.VocabularyRepository)
					taskRepo       = ctx.Get(repository.TaskRepoDef).(domain.TaskRepository)
					topicRepo      = ctx.Get(repository.TopicRepoDef).(domain.TopicRepository)
					handlers       = map[string]handler.MessageHandler{
						button.ButtonIdStart:     handler.NewStartHandler(api),
						button.ButtonIdNext:      handler.NewNextHandler(api, vocabularyRepo, taskRepo),
						button.ButtonIdSurrender: handler.NewSurrenderHandler(api, taskRepo),
						button.ButtonIdReturn:    handler.NewReturnHandler(api),
						button.ButtonIdTopics:    handler.NewTopicHandler(api, topicRepo),
					}
					answerHandler = handler.NewAnswer(taskRepo, api, logger)
				)

				return message2.NewMessageNew(cfg, handlers, logger, answerHandler), nil
			},
		})
	})
}
