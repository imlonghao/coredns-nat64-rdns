package coredns_nat64_rdns

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("nat64-rdns", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'nat64-rdns'

	suffix := ""

	if c.Next() {
		suffix = c.Val()
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Nat64rDNS{Next: next, Suffix: suffix}
	})

	return nil
}
