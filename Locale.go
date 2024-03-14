package goi18n

// https://en.wikipedia.org/wiki/IETF_language_tag

type Locale struct {
	PrimaryLanguage PrimaryLanguage
	ExtendedLanguages []ExtendedLanguage
	Script ScriptNumeric
	Region Region
	Variants []Variant
}
