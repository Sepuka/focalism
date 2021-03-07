package message

import (
	"encoding/json"
	"fmt"
	"github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/focalism/internal/message/handler"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	"net/http"
)

type (
	MessageNew struct {
		cfg           *config.Config
		logger        *zap.SugaredLogger
		handlers      map[string]handler.MessageHandler
		answerHandler *handler.Answer
	}
)

func NewMessageNew(
	cfg *config.Config,
	handlers map[string]handler.MessageHandler,
	logger *zap.SugaredLogger,
	answer *handler.Answer,
) *MessageNew {
	return &MessageNew{
		cfg:           cfg,
		handlers:      handlers,
		logger:        logger,
		answerHandler: answer,
	}
}

func (o *MessageNew) Exec(req *domain.Request, resp http.ResponseWriter) error {
	var (
		err            error
		rawPayload     = req.Object.Message.Payload
		payloadData    button.Payload
		buttonHandler  handler.MessageHandler
		isHandlerKnown bool
	)

	defer func() {
		_, err = resp.Write(api.DefaultResponseBody())
	}()

	if req.IsKeyboardButton() {
		if err = json.Unmarshal([]byte(rawPayload), &payloadData); err != nil {
			return o.buildPayloadError(req.Object.Message, err, `invalid payload JSON`)
		}

		if buttonHandler, isHandlerKnown = o.handlers[payloadData.Command]; isHandlerKnown {
			o.logger.Debugf(`Got "%s" command`, payloadData.Command)
			if err = buttonHandler.Handle(req, &payloadData); err != nil {
				o.logger.Error(err)
			}
		}
	} else {
		if err = o.answerHandler.Handle(req); err != nil {
			return err
		}
	}

	return err
}

func (o *MessageNew) String() string {
	return `message_new`
}

func (o *MessageNew) buildPayloadError(msg domain.Message, err error, text string) error {
	var (
		rawPayload = msg.Payload
		userId     = msg.FromId
	)

	o.
		logger.
		With(
			zap.String(`json`, rawPayload),
			zap.Int32(`user_id`, userId),
			zap.Error(err),
		).
		Error(text)

	return errors.NewInvalidJsonError(fmt.Sprintf(`%v`, msg), err)
}
