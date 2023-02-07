package main

import (
	_ "git.esd.cc/imlonghao/coredns-nat64-rdns"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

var directives = []string{
	"nat64-rdns",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	coremain.Run()
}
