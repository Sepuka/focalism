package context

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type PrinterBuilder = func(tag language.Tag) *message.Printer
