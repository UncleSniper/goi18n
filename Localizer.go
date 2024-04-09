package goi18n

import (
	"fmt"
	"reflect"
)

type GenLocalizer interface {
	MessageKeyTypeName() string
	GenMessage(*Locale, any) (string, bool)
	NumberFormats(*Locale) NumberFormats
}

type Localizer[KeyT any] interface {
	GenLocalizer
	Message(*Locale, KeyT) (string, bool)
}

type NumberFormats interface {
	DecimalSeparator() rune
	DigitGroupSeparator() rune
}

type NumberFormatsTemplate struct {
	DecimalSeparator rune
	DigitGroupSeparator rune
}

type numberFormatsImpl struct {
	decimalSeparator rune
	digitGroupSeparator rune
}

func(impl *numberFormatsImpl) DecimalSeparator() rune {
	return impl.decimalSeparator
}

func(impl *numberFormatsImpl) DigitGroupSeparator() rune {
	return impl.digitGroupSeparator
}

var fallbackNumberFormats *numberFormatsImpl = &numberFormatsImpl {
	decimalSeparator: '.',
	digitGroupSeparator: ',',
}

func ProtectNumberFormats(tpl *NumberFormatsTemplate) NumberFormats {
	if tpl == nil {
		return fallbackNumberFormats
	}
	return &numberFormatsImpl {
		decimalSeparator: tpl.DecimalSeparator,
		digitGroupSeparator: tpl.DigitGroupSeparator,
	}
}

func FallbackNumberFormats() NumberFormats {
	return fallbackNumberFormats
}

func GetMessageKeyTypeName[KeyT any]() string {
	return reflect.TypeOf(new(KeyT)).Elem().String()
}

func MessageBySpecific[KeyT any](
	localizer Localizer[KeyT],
	locale *Locale,
	key any,
	specific func(*Locale, KeyT) (string, bool),
) (string, bool) {
	typed, accept := key.(KeyT)
	if !accept {
		var gotType string
		if key == nil {
			gotType = "<nil>"
		} else {
			gotType = reflect.TypeOf(key).String()
		}
		panic(fmt.Sprintf(
			"Bad I18N message key: Expected type %s for localizer, but got %s",
			localizer.MessageKeyTypeName(),
			gotType,
		))
	}
	return specific(locale, typed)
}
