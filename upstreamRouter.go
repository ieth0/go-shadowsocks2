package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func unique(s []net.IP) []net.IP {
	inResult := make(map[string]bool)
	var result []net.IP
	for _, ip := range s {
		ipstr := ip.String()
		if _, ok := inResult[ipstr]; !ok {
			inResult[ipstr] = true
			result = append(result, ip)
		}
	}
	return result
}

type UpstreamRouter struct {
	Hosts []string
	Ips   []net.IP
}

func (router UpstreamRouter) Resolve() {
	for {
		newIps := router.Ips
		for _, host := range router.Hosts {
			t, _ := net.LookupIP(host)
			newIps = append(newIps, t...)
		}
		router.Ips = unique(newIps)

		logf("Next IPs are routed to upstream: %s\n", router.Ips)
		time.Sleep(time.Minute)
	}
}

func (router UpstreamRouter) shouldRoute(addr string) bool {
	if router.Ips == nil {
		return false
	}

	host, _, _ := net.SplitHostPort(addr)
	ip := net.ParseIP(host)

	if ip != nil {
		for _, v := range router.Ips {
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
