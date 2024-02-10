package Proxy

import "net/url"

var AvailableProxyServices = make(map[string]ProxyService)

type Proxy struct {
	Strategy string
	Addr     *url.URL
	Others   map[string]interface{}
}

type singleProxyStrategy struct {
	proxyAddr *url.URL
}

const ProxyTypeSingle = "single"
