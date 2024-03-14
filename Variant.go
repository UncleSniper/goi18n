package goi18n

import (
	"fmt"
	"sync"
	"errors"
	"github.com/UncleSniper/goutil"
)

type Variant uint32

const (
	VAR_NONE Variant = 0
)

type VariantInfo struct {
	id Variant
	language PrimaryLanguage
	prefix string
	shortNames []string
	longNames []string
}

func(info *VariantInfo) ID() Variant {
	if info == nil {
		return VAR_NONE
	}
	return info.id
}

func(info *VariantInfo) Language() PrimaryLanguage {
	if info == nil {
		return NO_PRIMARY_LANGUAGE
	}
	return info.language
}

func(info *VariantInfo) Prefix() string {
	if info == nil {
		return ""
	}
	return info.prefix
}

func(info *VariantInfo) ShortNames() []string {
	if info == nil || len(info.shortNames) == 0 {
		return nil
	}
	return append([]string(nil), info.shortNames...)
}

func(info *VariantInfo) ShortNamesWithPrefix() []string {
	if info == nil || len(info.shortNames) == 0 {
		return nil
	}
	if len(info.prefix) == 0 {
		return append([]string(nil), info.shortNames...)
	}
	names := make([]string, 0, len(info.shortNames))
	for _, name := range info.shortNames {
		names = append(names, fmt.Sprintf("%s-%s", info.prefix, name))
	}
	return names
}

func(info *VariantInfo) LongNames() []string {
	if info == nil || len(info.longNames) == 0 {
		return nil
	}
	return append([]string(nil), info.longNames...)
}

var variants []*VariantInfo
var variantShortNameMap map[PrimaryLanguage]map[string]*VariantInfo
var variantPrefixMap map[string]*VariantInfo

var variantMutex sync.Mutex

type VariantLongNamesBuilder interface {
	Names(longNames ...string) Variant
}

type variantBuilder struct {
	info *VariantInfo
}

func NewVariant(language PrimaryLanguage, variantPrefix string, shortNames ...string) VariantLongNamesBuilder {
	info := &VariantInfo {
		language: language,
		prefix: variantPrefix,
	}
	for _, name := range shortNames {
		if len(name) > 0 && !info.hasShortName(name) {
			info.shortNames = append(info.shortNames, name)
		}
	}
	return variantBuilder{info}
}

func(builder variantBuilder) Names(longNames ...string) Variant {
	info := builder.info
	for _, name := range longNames {
		if len(name) > 0 && !info.hasLongName(name) {
			info.longNames = append(info.longNames, name)
		}
	}
	return registerVariant(info)
}

func registerVariant(info *VariantInfo) (variant Variant) {
	if info.language == NO_PRIMARY_LANGUAGE {
		panic("Cannot register new Variant: No associated language set")
	}
	variantMutex.Lock()
	if variantShortNameMap == nil {
		variantShortNameMap = make(map[PrimaryLanguage]map[string]*VariantInfo)
		variantPrefixMap = make(map[string]*VariantInfo)
	}
	langInfo := GetPrimaryLanguageInfo(info.language)
	if langInfo == nil {
		panic(fmt.Sprintf(
			"Cannot register new Variant: PrimaryLanguage %d is not defined",
			uint32(info.language),
		))
	}
	ofLanguage := variantShortNameMap[info.language]
	if ofLanguage == nil {
		ofLanguage = make(map[string]*VariantInfo)
		variantShortNameMap[info.language] = ofLanguage
	}
	var fullPrefix string
	if len(langInfo.variantPrefix) > 0 {
		fullPrefix = langInfo.variantPrefix + "-"
	}
	for _, name := range info.shortNames {
		prev := ofLanguage[name]
		if prev != nil {
			panic(fmt.Sprintf(
				"Cannot register new Variant with short name '%s' for PrimaryLanguage %d: " +
					"Variant %d already has this short name",
				name,
				uint32(info.language),
				uint32(prev.id),
			))
		}
		fullName := fullPrefix + name
		prev = variantPrefixMap[fullName]
		if prev != nil {
			panic(fmt.Sprintf(
				"Cannot register new Variant with full short name '%s': Variant %d already has this short name",
				fullName,
				uint32(prev.id),
			))
		}
	}
	variant = Variant(len(variants) + 1)
	if variant == 0 {
		panic("Too many Variant instances registered")
	}
	newInfo := &VariantInfo {
		id: variant,
		language: info.language,
		prefix: langInfo.variantPrefix,
		shortNames: append([]string(nil), info.shortNames...),
		longNames: append([]string(nil), info.longNames...),
	}
	for _, name := range info.shortNames {
		ofLanguage[name] = newInfo
		variantPrefixMap[fullPrefix + name] = newInfo
	}
	variantMutex.Unlock()
	return
}

func(info *VariantInfo) hasShortName(shortName string) bool {
	for _, name := range info.shortNames {
		if name == shortName {
			return true
		}
	}
	return false
}

func(info *VariantInfo) AddShortNames(shortNames ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&VariantInfo{}, "AddShortNames")
	}
	if info.id == VAR_NONE {
		return errors.New("Trying to call VariantInfo.AddShortNames() when ID() == VAR_NONE")
	}
	variantMutex.Lock()
	if variantShortNameMap == nil {
		variantShortNameMap = make(map[PrimaryLanguage]map[string]*VariantInfo)
		variantPrefixMap = make(map[string]*VariantInfo)
	}
	langInfo := GetPrimaryLanguageInfo(info.language)
	if langInfo == nil {
		variantMutex.Unlock()
		return errors.New("Trying to call VariantInfo.AddShortNames() when Language() == NO_PRIMARY_LANGUAGE")
	}
	ofLanguage := variantShortNameMap[info.language]
	if ofLanguage == nil {
		ofLanguage = make(map[string]*VariantInfo)
		variantShortNameMap[info.language] = ofLanguage
	}
	var fullPrefix string
	if len(langInfo.variantPrefix) > 0 {
		fullPrefix = langInfo.variantPrefix + "-"
	}
	for _, name := range shortNames {
		if len(name) == 0 || info.hasShortName(name) {
			continue
		}
		prev := ofLanguage[name]
		if prev != nil && prev.id != info.id {
			variantMutex.Unlock()
			return errors.New(fmt.Sprintf(
				"Cannot add short name '%s' for PrimaryLanguage %d to Variant %d, " +
						"since %d already has this short name for that PrimaryLanguage",
				name,
				uint32(info.id),
				uint32(prev.id),
			))
		}
		fullName := fullPrefix + name
		prev = variantPrefixMap[fullName]
		if prev != nil && prev.id != info.id {
			variantMutex.Unlock()
			return errors.New(fmt.Sprintf(
				"Cannot add full short name '%s' to Variant %d, since %d already has this full short name",
				fullName,
				uint32(info.id),
				uint32(prev.id),
			))
		}
	}
	for _, name := range shortNames {
		ofLanguage[name] = info
		variantPrefixMap[fullPrefix + name] = info
	}
	variantMutex.Unlock()
	return nil
}

func(info *VariantInfo) hasLongName(longName string) bool {
	for _, name := range info.longNames {
		if name == longName {
			return true
		}
	}
	return false
}

func(info *VariantInfo) AddLongNames(longNames ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&VariantInfo{}, "AddLongNames")
	}
	if info.id == VAR_NONE {
		return errors.New("Trying to call VariantInfo.AddLongNames() when ID() == VAR_NONE")
	}
	for _, name := range longNames {
		if len(name) > 0 && !info.hasLongName(name) {
			info.longNames = append(info.longNames, name)
		}
	}
	return nil
}

func GetVariantInfo(variant Variant) (info *VariantInfo) {
	if variant == VAR_NONE {
		return
	}
	variantMutex.Lock()
	if variant <= Variant(len(variants)) {
		info = variants[variant - 1]
	}
	variantMutex.Unlock()
	return
}

func GetVariantInfoByLanguageAndShortName(language PrimaryLanguage, shortName string) (info *VariantInfo) {
	if language == NO_PRIMARY_LANGUAGE || len(shortName) == 0 {
		return
	}
	variantMutex.Lock()
	if variantShortNameMap != nil {
		ofLanguage := variantShortNameMap[language]
		if ofLanguage != nil {
			info = ofLanguage[shortName]
		}
	}
	variantMutex.Unlock()
	return
}

func GetVariantInfoByFullShortName(shortName string) (info *VariantInfo) {
	if len(shortName) == 0 {
		return
	}
	variantMutex.Lock()
	if variantPrefixMap != nil {
		info = variantPrefixMap[shortName]
	}
	variantMutex.Unlock()
	return
}

func initVariants() {
	// from https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry with 'Type: variant
	//TODO
}
