package rule

import "strings"

type DomainSuffix struct {
	suffix string
}

func (ds *DomainSuffix) RuleType() RuleType {
	return RuleTypeSuffix
}

func (ds *DomainSuffix) Match(domain string) bool {
	return strings.HasSuffix(domain, "."+ds.suffix) || domain == ds.suffix
}

func NewDomainSuffix(suffix string) *DomainSuffix {
	return &DomainSuffix{
		suffix: strings.ToLower(suffix),
	}
}
