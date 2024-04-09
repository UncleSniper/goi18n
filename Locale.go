package goi18n

// https://en.wikipedia.org/wiki/IETF_language_tag

type Locale struct {
	PrimaryLanguage PrimaryLanguage
	ExtendedLanguages []ExtendedLanguage
	Script ScriptNumeric
	Region Region
	Variants []Variant
	Extensions []Extension
}

func(loc *Locale) SortExtensions() error {
	//TODO
	return nil
}

func(loc *Locale) Equals(other *Locale) bool {
	if other == nil {
		other = fallbackLocale
	}
	if loc.PrimaryLanguage != other.PrimaryLanguage || loc.Script != other.Script {
		return false
	}
	if loc.Region != other.Region {
		return false
	}
	return compareSlices(loc.ExtendedLanguages, other.ExtendedLanguages) &&
			compareSlices(loc.Variants, other.Variants) && compareSlices(loc.Extensions, other.Extensions)
}

var fallbackLocale *Locale = &Locale {
	PrimaryLanguage: PL_English,
	ExtendedLanguages: nil,
	Script: SCNM_NONE, //TODO: use Latin
	Region: RG_NONE, //TODO: use GB
	Variants: nil,
	Extensions: nil,
}

func FallbackLocale() *Locale {
	return fallbackLocale
}

type LocaleMapper func(*Locale) *Locale
