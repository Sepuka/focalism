package lang

import (
	"github.com/sarulabs/di"
	"github.com/sepuka/focalism/def"
	"github.com/sepuka/focalism/internal/config"
	"github.com/sepuka/focalism/internal/context"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	message2 "golang.org/x/text/message"
)

const (
	dictionaryDef     = `def.lang.dictionary`
	PrinterBuilderDef = `def.lang.printer.builder`
)

func init() {
	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				return nil, message2.Set(language.Russian, context.KeyLangTasksPerDay, plural.Selectf(1, `%d`,
					plural.One, `%d слово`,
					plural.Few, `%d слова`,
					plural.Many, `%d слов`,
					plural.Other, `%d слов`,
				),
				)
			},
			Close: nil,
			Name:  dictionaryDef,
			Scope: di.App,
		})
	})

	def.Register(func(builder *di.Builder, cfg *config.Config) error {
		return builder.Add(di.Def{
			Build: func(container di.Container) (interface{}, error) {
				var (
					err error
				)

				if _, err = container.SafeGet(dictionaryDef); err != nil {
					return nil, err
				}

				return func(tag language.Tag) *message2.Printer {
					return message2.NewPrinter(tag)
				}, nil
			},
			Scope: di.Request,
			Name:  PrinterBuilderDef,
		})
	})
}
