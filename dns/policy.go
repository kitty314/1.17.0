package dns

import (
	"github.com/kitty314/1.17.0/component/trie"
	C "github.com/kitty314/1.17.0/constant"
	"github.com/kitty314/1.17.0/constant/provider"
)

type dnsPolicy interface {
	Match(domain string) []dnsClient
}

type domainTriePolicy struct {
	*trie.DomainTrie[[]dnsClient]
}

func (p domainTriePolicy) Match(domain string) []dnsClient {
	record := p.DomainTrie.Search(domain)
	if record != nil {
		return record.Data()
	}
	return nil
}

type geositePolicy struct {
	matcher    fallbackDomainFilter
	inverse    bool
	dnsClients []dnsClient
}

func (p geositePolicy) Match(domain string) []dnsClient {
	matched := p.matcher.Match(domain)
	if matched != p.inverse {
		return p.dnsClients
	}
	return nil
}

type domainSetPolicy struct {
	domainSetProvider provider.RuleProvider
	dnsClients        []dnsClient
}

func (p domainSetPolicy) Match(domain string) []dnsClient {
	metadata := &C.Metadata{Host: domain}
	if ok := p.domainSetProvider.Match(metadata); ok {
		return p.dnsClients
	}
	return nil
}
