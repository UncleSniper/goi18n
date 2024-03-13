package goi18n

import (
	"fmt"
	"sync"
	"errors"
	"github.com/UncleSniper/goutil"
)

// https://en.wikipedia.org/wiki/ISO_15924
// https://unicode.org/iso15924/iso15924-codes.html

type ScriptNumeric uint32

const (
	SCNM_NONE = 0
	SCNM_PRIV_LOWEST ScriptNumeric = 900
	SCNM_PRIV_HIGHEST ScriptNumeric = 949
	SCNM_EMOJI ScriptNumeric = 993
	SCNM_INHERITED ScriptNumeric = 994
	SCNM_MATHEMATICAL ScriptNumeric = 995
	SCNM_SYMBOLS ScriptNumeric = 996
	SCNM_UNWRITTEN ScriptNumeric = 997
	SCNM_UNDETERMINED ScriptNumeric = 998
	SCNM_UNCODED ScriptNumeric = 999
	SCNM_ROOT ScriptNumeric = 1000000
	SCNM_TRUE ScriptNumeric = 1000001 // nobody seems to know what this is for?
)

type ScriptBlock uint

const (
	SCBL_NONE ScriptBlock = iota
	SCBL_Hieroglyphic_and_cuneiform
	SCBL_Right_to_left_alphabetic
	SCBL_Left_to_right_alphabetic
	SCBL_Alphasyllabic
	SCBL_Syllabic
	SCBL_Ideographic
	SCBL_Undeciphered
	SCBL_Shorthands
	SCBL_unassigned
	SCBL_Private_use_and_alias_and_special
)

const scbl_lowest_invalid ScriptNumeric = ScriptNumeric(SCBL_Private_use_and_alias_and_special * 100)

func(script ScriptNumeric) Block() ScriptBlock {
	if script == SCNM_NONE || script >= scbl_lowest_invalid {
		return SCBL_NONE
	} else {
		return ScriptBlock(script / 100 + 1)
	}
}

type ScriptSpecial uint

const (
	SCSP_NONE ScriptSpecial = iota
	SCSP_PRIVATE_USE
	SCSP_EMOJI
	SCSP_INHERITED
	SCSP_MATHEMATICAL
	SCSP_SYMBOLS
	SCSP_UNWRITTEN
	SCSP_UNDETERMINED
	SCSP_UNCODED
)

func(script ScriptNumeric) Special() ScriptSpecial {
	switch script {
		case SCNM_EMOJI:
			return SCSP_EMOJI
		case SCNM_INHERITED:
			return SCSP_INHERITED
		case SCNM_MATHEMATICAL:
			return SCSP_MATHEMATICAL
		case SCNM_SYMBOLS:
			return SCSP_SYMBOLS
		case SCNM_UNWRITTEN:
			return SCSP_UNWRITTEN
		case SCNM_UNDETERMINED:
			return SCSP_UNDETERMINED
		case SCNM_UNCODED:
			return SCSP_UNCODED
		default:
			if script >= SCNM_PRIV_LOWEST && script <= SCNM_PRIV_HIGHEST {
				return SCSP_PRIVATE_USE
			} else {
				return SCSP_NONE
			}
	}
}

const (
	SCCD_PRIV_LOWEST string = "Qaaa"
	SCCD_PRIV_HIGHEST string = "Qabx"
	SCCD_EMOJI string = "Zsye"
	SCCD_INHERITED string = "Zinh"
	SCCD_MATHEMATICAL string = "Zmth"
	SCCD_SYMBOLS string = "Zsym"
	SCCD_UNWRITTEN string = "Zxxx"
	SCCD_UNDETERMINED string = "Zyyy"
	SCCD_UNCODED string = "Zzzz"
	SCCD_ROOT string = "Root"
	SCCD_TRUE string = "True"
)

type ScriptInfo struct {
	numeric ScriptNumeric
	code string
	formalNames []string
	aliases []string
	directionality Directionality
	superset ScriptNumeric
}

func(info *ScriptInfo) Numeric() ScriptNumeric {
	if info == nil {
		return SCNM_NONE
	} else {
		return info.numeric
	}
}

func(info *ScriptInfo) Code() string {
	if info == nil {
		return ""
	} else {
		return info.code
	}
}

func(info *ScriptInfo) FormalNames() []string {
	if info == nil || len(info.formalNames) == 0 {
		return nil
	}
	return append([]string(nil), info.formalNames...)
}

func(info *ScriptInfo) Aliases() []string {
	if info == nil || len(info.aliases) == 0 {
		return nil
	}
	return append([]string(nil), info.aliases...)
}

func(info *ScriptInfo) Directionality() Directionality {
	if info == nil {
		return DIR_UNKNOWN
	} else {
		return info.directionality
	}
}

func(info *ScriptInfo) Superset() ScriptNumeric {
	if info == nil {
		return SCNM_NONE
	} else {
		return info.superset
	}
}

var scriptNumericMap map[ScriptNumeric]*ScriptInfo
var scriptCodeMap map[string]*ScriptInfo

var scriptMutex sync.Mutex

type ScriptFormalNamesBuilder interface {
	Names(formalNames ...string) ScriptAliasesBuilder
}

type ScriptAliasesBuilder interface {
	Aliases(aliases ...string) (*ScriptInfo, error)
}

type scriptBuilder struct {
	info *ScriptInfo
}

func NewScript(
	numeric ScriptNumeric,
	code string,
	directionality Directionality,
	superset ScriptNumeric,
) ScriptFormalNamesBuilder {
	info := &ScriptInfo {
		numeric: numeric,
		code: code,
		directionality: directionality,
		superset: superset,
	}
	return scriptBuilder{info}
}

func(builder scriptBuilder) Names(formalNames ...string) ScriptAliasesBuilder {
	info := builder.info
	for _, name := range formalNames {
		if len(name) > 0 && !info.hasFormalName(name) {
			info.formalNames = append(info.formalNames, name)
		}
	}
	return builder
}

func(builder scriptBuilder) Aliases(aliases ...string) (*ScriptInfo, error) {
	info := builder.info
	for _, alias := range aliases {
		if len(alias) > 0 && !info.hasAlias(alias) {
			info.aliases = append(info.aliases, alias)
		}
	}
	return registerScript(info)
}

func dieRegisterScript(numeric ScriptNumeric, msg string) error {
	if numeric < SCNM_PRIV_LOWEST || numeric > SCNM_PRIV_HIGHEST {
		panic(msg)
	}
	scriptMutex.Unlock()
	return errors.New(msg)
}

func registerScript(info *ScriptInfo) (*ScriptInfo, error) {
	if info.numeric == SCNM_NONE {
		panic("Cannot register script with ScriptNumeric SCNM_NONE")
	}
	scriptMutex.Lock()
	if scriptNumericMap == nil {
		scriptNumericMap = make(map[ScriptNumeric]*ScriptInfo)
	}
	if scriptCodeMap == nil {
		scriptCodeMap = make(map[string]*ScriptInfo)
	}
	byNumeric := scriptNumericMap[info.numeric]
	var byCode *ScriptInfo
	if len(info.code) > 0 {
		byCode = scriptCodeMap[info.code]
	}
	// merge numeric
	var setByNumericCode string
	if byNumeric != nil {
		if len(byNumeric.code) == 0 {
			if len(info.code) > 0 {
				if byCode != nil && byCode.numeric != info.numeric {
					msg := fmt.Sprintf(
						"Cannot register script %d with code '%s', since %d already has that code",
						uint32(info.numeric),
						info.code,
						uint32(byCode.numeric),
					)
					return nil, dieRegisterScript(info.numeric, msg)
				}
				setByNumericCode = info.code
			}
		} else if len(info.code) > 0 && info.code != byNumeric.code {
			msg := fmt.Sprintf(
				"Conflicting codes for previously registered script %d: '%s' vs. '%s'",
				uint32(info.numeric),
				info.code,
				byNumeric.code,
			)
			return nil, dieRegisterScript(info.numeric, msg)
		}
	}
	// merge code
	if byCode != nil {
		if byCode.numeric != info.numeric {
			msg := fmt.Sprintf(
				"Conflicting IDs for different scripts by code '%s': %d vs. %d",
				info.code,
				uint32(info.numeric),
				uint32(byCode.numeric),
			)
			return nil, dieRegisterScript(info.numeric, msg)
		}
		if byNumeric == nil {
			// this cannot happen
			panic(fmt.Sprintf(
				"In goi18n.registerScript(): byCode != nil && byCode.numeric == info.numeric && byNumeric == nil" +
						": info.numeric = %d, info.code = '%s', byNumeric.numeric = %d, byCode.numeric = %d",
				uint32(info.numeric),
				info.code,
				uint32(byNumeric.numeric),
				uint32(byCode.numeric),
			))
		}
	}
	// merge directionality
	var setByNumericDir Directionality
	if byNumeric != nil && info.directionality != DIR_UNKNOWN {
		if byNumeric.directionality == DIR_UNKNOWN {
			setByNumericDir = info.directionality
		} else if info.directionality != byNumeric.directionality {
			msg := fmt.Sprintf(
				"Conflicting Directionality values for script %d: %d vs. %d",
				uint32(info.numeric),
				uint(info.directionality),
				uint(byNumeric.directionality),
			)
			return nil, dieRegisterScript(info.numeric, msg)
		}
	}
	// merge superset
	var setByNumericSuperset ScriptNumeric
	if byNumeric != nil && info.superset != SCNM_NONE {
		if byNumeric.superset == SCNM_NONE {
			setByNumericSuperset = info.superset
		} else if info.superset != byNumeric.superset {
			msg := fmt.Sprintf(
				"Conflicting supersets for script %d: %d vs. %d",
				uint32(info.numeric),
				uint32(info.superset),
				uint32(byNumeric.superset),
			)
			return nil, dieRegisterScript(info.numeric, msg)
		}
	}
	// flush
	var newInfo *ScriptInfo
	if byNumeric == nil {
		newInfo = &ScriptInfo {
			numeric: info.numeric,
			code: info.code,
			formalNames: append([]string(nil), info.formalNames...),
			aliases: append([]string(nil), info.aliases...),
			directionality: info.directionality,
			superset: info.superset,
		}
		scriptNumericMap[info.numeric] = newInfo
		if len(info.code) > 0 {
			scriptCodeMap[info.code] = newInfo
		}
	} else {
		newInfo = byNumeric
		if len(setByNumericCode) > 0 {
			byNumeric.code = setByNumericCode
		}
		if setByNumericDir != DIR_UNKNOWN {
			byNumeric.directionality = setByNumericDir
		}
		if setByNumericSuperset != SCNM_NONE {
			byNumeric.superset = setByNumericSuperset
		}
		if byCode == nil && len(info.code) > 0 {
			scriptCodeMap[info.code] = byNumeric
		}
	}
	scriptMutex.Unlock()
	return newInfo, nil
}

func(info *ScriptInfo) hasFormalName(formalName string) bool {
	for _, name := range info.formalNames {
		if name == formalName {
			return true
		}
	}
	return false
}

func(info *ScriptInfo) AddFormalNames(formalNames ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&ScriptInfo{}, "AddFormalNames")
	}
	for _, name := range formalNames {
		if len(name) > 0 && !info.hasFormalName(name) {
			info.formalNames = append(info.formalNames, name)
		}
	}
	return nil
}

func(info *ScriptInfo) hasAlias(alias string) bool {
	for _, name := range info.aliases {
		if name == alias {
			return true
		}
	}
	return false
}

func(info *ScriptInfo) AddAliases(aliases ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&ScriptInfo{}, "AddAliases")
	}
	for _, alias := range aliases {
		if len(alias) > 0 && !info.hasAlias(alias) {
			info.aliases = append(info.aliases, alias)
		}
	}
	return nil
}

func GetScriptByNumeric(numeric ScriptNumeric) (info *ScriptInfo) {
	if numeric == SCNM_NONE {
		return
	}
	scriptMutex.Lock()
	if scriptNumericMap != nil {
		info = scriptNumericMap[numeric]
	}
	scriptMutex.Unlock()
	return
}

func GetScriptByCode(code string) (info *ScriptInfo) {
	if len(code) == 0 {
		return
	}
	scriptMutex.Lock()
	if scriptCodeMap != nil {
		info = scriptCodeMap[code]
	}
	scriptMutex.Unlock()
	return
}
