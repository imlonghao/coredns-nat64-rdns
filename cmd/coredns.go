package main

import (
	_ "git.esd.cc/imlonghao/coredns-nat64-rdns"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/dns64"
	_ "github.com/coredns/coredns/plugin/forward"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

var directives = []string{
	"bind",
	"forward",
	"dns64",
	"nat64-rdns",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
