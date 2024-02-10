package Proxy

import "net/url"

type ProxyService interface {
	Init(Proxy) error
	GetAll() []*url.URL
	GetProxy() *url.URL
	ReportProxy(addr *url.URL, reason string) *url.URL
	GetProxyCountry(*url.URL) string
	Done() error
}
