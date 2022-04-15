package tools

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	RequestTimeout     = 1 * time.Minute
	InsecureSkipVerify = true
)

func NewProxy(proxy string) *http.Client {
	proxyClient, err := proxyHTTPClient(proxy)
	if err != nil {
		panic(err)
	}
	return proxyClient
}

func proxyHTTPClient(address string) (client *http.Client, err error) {
	tr := &http.Transport{
		Proxy:           nil,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: InsecureSkipVerify},
		DialContext:     (&net.Dialer{Timeout: RequestTimeout}).DialContext,
	}
	if len(address) > 0 {
		proxy, err := url.Parse(address)
		if err != nil {
			return nil, err
		}
		tr.Proxy = http.ProxyURL(proxy)
	}
	client = &http.Client{Transport: tr}
	return
}
