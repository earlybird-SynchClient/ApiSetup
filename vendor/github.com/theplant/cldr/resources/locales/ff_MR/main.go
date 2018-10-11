package ff_MR

import "github.com/theplant/cldr"

var Locale = &cldr.Locale{
	Locale: "ff_MR",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
