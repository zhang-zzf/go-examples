package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

func FromRequest(r *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not ip:port", r.RemoteAddr)
	}
	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		return nil, fmt.Errorf("userip: %q is not ip:port", r.RemoteAddr)
	}
	return parseIP, nil
}

type key int

const userIPKey key = 0

func NewContext(ctx context.Context, ip net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, ip)
}

func FromContext(ctx context.Context) (net.IP, bool) {
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
