package goi18n

import (
	"fmt"
	"sync"
	"errors"
	"github.com/UncleSniper/goutil"
)

type Region uint32

type RegionNumeric uint32

type RegionInfo struct {
	id Region
	parent Region
	level uint32
	numeric []RegionNumeric
	codes []string
	regionType RegionType
	subregions []Region
}

type RegionType uint

const (
	RGTY_NONE RegionType = iota
	RGTY_LANDMASS
	RGTY_CONTINENT
	RGTY_SUBCONTINENT
	RGTY_SUBSUBCONTINENT
	RGTY_COUNTRY
)

const RG_NONE Region = 0

const RGNM_NONE RegionNumeric = 0

func(info *RegionInfo) ID() Region {
	if info == nil {
		return RG_NONE
	} else {
		return info.id
	}
}

func(info *RegionInfo) Parent() Region {
	if info == nil {
		return RG_NONE
	} else {
		return info.parent
	}
}

func(info *RegionInfo) Level() uint32 {
	if info == nil {
		return 0
	} else {
		return info.level
	}
}

func(info *RegionInfo) Numeric() []RegionNumeric {
	if info == nil {
		return nil
	} else {
		return append([]RegionNumeric(nil), info.numeric...)
	}
}

func(info *RegionInfo) Codes() []string {
	if info == nil {
		return nil
	} else {
		return append([]string(nil), info.codes...)
	}
}

func(info *RegionInfo) Type() RegionType {
	if info == nil {
		return RGTY_NONE
	} else {
		return info.regionType
	}
}

func(info *RegionInfo) Subregions() []Region {
	if info == nil {
		return nil
	} else {
		return append([]Region(nil), info.subregions...)
	}
}

var regions []*RegionInfo
var regionNumericMap map[RegionNumeric]*RegionInfo
var regionCodeMap map[string]*RegionInfo

var regionMutex sync.Mutex

type RegionCodesBuilder interface {
	Codes(codes ...string) Region
}

type regionBuilder struct {
	info *RegionInfo
}

func NewRegion(parent Region, regionType RegionType, numeric ...RegionNumeric) RegionCodesBuilder {
	info := &RegionInfo {
		parent: parent,
		regionType: regionType,
	}
	for _, n := range numeric {
		if n != RGNM_NONE && !info.hasNumeric(n) {
			info.numeric = append(info.numeric, n)
		}
	}
	return regionBuilder{info}
}

func(builder regionBuilder) Codes(codes ...string) Region {
	info := builder.info
	for _, code := range codes {
		if len(code) > 0 && !info.hasCode(code) {
			info.codes = append(info.codes, code)
		}
	}
	return registerRegion(info)
}

func registerRegion(info *RegionInfo) (region Region) {
	if info.regionType == RGTY_NONE {
		panic("Cannot register new Region: RegionType is RGTY_NONE")
	}
	regionMutex.Lock()
	region = Region(len(regions) + 1)
	if info.parent >= region {
		panic(fmt.Sprintf("Cannot register new Region: Undefined parent: %d", uint32(info.parent)))
	}
	var newLevel uint32
	var parentInfo *RegionInfo
	if info.parent != RG_NONE {
		parentInfo = regions[info.parent - 1]
		newLevel = parentInfo.level + 1
		if newLevel == 0 {
			panic("Cannot register new Region: Level is too deep")
		}
	}
	if regionNumericMap == nil {
		regionNumericMap = make(map[RegionNumeric]*RegionInfo)
		regionCodeMap = make(map[string]*RegionInfo)
	}
	for _, n := range info.numeric {
		if regionNumericMap[n] != nil {
			panic(fmt.Sprintf("Cannot register new Region: Numeric code '%d' is already registered", uint32(n)))
		}
	}
	for _, code := range info.codes {
		if regionCodeMap[code] != nil {
			panic(fmt.Sprintf("Cannot register new Region: String code '%s' is already registered", code))
		}
	}
	newInfo := &RegionInfo {
		id: region,
		parent: info.parent,
		level: newLevel,
		numeric: append([]RegionNumeric(nil), info.numeric...),
		codes: append([]string(nil), info.codes...),
		regionType: info.regionType,
	}
	regions = append(regions, newInfo)
	for _, n := range info.numeric {
		regionNumericMap[n] = newInfo
	}
	for _, code := range info.codes {
		regionCodeMap[code] = newInfo
	}
	if parentInfo != nil {
		parentInfo.subregions = append(parentInfo.subregions, region)
	}
	regionMutex.Unlock()
	return
}

func(info *RegionInfo) hasNumeric(numeric RegionNumeric) bool {
	for _, n := range info.numeric {
		if n == numeric {
			return true
		}
	}
	return false
}

func(info *RegionInfo) AddNumeric(numeric ...RegionNumeric) error {
	if info == nil {
		return goutil.NewNilTargetError(&RegionInfo{}, "AddNumeric")
	}
	if info.id == RG_NONE {
		return errors.New("Trying to call RegionInfo.AddNumeric() when ID() == RG_NONE")
	}
	regionMutex.Lock()
	if regionNumericMap == nil {
		regionNumericMap = make(map[RegionNumeric]*RegionInfo)
		regionCodeMap = make(map[string]*RegionInfo)
	}
	for _, n := range numeric {
		if n == RGNM_NONE {
			continue
		}
		prev := regionNumericMap[n]
		if prev != nil && prev.id != info.id {
			regionMutex.Unlock()
			return errors.New(fmt.Sprintf(
				"Cannot add numeric code %d to region %d: Region %d already has this code",
				uint32(n),
				uint32(info.id),
				uint32(prev.id),
			))
		}
	}
	for _, n := range numeric {
		if n != RGNM_NONE {
			regionNumericMap[n] = info
		}
	}
	regionMutex.Unlock()
	return nil
}

func(info *RegionInfo) hasCode(code string) bool {
	for _, c := range info.codes {
		if c == code {
			return true
		}
	}
	return false
}

func(info *RegionInfo) AddCodes(codes ...string) error {
	if info == nil {
		return goutil.NewNilTargetError(&RegionInfo{}, "AddCodes")
	}
	if info.id == RG_NONE {
		return errors.New("Trying to call RegionInfo.AddCodes() when ID() == RG_NONE")
	}
	regionMutex.Lock()
	if regionNumericMap == nil {
		regionNumericMap = make(map[RegionNumeric]*RegionInfo)
		regionCodeMap = make(map[string]*RegionInfo)
	}
	for _, code := range codes {
		if len(code) == 0 {
			continue
		}
		prev := regionCodeMap[code]
		if prev != nil && prev.id != info.id {
			regionMutex.Unlock()
			return errors.New(fmt.Sprintf(
				"Cannot add string code '%s' to region %d: Region %d already has this code",
				code,
				uint32(info.id),
				uint32(prev.id),
			))
		}
	}
	for _, code := range codes {
		if len(code) > 0 {
			regionCodeMap[code] = info
		}
	}
	regionMutex.Unlock()
	return nil
}

func GetRegionInfo(region Region) (info *RegionInfo) {
	if region == RG_NONE {
		return
	}
	regionMutex.Lock()
	if region <= Region(len(regions)) {
		info = regions[region - 1]
	}
	regionMutex.Unlock()
	return
}

func GetRegionByNumeric(numeric RegionNumeric) (info *RegionInfo) {
	if numeric == RGNM_NONE {
		return
	}
	regionMutex.Lock()
	if regionNumericMap != nil {
		info = regionNumericMap[numeric]
	}
	regionMutex.Unlock()
	return
}

func GetRegionByCode(code string) (info *RegionInfo) {
	if len(code) == 0 {
		return
	}
	regionMutex.Lock()
	if regionCodeMap != nil {
		info = regionCodeMap[code]
	}
	regionMutex.Unlock()
	return
}

func initRegions() {
	// from https://en.wikipedia.org/wiki/UN_M49
	//TODO
	// from https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	//TODO
}
