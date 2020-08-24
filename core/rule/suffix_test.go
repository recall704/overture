package rule

import "testing"

func TestDomain_Suffix(t *testing.T) {

	r := NewDomainSuffix("google.com")
	result := r.Match("www.google.com")
	if !result {
		t.FailNow()
	}
}
