package handler

import (
	"github.com/sepuka/focalism/def/lang"
	"github.com/sepuka/focalism/internal/context"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	MessageHandler interface {
		Handle(*domain.Request, *button.Payload) error
	}
	Builder = func(tag language.Tag) *message.Printer
)

func Printer(req *domain.Request) (*message.Printer, error) {
	var (
		err            error
		reqContext     = context.GetContext(req)
		printerBuilder Builder
		printer        *message.Printer
	)

	if err = reqContext.Container.Fill(lang.PrinterBuilderDef, &printerBuilder); err != nil {
		return nil, err
	}

	printer = printerBuilder(reqContext.Lang)

	return printer, err
}
