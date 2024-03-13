package goi18n

import (
	"fmt"
	"sync"
	"errors"
	"github.com/UncleSniper/goutil"
)

type PrimaryLanguage uint32

const (
	NO_PRIMARY_LANGUAGE PrimaryLanguage = 0
	PL_Abkhazian PrimaryLanguage = 1
	PL_Abkhaz PrimaryLanguage = PL_Abkhazian
	PL_Afar PrimaryLanguage = 2
	PL_Afrikaans PrimaryLanguage = 3
	PL_Akan PrimaryLanguage = 4
	PL_Albanian PrimaryLanguage = 5
	PL_Amharic PrimaryLanguage = 6
	PL_Arabic PrimaryLanguage = 7
	PL_Aragonese PrimaryLanguage = 8
	PL_Armenian PrimaryLanguage = 9
	PL_Assamese PrimaryLanguage = 10
	PL_Avaric PrimaryLanguage = 11
	PL_Avar PrimaryLanguage = PL_Avaric
	PL_Avestan PrimaryLanguage = 12
	PL_Aymara PrimaryLanguage = 13
	PL_Azerbaijani PrimaryLanguage = 14
	PL_Azeri PrimaryLanguage = PL_Azerbaijani
	PL_Bambara PrimaryLanguage = 15
	PL_Bashkir PrimaryLanguage = 16
	PL_Basque PrimaryLanguage = 17
	PL_Belarusian PrimaryLanguage = 18
	PL_Bengali PrimaryLanguage = 19
	PL_Bangla PrimaryLanguage = PL_Bengali
	PL_Bislama PrimaryLanguage = 20
	PL_Bosnian PrimaryLanguage = 21
	PL_Breton PrimaryLanguage = 22
	PL_Bulgarian PrimaryLanguage = 23
	PL_Burmese PrimaryLanguage = 24
	PL_Myanmar PrimaryLanguage = PL_Burmese
	PL_Catalan PrimaryLanguage = 25
	PL_Valencian PrimaryLanguage = PL_Catalan
	PL_Chamorro PrimaryLanguage = 26
	PL_Chechen PrimaryLanguage = 27
	PL_Chichewa PrimaryLanguage = 28
	PL_Chewa PrimaryLanguage = PL_Chichewa
	PL_Nyanja PrimaryLanguage = PL_Chichewa
	PL_Chinese PrimaryLanguage = 29
	PL_Church_Slavonic PrimaryLanguage = 30
	PL_Old_Slavonic PrimaryLanguage = PL_Church_Slavonic
	PL_Old_Church_Slavonic PrimaryLanguage = PL_Church_Slavonic
	PL_Chuvash PrimaryLanguage = 31
	PL_Cornish PrimaryLanguage = 32
	PL_Corsican PrimaryLanguage = 33
	PL_Cree PrimaryLanguage = 34
	PL_Croatian PrimaryLanguage = 35
	PL_Czech PrimaryLanguage = 36
	PL_Danish PrimaryLanguage = 37
	PL_Divehi PrimaryLanguage = 38
	PL_Dhivehi PrimaryLanguage = PL_Divehi
	PL_Maldivian PrimaryLanguage = PL_Divehi
	PL_Dutch PrimaryLanguage = 39
	PL_Flemish PrimaryLanguage = PL_Dutch
	PL_Dzongkha PrimaryLanguage = 40
	PL_English PrimaryLanguage = 41
	PL_Esperanto PrimaryLanguage = 42
	PL_Estonian PrimaryLanguage = 43
	PL_Ewe PrimaryLanguage = 44
	PL_Faroese PrimaryLanguage = 45
	PL_Fijian PrimaryLanguage = 46
	PL_Finnish PrimaryLanguage = 47
	PL_French PrimaryLanguage = 48
	PL_Western_Frisian PrimaryLanguage = 49
	PL_Frisian PrimaryLanguage = PL_Western_Frisian
	PL_Fries PrimaryLanguage = PL_Western_Frisian
	PL_Fulah PrimaryLanguage = 50
	PL_Fula PrimaryLanguage = PL_Fulah
	PL_Gaelic PrimaryLanguage = 51
	PL_Scottish_Gaelic PrimaryLanguage = PL_Gaelic
	PL_Galician PrimaryLanguage = 52
	PL_Ganda PrimaryLanguage = 53
	PL_Georgian PrimaryLanguage = 54
	PL_German PrimaryLanguage = 55
	PL_Greek PrimaryLanguage = 56
	PL_Kalaallisut PrimaryLanguage = 57
	PL_Greenlandic PrimaryLanguage = PL_Kalaallisut
	PL_Guarani PrimaryLanguage = 58
	PL_Gujarati PrimaryLanguage = 59
	PL_Haitian PrimaryLanguage = 60
	PL_Haitian_Creole PrimaryLanguage = PL_Haitian
	PL_Hausa PrimaryLanguage = 61
	PL_Hebrew PrimaryLanguage = 62
	PL_Herero PrimaryLanguage = 63
	PL_Hindi PrimaryLanguage = 64
	PL_Hiri_Motu PrimaryLanguage = 65
	PL_Hungarian PrimaryLanguage = 66
	PL_Icelandic PrimaryLanguage = 67
	PL_Ido PrimaryLanguage = 68
	PL_Igbo PrimaryLanguage = 69
	PL_Indonesian PrimaryLanguage = 70
	PL_Interlingua PrimaryLanguage = 71
	PL_Interlingue PrimaryLanguage = 72
	PL_Occidental PrimaryLanguage = PL_Interlingue
	PL_Inuktitut PrimaryLanguage = 73
	PL_Inupiaq PrimaryLanguage = 74
	PL_Irish PrimaryLanguage = 75
	PL_Italian PrimaryLanguage = 76
	PL_Japanese PrimaryLanguage = 77
	PL_Javanese PrimaryLanguage = 78
	PL_Kannada PrimaryLanguage = 79
	PL_Kanuri PrimaryLanguage = 80
	PL_Kashmiri PrimaryLanguage = 81
	PL_Kazakh PrimaryLanguage = 82
	PL_Central_Khmer PrimaryLanguage = 83
	PL_Khmer PrimaryLanguage = PL_Central_Khmer
	PL_Cambodian PrimaryLanguage = PL_Central_Khmer
	PL_Kikuyu PrimaryLanguage = 84
	PL_Gikuyu PrimaryLanguage = PL_Kikuyu
	PL_Kinyarwanda PrimaryLanguage = 85
	PL_Kirghiz PrimaryLanguage = 86
	PL_Kyrgyz PrimaryLanguage = PL_Kirghiz
	PL_Komi PrimaryLanguage = 87
	PL_Kongo PrimaryLanguage = 88
	PL_Korean PrimaryLanguage = 89
	PL_Kuanyama PrimaryLanguage = 90
	PL_Kwanyama PrimaryLanguage = PL_Kuanyama
	PL_Kurdish PrimaryLanguage = 91
	PL_Lao PrimaryLanguage = 92
	PL_Latin PrimaryLanguage = 93
	PL_Latvian PrimaryLanguage = 94
	PL_Limburgan PrimaryLanguage = 95
	PL_Limburger PrimaryLanguage = PL_Limburgan
	PL_Limburgish PrimaryLanguage = PL_Limburgan
	PL_Lingala PrimaryLanguage = 96
	PL_Lithuanian PrimaryLanguage = 97
	PL_Luba_Katanga PrimaryLanguage = 98
	PL_Luba_Shaba PrimaryLanguage = PL_Luba_Katanga
	PL_Luxembourgish PrimaryLanguage = 99
	PL_Letzeburgesch PrimaryLanguage = PL_Luxembourgish
	PL_Macedonian PrimaryLanguage = 100
	PL_Malagasy PrimaryLanguage = 101
	PL_Malay PrimaryLanguage = 102
	PL_Malayalam PrimaryLanguage = 103
	PL_Maltese PrimaryLanguage = 104
	PL_Manx PrimaryLanguage = 105
	PL_Maori PrimaryLanguage = 106
	PL_Marathi PrimaryLanguage = 107
	PL_Marshallese PrimaryLanguage = 108
	PL_Mongolian PrimaryLanguage = 109
	PL_Nauru PrimaryLanguage = 110
	PL_Nauruan PrimaryLanguage = PL_Nauru
	PL_Navajo PrimaryLanguage = 111
	PL_Navaho PrimaryLanguage = PL_Navajo
	PL_North_Ndebele PrimaryLanguage = 112
	PL_Northern_Ndebele PrimaryLanguage = PL_North_Ndebele
	PL_South_Ndebele PrimaryLanguage = 113
	PL_Southern_Ndebele PrimaryLanguage = PL_South_Ndebele
	PL_Ndonga PrimaryLanguage = 114
	PL_Nepali PrimaryLanguage = 115
	PL_Norwegian PrimaryLanguage = 116
	PL_Norwegian_Bokmal PrimaryLanguage = 117
	PL_Norwegian_Bokmål PrimaryLanguage = PL_Norwegian_Bokmal
	PL_Norwegian_Nynorsk PrimaryLanguage = 118
	PL_Sichuan_Yi PrimaryLanguage = 119
	PL_Nuosu PrimaryLanguage = PL_Sichuan_Yi
	PL_Occitan PrimaryLanguage = 120
	PL_Ojibwa PrimaryLanguage = 121
	PL_Ojibwe PrimaryLanguage = PL_Ojibwa
	PL_Oriya PrimaryLanguage = 122
	PL_Odia PrimaryLanguage = PL_Oriya
	PL_Oromo PrimaryLanguage = 123
	PL_Ossetian PrimaryLanguage = 124
	PL_Ossetic PrimaryLanguage = PL_Ossetian
	PL_Pali PrimaryLanguage = 125
	PL_Pāli PrimaryLanguage = PL_Pali
	PL_Pashto PrimaryLanguage = 126
	PL_Pushto PrimaryLanguage = PL_Pashto
	PL_Persian PrimaryLanguage = 127
	PL_Farsi PrimaryLanguage = PL_Persian
	PL_Polish PrimaryLanguage = 128
	PL_Portuguese PrimaryLanguage = 129
	PL_Punjabi PrimaryLanguage = 130
	PL_Panjabi PrimaryLanguage = PL_Punjabi
	PL_Romanian PrimaryLanguage = 131
	PL_Moldavian PrimaryLanguage = PL_Romanian
	PL_Moldovan PrimaryLanguage = PL_Romanian
	PL_Romansh PrimaryLanguage = 132
	PL_Rundi PrimaryLanguage = 133
	PL_Kirundi PrimaryLanguage = PL_Rundi
	PL_Russian PrimaryLanguage = 134
	PL_Northern_Sami PrimaryLanguage = 135
	PL_Samoan PrimaryLanguage = 136
	PL_Sango PrimaryLanguage = 137
	PL_Sanskrit PrimaryLanguage = 138
	PL_Sardinian PrimaryLanguage = 139
	PL_Serbian PrimaryLanguage = 140
	PL_Shona PrimaryLanguage = 141
	PL_Sindhi PrimaryLanguage = 142
	PL_Sinhala PrimaryLanguage = 143
	PL_Sinhalese PrimaryLanguage = PL_Sinhala
	PL_Slovak PrimaryLanguage = 144
	PL_Slovenian PrimaryLanguage = 145
	PL_Slovene PrimaryLanguage = PL_Slovenian
	PL_Somali PrimaryLanguage = 146
	PL_Southern_Sotho PrimaryLanguage = 147
	PL_Spanish PrimaryLanguage = 148
	PL_Castilian PrimaryLanguage = PL_Spanish
	PL_Sundanese PrimaryLanguage = 149
	PL_Swahili PrimaryLanguage = 150
	PL_Swati PrimaryLanguage = 151
	PL_Swazi PrimaryLanguage = PL_Swati
	PL_Swedish PrimaryLanguage = 152
	PL_Tagalog PrimaryLanguage = 153
	PL_Tahitian PrimaryLanguage = 154
	PL_Tajik PrimaryLanguage = 155
	PL_Tamil PrimaryLanguage = 156
	PL_Tamizh PrimaryLanguage = PL_Tamil
	PL_Tatar PrimaryLanguage = 157
	PL_Telugu PrimaryLanguage = 158
	PL_Thai PrimaryLanguage = 159
	PL_Tibetan PrimaryLanguage = 160
	PL_Tigrinya PrimaryLanguage = 161
	PL_Tonga PrimaryLanguage = 162
	PL_Tongan PrimaryLanguage = PL_Tonga
	PL_Tsonga PrimaryLanguage = 163
	PL_Tswana PrimaryLanguage = 164
	PL_Turkish PrimaryLanguage = 165
	PL_Turkmen PrimaryLanguage = 166
	PL_Twi PrimaryLanguage = 167
	PL_Uighur PrimaryLanguage = 168
	PL_Uyghur PrimaryLanguage = PL_Uighur
	PL_Ukrainian PrimaryLanguage = 169
	PL_Urdu PrimaryLanguage = 170
	PL_Uzbek PrimaryLanguage = 171
	PL_Venda PrimaryLanguage = 172
	PL_Vietnamese PrimaryLanguage = 173
	PL_Volapük PrimaryLanguage = 174
	PL_Walloon PrimaryLanguage = 175
	PL_Welsh PrimaryLanguage = 176
	PL_Wolof PrimaryLanguage = 177
	PL_Xhosa PrimaryLanguage = 178
	PL_Yiddish PrimaryLanguage = 179
	PL_Yoruba PrimaryLanguage = 180
	PL_Zhuang PrimaryLanguage = 181
	PL_Chuang PrimaryLanguage = PL_Zhuang
	PL_Zulu PrimaryLanguage = 182
)

type PrimaryLanguageInfo struct {
	id PrimaryLanguage
	parent PrimaryLanguage
	shortNames []string
	longNames []string
}

func(info *PrimaryLanguageInfo) ID() PrimaryLanguage {
	if info == nil {
		return NO_PRIMARY_LANGUAGE
	}
	return info.id
}

func(info *PrimaryLanguageInfo) Parent() PrimaryLanguage {
	if info == nil {
		return NO_PRIMARY_LANGUAGE
	}
	return info.parent
}

func(info *PrimaryLanguageInfo) ShortNames() []string {
	if info == nil || len(info.shortNames) == 0 {
		return nil
	}
	return append([]string(nil), info.shortNames...)
}

func(info *PrimaryLanguageInfo) LongNames() []string {
	if info == nil || len(info.longNames) == 0 {
		return nil
	}
	return append([]string(nil), info.longNames...)
}

var primaryLanguages []*PrimaryLanguageInfo
var primaryLanguageMap map[string]PrimaryLanguage

var primaryLanguageMutex sync.Mutex

type PrimaryLanguageLongNamesBuilder interface {
	Names(longNames ...string) PrimaryLanguage
}

type primaryLanguageBuilder struct {
	info *PrimaryLanguageInfo
}

func NewPrimaryLanguage(parent PrimaryLanguage, shortNames ...string) PrimaryLanguageLongNamesBuilder {
	info := &PrimaryLanguageInfo {
		parent: parent,
	}
	for _, name := range shortNames {
		if len(name) > 0 && !info.hasShortName(name) {
			info.shortNames = append(info.shortNames, name)
		}
	}
	return primaryLanguageBuilder{info}
}

func(builder primaryLanguageBuilder) Names(longNames ...string) PrimaryLanguage {
	info := builder.info
	for _, name := range longNames {
		if len(name) > 0 && !info.hasLongName(name) {
			info.longNames = append(info.longNames, name)
		}
	}
	return registerPrimaryLanguage(info)
}

func registerPrimaryLanguage(info *PrimaryLanguageInfo) (lang PrimaryLanguage) {
	primaryLanguageMutex.Lock()
	if primaryLanguageMap == nil {
		primaryLanguageMap = make(map[string]PrimaryLanguage)
	}
	for _, name := range info.shortNames {
		if primaryLanguageMap[name] != NO_PRIMARY_LANGUAGE {
			panic(fmt.Sprintf("Cannot register new PrimaryLanguage: Short name '%s' is already registered", name))
		}
	}
	lang = PrimaryLanguage(len(primaryLanguages) + 1)
	if lang == 0 {
		panic("Too many PrimaryLanguage instances registered")
	}
	primaryLanguages = append(primaryLanguages, &PrimaryLanguageInfo {
		id: lang,
		parent: info.parent,
		shortNames: append([]string(nil), info.shortNames...),
		longNames: append([]string(nil), info.longNames...),
	})
	for _, name := range info.shortNames {
		primaryLanguageMap[name] = lang
	}
	primaryLanguageMutex.Unlock()
	return
}

func(info *PrimaryLanguageInfo) hasShortName(shortName string) bool {
	for _, name := range info.shortNames {
		if name == shortName {
			return true
		}
	}
	return false
}

func(info *PrimaryLanguageInfo) AddShortNames(shortNames ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&PrimaryLanguageInfo{}, "AddShortNames")
	}
	if info.id == NO_PRIMARY_LANGUAGE {
		return errors.New("Trying to call PrimaryLanguageInfo.AddShortNames() when ID() == NO_PRIMARY_LANGUAGE")
	}
	primaryLanguageMutex.Lock()
	if primaryLanguageMap == nil {
		primaryLanguageMap = make(map[string]PrimaryLanguage)
	}
	for _, name := range shortNames {
		if len(name) > 0 && !info.hasShortName(name) {
			prev, _ := primaryLanguageMap[name]
			if prev != NO_PRIMARY_LANGUAGE {
				if prev == info.id {
					continue
				}
				primaryLanguageMutex.Unlock()
				return errors.New(fmt.Sprintf(
					"Cannot add short name '%s' to PrimaryLanguage %d, since %d already has this short name",
					name,
					uint32(info.id),
					uint32(prev),
				))
			}
		}
	}
	for _, name := range shortNames {
		if len(name) > 0 {
			primaryLanguageMap[name] = info.id
		}
	}
	primaryLanguageMutex.Unlock()
	return nil
}

func(info *PrimaryLanguageInfo) hasLongName(longName string) bool {
	for _, name := range info.longNames {
		if name == longName {
			return true
		}
	}
	return false
}

func(info *PrimaryLanguageInfo) AddLongNames(longNames ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&PrimaryLanguageInfo{}, "AddLongNames")
	}
	for _, name := range info.longNames {
		if len(name) > 0 && !info.hasLongName(name) {
			info.longNames = append(info.longNames, name)
		}
	}
	return nil
}

func GetPrimaryLanguageInfo(lang PrimaryLanguage) (info *PrimaryLanguageInfo) {
	if lang == NO_PRIMARY_LANGUAGE {
		return
	}
	primaryLanguageMutex.Lock()
	if lang <= PrimaryLanguage(len(primaryLanguages)) {
		info = primaryLanguages[lang - 1]
	}
	primaryLanguageMutex.Unlock()
	return
}

// https://en.wikipedia.org/wiki/ISO_639-1
// https://en.wikipedia.org/wiki/ISO_639-2

// https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes

func verifyPrimaryLanguage(registered PrimaryLanguage, expected ...PrimaryLanguage) {
	for _, exp := range expected {
		if registered != exp {
			panic(fmt.Sprintf(
				"Something is wrong in initPrimaryLanguages(): Registered %d, expected it to be equal to %d",
				uint32(registered),
				uint32(exp),
			))
		}
	}
}

func initPrimaryLanguages() {
	// from https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ab",
		"abk",
	).Names("Abkhazian", "Abkhaz"), PL_Abkhazian, PL_Abkhazian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"aa",
		"aar",
	).Names("Afar"), PL_Afar)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"af",
		"afr",
	).Names("Afrikaans"), PL_Afrikaans)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ak",
		"aka",
	).Names("Akan"), PL_Akan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sq",
		"sqi",
		"alb",
	).Names("Albanian"), PL_Albanian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"am",
		"amh",
	).Names("Amharic"), PL_Amharic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ar",
		"ara",
	).Names("Arabic"), PL_Arabic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"an",
		"arg",
	).Names("Aragonese"), PL_Aragonese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"hy",
		"hye",
		"arm",
	).Names("Armenian"), PL_Armenian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"as",
		"asm",
	).Names("Assamese"), PL_Assamese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"av",
		"ava",
	).Names("Avaric", "Avar"), PL_Avaric, PL_Avar)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ae",
		"ave",
	).Names("Avestan"), PL_Avestan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ay",
		"aym",
	).Names("Aymara"), PL_Aymara)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"az",
		"aze",
	).Names("Azerbaijani", "Azeri"), PL_Azerbaijani, PL_Azeri)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bm",
		"bam",
	).Names("Bambara"), PL_Bambara)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ba",
		"bak",
	).Names("Bashkir"), PL_Bashkir)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"eu",
		"eus",
		"baq",
	).Names("Basque"), PL_Basque)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"be",
		"bel",
	).Names("Belarusian"), PL_Belarusian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bn",
		"ben",
	).Names("Bengali", "Bangla"), PL_Bengali, PL_Bangla)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bi",
		"bis",
	).Names("Bislama"), PL_Bislama)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bs",
		"bos",
	).Names("Bosnian"), PL_Bosnian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"br",
		"bre",
	).Names("Breton"), PL_Breton)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bg",
		"bul",
	).Names("Bulgarian"), PL_Bulgarian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"my",
		"mya",
		"bur",
	).Names("Burmese", "Myanmar"), PL_Burmese, PL_Myanmar)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ca",
		"cat",
	).Names("Catalan", "Valencian"), PL_Catalan, PL_Valencian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ch",
		"cha",
	).Names("Chamorro"), PL_Chamorro)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ce",
		"che",
	).Names("Chechen"), PL_Chechen)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ny",
		"nya",
	).Names("Chichewa", "Chewa", "Nyanja"), PL_Chichewa, PL_Chewa, PL_Nyanja)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"zh",
		"zho",
		"chi",
	).Names("Chinese"), PL_Chinese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"cu",
		"chu",
	).Names(
		"Church Slavonic",
		"Old Slavonic",
		"Old Church Slavonic",
	), PL_Church_Slavonic, PL_Old_Slavonic, PL_Old_Church_Slavonic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"cv",
		"chv",
	).Names("Chuvash"), PL_Chuvash)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kw",
		"cor",
	).Names("Cornish"), PL_Cornish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"co",
		"cos",
	).Names("Corsican"), PL_Corsican)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"cr",
		"cre",
	).Names("Cree"), PL_Cree)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"hr",
		"hrv",
	).Names("Croatian"), PL_Croatian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"cs",
		"ces",
		"cze",
	).Names("Czech"), PL_Czech)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"da",
		"dan",
	).Names("Danish"), PL_Danish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"dv",
		"div",
	).Names("Divehi", "Dhivehi", "Maldivian"), PL_Divehi, PL_Dhivehi, PL_Maldivian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"nl",
		"nld",
		"dut",
	).Names("Dutch", "Flemish"), PL_Dutch, PL_Flemish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"dz",
		"dzo",
	).Names("Dzongkha"), PL_Dzongkha)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"en",
		"eng",
	).Names("English"), PL_English)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"eo",
		"epo",
	).Names("Esperanto"), PL_Esperanto)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"et",
		"est",
	).Names("Estonian"), PL_Estonian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ee",
		"ewe",
	).Names("Ewe"), PL_Ewe)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fo",
		"fao",
	).Names("Faroese"), PL_Faroese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fj",
		"fij",
	).Names("Fijian"), PL_Fijian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fi",
		"fin",
	).Names("Finnish"), PL_Finnish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fr",
		"fra",
		"fre",
	).Names("French"), PL_French)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fy",
		"fry",
	).Names("Western Frisian", "Frisian", "Fries"), PL_Western_Frisian, PL_Frisian, PL_Fries)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ff",
		"ful",
	).Names("Fulah", "Fula"), PL_Fulah, PL_Fula)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"gd",
		"gla",
	).Names("Gaelic", "Scottish Gaelic"), PL_Gaelic, PL_Scottish_Gaelic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"gl",
		"glg",
	).Names("Galician"), PL_Galician)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lg",
		"lug",
	).Names("Ganda"), PL_Ganda)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ka",
		"kat",
		"geo",
	).Names("Georgian"), PL_Georgian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"de",
		"deu",
		"ger",
	).Names("German"), PL_German)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"el",
		"ell",
		"gre",
	).Names("Greek"), PL_Greek)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kl",
		"kal",
	).Names("Kalaallisut", "Greenlandic"), PL_Kalaallisut, PL_Greenlandic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"gn",
		"grn",
	).Names("Guarani"), PL_Guarani)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"gu",
		"guj",
	).Names("Gujarati"), PL_Gujarati)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ht",
		"hat",
	).Names("Haitian", "Haitian Creole"), PL_Haitian, PL_Haitian_Creole)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ha",
		"hau",
	).Names("Hausa"), PL_Hausa)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"he",
		"heb",
	).Names("Hebrew"), PL_Hebrew)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"hz",
		"her",
	).Names("Herero"), PL_Herero)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"hi",
		"hin",
	).Names("Hindi"), PL_Hindi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ho",
		"hmo",
	).Names("Hiri Motu"), PL_Hiri_Motu)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"hu",
		"hun",
	).Names("Hungarian"), PL_Hungarian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"is",
		"isl",
		"ice",
	).Names("Icelandic"), PL_Icelandic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"io",
		"ido",
	).Names("Ido"), PL_Ido)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ig",
		"ibo",
	).Names("Igbo"), PL_Igbo)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		PL_Malay,
		"id",
		"ind",
	).Names("Indonesian"), PL_Indonesian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ia",
		"ina",
	).Names("Interlingua"), PL_Interlingua)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ie",
		"ile",
	).Names("Interlingue", "Occidental"), PL_Interlingue, PL_Occidental)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"iu",
		"iku",
	).Names("Inuktitut"), PL_Inuktitut)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ik",
		"ipk",
	).Names("Inupiaq"), PL_Inupiaq)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ga",
		"gle",
	).Names("Irish"), PL_Irish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"it",
		"ita",
	).Names("Italian"), PL_Italian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ja",
		"jpn",
	).Names("Japanese"), PL_Japanese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"jv",
		"jav",
	).Names("Javanese"), PL_Javanese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kn",
		"kan",
	).Names("Kannada"), PL_Kannada)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kr",
		"kau",
	).Names("Kanuri"), PL_Kanuri)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ks",
		"kas",
	).Names("Kashmiri"), PL_Kashmiri)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kk",
		"kaz",
	).Names("Kazakh"), PL_Kazakh)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"km",
		"khm",
	).Names("Central Khmer", "Khmer", "Cambodian"), PL_Central_Khmer, PL_Khmer, PL_Cambodian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ki",
		"kik",
	).Names("Kikuyu", "Gikuyu"), PL_Kikuyu, PL_Gikuyu)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"rw",
		"kin",
	).Names("Kinyarwanda"), PL_Kinyarwanda)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ky",
		"kir",
	).Names("Kirghiz", "Kyrgyz"), PL_Kirghiz, PL_Kyrgyz)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kv",
		"kom",
	).Names("Komi"), PL_Komi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kg",
		"kon",
	).Names("Kongo"), PL_Kongo)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ko",
		"kor",
	).Names("Korean"), PL_Korean)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"kj",
		"kua",
	).Names("Kuanyama", "Kwanyama"), PL_Kuanyama, PL_Kwanyama)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ku",
		"kur",
	).Names("Kurdish"), PL_Kurdish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lo",
		"lao",
	).Names("Lao"), PL_Lao)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"la",
		"lat",
	).Names("Latin"), PL_Latin)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lv",
		"lav",
	).Names("Latvian"), PL_Latvian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"li",
		"lim",
	).Names("Limburgan", "Limburger", "Limburgish"), PL_Limburgan, PL_Limburger, PL_Limburgish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ln",
		"lin",
	).Names("Lingala"), PL_Lingala)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lt",
		"lit",
	).Names("Lithuanian"), PL_Lithuanian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lu",
		"lub",
	).Names("Luba-Katanga", "Luba-Shaba"), PL_Luba_Katanga)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"lb",
		"ltz",
	).Names("Luxembourgish", "Letzeburgesch"), PL_Luxembourgish, PL_Letzeburgesch)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mk",
		"mkd",
		"mac",
	).Names("Macedonian"), PL_Macedonian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mg",
		"mlg",
	).Names("Malagasy"), PL_Malagasy)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ms",
		"msa",
		"may",
	).Names("Malay"), PL_Malay)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ml",
		"mal",
	).Names("Malayalam"), PL_Malayalam)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mt",
		"mlt",
	).Names("Maltese"), PL_Maltese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"gv",
		"glv",
	).Names("Manx"), PL_Manx)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mi",
		"mri",
		"mao",
	).Names("Maori"), PL_Maori)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mr",
		"mar",
	).Names("Marathi"), PL_Marathi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mh",
		"mah",
	).Names("Marshallese"), PL_Marshallese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"mn",
		"mon",
	).Names("Mongolian"), PL_Mongolian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"na",
		"nau",
	).Names("Nauru", "Nauruan"), PL_Nauru, PL_Nauruan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"nv",
		"nav",
	).Names("Navajo", "Navaho"), PL_Navajo, PL_Navaho)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"nd",
		"nde",
	).Names("North Ndebele", "Northern Ndebele"), PL_North_Ndebele, PL_Northern_Ndebele)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"nr",
		"nbl",
	).Names("South Ndebele", "Southern Ndebele"), PL_South_Ndebele, PL_Southern_Ndebele)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ng",
		"ndo",
	).Names("Ndonga"), PL_Ndonga)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ne",
		"nep",
	).Names("Nepali"), PL_Nepali)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"no",
		"nor",
	).Names("Norwegian"), PL_Norwegian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		PL_Norwegian,
		"nb",
		"nob",
	).Names("Norwegian Bokmal", "Norwegian Bokmål"), PL_Norwegian_Bokmal, PL_Norwegian_Bokmål)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		PL_Norwegian,
		"nn",
		"nno",
	).Names("Norwegian Nynorsk"), PL_Norwegian_Nynorsk)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ii",
		"iii",
	).Names("Sichuan Yi", "Nuosu"), PL_Sichuan_Yi, PL_Nuosu)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"oc",
		"oci",
	).Names("Occitan"), PL_Occitan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"oj",
		"oji",
	).Names("Ojibwa", "Ojibwe"), PL_Ojibwa, PL_Ojibwe)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"or",
		"ori",
	).Names("Oriya", "Odia"), PL_Oriya, PL_Odia)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"om",
		"orm",
	).Names("Oromo"), PL_Oromo)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"os",
		"oss",
	).Names("Ossetian", "Ossetic"), PL_Ossetian, PL_Ossetic)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"pi",
		"pli",
	).Names("Pali", "Pāli"), PL_Pali, PL_Pāli)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ps",
		"pus",
	).Names("Pashto", "Pushto"), PL_Pashto, PL_Pushto)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"fa",
		"fas",
		"per",
	).Names("Persian", "Farsi"), PL_Persian, PL_Farsi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"pl",
		"pol",
	).Names("Polish"), PL_Polish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"pt",
		"por",
	).Names("Portuguese"), PL_Portuguese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"pa",
		"pan",
	).Names("Punjabi", "Panjabi"), PL_Punjabi, PL_Panjabi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ro",
		"ron",
		"rum",
		"mo",
		"mol",
	).Names("Romanian", "Moldavian", "Moldovan"), PL_Romanian, PL_Moldavian, PL_Moldovan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"rm",
		"roh",
	).Names("Romansh"), PL_Romansh)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"rn",
		"run",
	).Names("Rundi", "Kirundi"), PL_Rundi, PL_Kirundi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ru",
		"rus",
	).Names("Russian"), PL_Russian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"se",
		"sme",
	).Names("Northern Sami"), PL_Northern_Sami)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sm",
		"smo",
	).Names("Samoan"), PL_Samoan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sg",
		"sag",
	).Names("Sango"), PL_Sango)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sa",
		"san",
	).Names("Sanskrit"), PL_Sanskrit)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sc",
		"srd",
	).Names("Sardinian"), PL_Sardinian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sr",
		"srp",
	).Names("Serbian"), PL_Serbian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sn",
		"sna",
	).Names("Shona"), PL_Shona)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sd",
		"snd",
	).Names("Sindhi"), PL_Sindhi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"si",
		"sin",
	).Names("Sinhala", "Sinhalese"), PL_Sinhala, PL_Sinhalese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sk",
		"slk",
		"slo",
	).Names("Slovak"), PL_Slovak)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sl",
		"slv",
	).Names("Slovenian", "Slovene"), PL_Slovenian, PL_Slovene)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"so",
		"som",
	).Names("Somali"), PL_Somali)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"st",
		"sot",
	).Names("Southern Sotho"), PL_Southern_Sotho)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"es",
		"spa",
	).Names("Spanish", "Castilian"), PL_Spanish, PL_Castilian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"su",
		"sun",
	).Names("Sundanese"), PL_Sundanese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sw",
		"swa",
	).Names("Swahili"), PL_Swahili)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ss",
		"ssw",
	).Names("Swati", "Swazi"), PL_Swati, PL_Swazi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"sv",
		"swe",
	).Names("Swedish"), PL_Swedish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tl",
		"tgl",
	).Names("Tagalog"), PL_Tagalog)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ty",
		"tah",
	).Names("Tahitian"), PL_Tahitian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tg",
		"tgk",
	).Names("Tajik"), PL_Tajik)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ta",
		"tam",
	).Names("Tamil", "Tamizh"), PL_Tamil, PL_Tamizh)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tt",
		"tat",
	).Names("Tatar"), PL_Tatar)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"te",
		"tel",
	).Names("Telugu"), PL_Telugu)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"th",
		"tha",
	).Names("Thai"), PL_Thai)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"bo",
		"bod",
	).Names("Tibetan"), PL_Tibetan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ti",
		"tir",
		"tib",
	).Names("Tigrinya"), PL_Tigrinya)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"to",
		"ton",
	).Names("Tonga", "Tongan"), PL_Tonga, PL_Tongan)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ts",
		"tso",
	).Names("Tsonga"), PL_Tsonga)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tn",
		"tsn",
	).Names("Tswana"), PL_Tswana)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tr",
		"tur",
	).Names("Turkish"), PL_Turkish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"tk",
		"tuk",
	).Names("Turkmen"), PL_Turkmen)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		PL_Akan,
		"tw",
		"twi",
	).Names("Twi"), PL_Twi)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ug",
		"uig",
	).Names("Uighur", "Uyghur"), PL_Uighur, PL_Uyghur)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"uk",
		"ukr",
	).Names("Ukrainian"), PL_Ukrainian)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ur",
		"urd",
	).Names("Urdu"), PL_Urdu)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"uz",
		"uzb",
	).Names("Uzbek"), PL_Uzbek)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"ve",
		"ven",
	).Names("Venda"), PL_Venda)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"vi",
		"vie",
	).Names("Vietnamese"), PL_Vietnamese)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"vo",
		"vol",
	).Names("Volapük"), PL_Volapük)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"wa",
		"wln",
	).Names("Walloon"), PL_Walloon)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"cy",
		"cym",
		"wel",
	).Names("Welsh"), PL_Welsh)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"wo",
		"wol",
	).Names("Wolof"), PL_Wolof)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"xh",
		"xho",
	).Names("Xhosa"), PL_Xhosa)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"yi",
		"yid",
	).Names("Yiddish"), PL_Yiddish)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"yo",
		"yor",
	).Names("Yoruba"), PL_Yoruba)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"za",
		"zha",
	).Names("Zhuang", "Chuang"), PL_Zhuang, PL_Chuang)
	verifyPrimaryLanguage(NewPrimaryLanguage(
		NO_PRIMARY_LANGUAGE,
		"zu",
		"zul",
	).Names("Zulu"), PL_Zulu)
	// from https://en.wikipedia.org/wiki/List_of_ISO_639-2_codes
	// from https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes
}
