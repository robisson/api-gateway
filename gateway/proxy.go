package gateway

import (
	"net/http/httputil"
	"net/url"
)

// NewProxy cria um novo proxy reverso para a URL de destino fornecida
func NewProxy(target string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}
