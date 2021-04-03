package middleware

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/internal/context"
	"github.com/sepuka/vkbotserver/domain"
	"github.com/sepuka/vkbotserver/message"
	"github.com/sepuka/vkbotserver/middleware"
	"net/http"
)

func NewContext(ctx di.Container) func(handlerFunc middleware.HandlerFunc) middleware.HandlerFunc {
	return func(handlerFunc middleware.HandlerFunc) middleware.HandlerFunc {
		return func(exec message.Executor, req *domain.Request, writer http.ResponseWriter) error {
			var (
				reqContext = &context.Context{}
				err        error
			)

			if reqContext.Container, err = ctx.SubContainer(); err != nil {
				return err
			}

			defer reqContext.Container.Delete()

			return handlerFunc(exec, context.WithContext(req, reqContext), writer)
		}
	}
}
