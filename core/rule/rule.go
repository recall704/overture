package rule

type RuleType string

const (
	RuleTypeDomain        = "Domain"
	RuleTypeSuffix        = "DomainSuffix"
	RuleTypeDomainKeyword = "DomainKeyword"
)

type Rule interface {
	RuleType() RuleType
	Match(domain string) bool
}
