Fork of [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2) with minimal upstream proxy and routing support.

All features of the original Shadowsocks in Golang plus:
- `-upstream '127.0.0.1:8384'` flag configures an upstream Shadowsocks server which traffic will be forwarded to.
Server must use the same encryption protocol and password which current server does - no option to customise it at the moment.
- `-upstream-targets 'google.com,m.facebook.com,127.0.0.1'` flag configures targets which should be routed to the upstream.
A comma-separated list of host names or IP addresses. If an item is a valid IP address - it will be matched to the target
IP exactly. If an item is not a valid IP address - it's considered a host name and will be resolved to IP via DNS lookup.
DNS lookup is performed in the background once per minute for all hostname items. Target IP must match at least one item from the list
according to the rules above in order to be routed to the upstream. If target does not match any of them - 
traffic routed directly to the Internet.
