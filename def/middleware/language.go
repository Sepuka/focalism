package middleware

import (
	context2 "github.com/sepuka/focalism/internal/context"
	"github.com/sepuka/vkbotserver/domain"
	"github.com/sepuka/vkbotserver/message"
	"github.com/sepuka/vkbotserver/middleware"
	"golang.org/x/text/language"
	"net/http"
)

func Language(next middleware.HandlerFunc) middleware.HandlerFunc {
	return func(exec message.Executor, req *domain.Request, writer http.ResponseWriter) error {
		var (
			context *context2.Context
		)

		context = context2.GetContext(req)
		context.Lang = language.Russian

		return next(exec, req, writer)
	}
}
