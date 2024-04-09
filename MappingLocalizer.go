package goi18n

type MappingLocalizer[KeyT any] struct {
	MessagePaths []*MappingPath[KeyT]
	NumberFormatPaths []*MappingPath[KeyT]
}

type MappingPath[KeyT any] struct {
	LocaleMapper LocaleMapper
	Localizer Localizer[KeyT]
	SubPaths []*MappingPath[KeyT]
}

func(loc *MappingLocalizer[KeyT]) MessageKeyTypeName() string {
	return GetMessageKeyTypeName[KeyT]()
}

func(loc *MappingLocalizer[KeyT]) GenMessage(locale *Locale, key any) (string, bool) {
	return MessageBySpecific[KeyT](loc, locale, key, loc.Message)
}

func(loc *MappingLocalizer[KeyT]) NumberFormats(locale *Locale) (formats NumberFormats) {
	for _, path := range loc.NumberFormatPaths {
		formats = path.numberFormats(locale)
		if formats != nil {
			break
		}
	}
	return
}

func(path *MappingPath[KeyT]) mapLocale(locale *Locale) (mappedLocale *Locale) {
	if path == nil {
		return
	}
	if locale == nil {
		locale = fallbackLocale
	}
	if path.LocaleMapper != nil {
		mappedLocale = path.LocaleMapper(locale)
	} else {
		mappedLocale = locale
	}
	return
}

func(path *MappingPath[KeyT]) numberFormats(locale *Locale) (formats NumberFormats) {
	mappedLocale := path.mapLocale(locale)
	if mappedLocale == nil {
		return
	}
	if path.Localizer != nil {
		formats = path.Localizer.NumberFormats(mappedLocale)
		if formats != nil {
			return
		}
	}
	for _, subPath := range path.SubPaths {
		formats = subPath.numberFormats(mappedLocale)
		if formats != nil {
			break
		}
	}
	return
}

func(loc *MappingLocalizer[KeyT]) Message(locale *Locale, key KeyT) (msg string, ok bool) {
	for _, path := range loc.MessagePaths {
		msg, ok = path.message(locale, key)
		if ok {
			break
		}
	}
	return
}

func(path *MappingPath[KeyT]) message(locale *Locale, key KeyT) (msg string, ok bool) {
	mappedLocale := path.mapLocale(locale)
	if mappedLocale == nil {
		return
	}
	if path.Localizer != nil {
		msg, ok = path.Localizer.Message(locale, key)
		if ok {
			return
		}
	}
	for _, subPath := range path.SubPaths {
		msg, ok = subPath.message(mappedLocale, key)
		if ok {
			break
		}
	}
	return
}

// === builder ===

type MappingLocalizerBuilder[KeyT any] interface {
	Path(LocaleMapper) MappingLocalizerBuilder[KeyT]
	With(Localizer[KeyT]) MappingLocalizerBuilder[KeyT]
	PathWith(LocaleMapper, Localizer[KeyT]) MappingLocalizerBuilder[KeyT]
}

type mappingLocalizerBuilderImpl[KeyT any] struct {
	slice *[]*MappingPath[KeyT]
	localizer Localizer[KeyT]
}

func(builder *mappingLocalizerBuilderImpl[KeyT]) Path(mapper LocaleMapper) MappingLocalizerBuilder[KeyT] {
	newPath := &MappingPath[KeyT] {
		LocaleMapper: mapper,
		Localizer: builder.localizer,
	}
	*builder.slice = append(*builder.slice, newPath)
	return &mappingLocalizerBuilderImpl[KeyT] {
		slice: &newPath.SubPaths,
		localizer: builder.localizer,
	}
}

func(builder *mappingLocalizerBuilderImpl[KeyT]) With(localizer Localizer[KeyT]) MappingLocalizerBuilder[KeyT] {
	builder.localizer = localizer
	return builder
}

func(builder *mappingLocalizerBuilderImpl[KeyT]) PathWith(
	mapper LocaleMapper,
	localizer Localizer[KeyT],
) MappingLocalizerBuilder[KeyT] {
	return builder.Path(mapper).With(localizer)
}

func(loc *MappingLocalizer[KeyT]) MessagePathBuilder(localizer Localizer[KeyT]) MappingLocalizerBuilder[KeyT] {
	if loc == nil {
		return nil
	}
	return &mappingLocalizerBuilderImpl[KeyT] {
		slice: &loc.MessagePaths,
		localizer: localizer,
	}
}

func(loc *MappingLocalizer[KeyT]) NumberFormatPathBuilder(localizer Localizer[KeyT]) MappingLocalizerBuilder[KeyT] {
	if loc == nil {
		return nil
	}
	return &mappingLocalizerBuilderImpl[KeyT] {
		slice: &loc.NumberFormatPaths,
		localizer: localizer,
	}
}

// === mappers ===

type LocaleWithout uint

const (
	LOCWO_PrimaryLanguage LocaleWithout = 1 << iota
	LOCWO_ExtendedLanguages
	LOCWO_Script
	LOCWO_Region
	LOCWO_Variants
	LOCWO_Extensions
)

func Without(flags LocaleWithout) LocaleMapper {
	return func(locale *Locale) (result *Locale) {
		if locale == nil {
			return
		}
		result = &Locale {}
		if (flags & LOCWO_PrimaryLanguage) == 0 {
			result.PrimaryLanguage = locale.PrimaryLanguage
		}
		if (flags & LOCWO_ExtendedLanguages) == 0 {
			result.ExtendedLanguages = locale.ExtendedLanguages
		}
		if (flags & LOCWO_Script) == 0 {
			result.Script = locale.Script
		}
		if (flags & LOCWO_Region) == 0 {
			result.Region = locale.Region
		}
		if (flags & LOCWO_Variants) == 0 {
			result.Variants = locale.Variants
		}
		if (flags & LOCWO_Extensions) == 0 {
			result.Extensions = locale.Extensions
		}
		return
	}
}
