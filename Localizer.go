package goi18n

import (
	"reflect"
)

type GenLocalizer interface {
	MessageKeyTypeName() string
	GenMessage(*Locale, any)
	NumberFormats() NumberFormats
}

type Localizer[KeyT any] interface {
	GenLocalizer
	Message(*Locale, KeyT)
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
