package goi18n

import (
	"fmt"
	"sync"
)

type Domain uint

const NO_DOMAIN Domain = 0

type DomainInfo struct {
	Name string
	Localizer GenLocalizer
}

var domains []DomainInfo

var domainMutex sync.Mutex

func NewDomain(name string, localizer GenLocalizer) (dom Domain) {
	if localizer == nil {
		dom = NO_DOMAIN
		return
	}
	domainMutex.Lock()
	domains = append(domains, DomainInfo {
		Name: name,
		Localizer: localizer,
	})
	dom = Domain(len(domains))
	domainMutex.Unlock()
	return
}

func GetDomain(dom Domain) (info DomainInfo) {
	id := int(dom)
	if id <= 0 || id > len(domains) {
		return
	}
	domainMutex.Lock()
	info = domains[id - 1]
	domainMutex.Unlock()
	return
}

type DomainHandle[KeyT any] Domain

type DomainHandleInfo[KeyT any] struct {
	Name string
	Localizer Localizer[KeyT]
}

func NewDomainHandle[KeyT any](name string, localizer Localizer[KeyT]) DomainHandle[KeyT] {
	return DomainHandle[KeyT](NewDomain(name, localizer))
}

func GetDomainHandle[KeyT any](dom DomainHandle[KeyT]) DomainHandleInfo[KeyT] {
	info := GetDomain(Domain(dom))
	if info.Localizer == nil {
		return DomainHandleInfo[KeyT]{}
	}
	localizer, ok := info.Localizer.(Localizer[KeyT])
	if !ok {
		panic(fmt.Sprintf(
			"Bad I18N message key: Expected type %s for domain #%d (%s), but got %s",
			info.Localizer.MessageKeyTypeName(),
			uint(dom),
			info.Name,
			GetMessageKeyTypeName[KeyT](),
		))
	}
	return DomainHandleInfo[KeyT] {
		Name: info.Name,
		Localizer: localizer,
	}
}
