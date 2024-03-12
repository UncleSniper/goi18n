package goi18n

type ExtendedLanguage interface {
	PrimaryLanguage() PrimaryLanguage
}

type PrimaryExtendedLanguage struct {
	Language PrimaryLanguage
}

func(ext PrimaryExtendedLanguage) PrimaryLanguage() PrimaryLanguage {
	return ext.Language
}
