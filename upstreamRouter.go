package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type UpstreamRouter struct {
	Hosts []string
	Ips   **[]net.IP
}

func (router UpstreamRouter) Resolve() {
	for _, host := range router.Hosts {
		t, _ := net.LookupIP(host)
		newIps := append(**router.Ips, t...)
		*router.Ips = &newIps
	}

	time.Sleep(60 * 1000)
}

func (router UpstreamRouter) shouldRoute(addr string) bool {
	if router.Ips == nil {
		return false
	}

	host, _, _ := net.SplitHostPort(addr)
	ip := net.ParseIP(host)

	if ip != nil {
		for _, v := range **router.Ips {
			if ip.Equal(v) {
				return true
			}
		}
	}
	for _, v := range router.Hosts {
		if host == v || strings.HasSuffix(host, fmt.Sprintf(".%s", v)) {
			return true
		}
	}

	return false
}
