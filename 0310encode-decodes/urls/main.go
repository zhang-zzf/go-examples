package main

import (
	"fmt"
	"net"
	"net/url"
)

var p = fmt.Println

func main() {
	urlStr := "postgres://user:pass@host.com:5432/path?k1=v1&k=v#f"
	u, err := url.Parse(urlStr)
	if err != nil {
		p(err)
	}
	fmt.Println(u.Scheme, u.User, u.Host, u.Port(), u.Path, u.RawQuery)
	p(net.SplitHostPort(u.Host))
	p(url.ParseQuery(u.RawQuery))

}
