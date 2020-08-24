package rule

import "strings"

type DomainKeyword struct {
	keyword string
}

func (ds *DomainKeyword) RuleType() RuleType {
	return RuleTypeDomainKeyword
}

func (ds *DomainKeyword) Match(domain string) bool {
	return strings.Contains(domain, ds.keyword)
}

func NewDomainKeyword(keyword string) *DomainKeyword {
	return &DomainKeyword{
		keyword: strings.ToLower(keyword),
	}
}