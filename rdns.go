package coredns_nat64_rdns

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"sort"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

const name = "nat64-rdns"

func Reverse(s interface{}) {
	sort.SliceStable(s, func(i, j int) bool {
		return true
	})
}

type Nat64rDNS struct {
	Next   plugin.Handler
	Suffix string
}

// ServeDNS implements the plugin.Handler interface.
func (nr Nat64rDNS) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	a := new(dns.Msg)
	a.SetReply(r)
	a.Authoritative = true

	reqSplited := strings.Split(state.Name(), string('.'))
	if len(reqSplited) < 8 {
		return plugin.NextOrFailure(nr.Name(), nr.Next, ctx, w, r)
	}
	t := reqSplited[:8]
	Reverse(t)
	ip, err := hex.DecodeString(strings.Join(t, ""))
	if err != nil {
		return plugin.NextOrFailure(nr.Name(), nr.Next, ctx, w, r)
	}
	originIP := net.IP(ip)

	var rr dns.RR

	switch state.QType() {
	case dns.TypePTR:
		rr = new(dns.PTR)
		rr.(*dns.PTR).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypePTR, Class: state.QClass(), Ttl: 86400}
		rr.(*dns.PTR).Ptr = fmt.Sprintf("%s.%s", originIP, nr.Suffix)
	default:
		return plugin.NextOrFailure(nr.Name(), nr.Next, ctx, w, r)
	}

	a.Answer = []dns.RR{rr}

	w.WriteMsg(a)

	return 0, nil
}

// Name implements the Handler interface.
func (nr Nat64rDNS) Name() string { return name }
