package goi18n

import (
	"sync"
)

type Domain uint

const NO_DOMAIN Domain = 0

type DomainInfo struct {
	Name string
	Localizer Localizer
}

var domains []DomainInfo

var domainMutex sync.Mutex

func NewDomain(name string, localizer Localizer) (dom Domain) {
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

func GetDomain(dom Domain) DomainInfo {
	id := int(dom)
	if id <= 0 || id > len(domains) {
		return DomainInfo{}
	}
	return domains[id - 1]
}
