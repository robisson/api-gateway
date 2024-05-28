package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Handler representa o manipulador de requisições HTTP
type Handler struct {
	config  *OpenAPI
	proxies map[string]*httputil.ReverseProxy
	targets map[string]*url.URL
}

// NewHandler cria um novo manipulador com base na configuração fornecida
func NewHandler(config *OpenAPI) *Handler {
	proxies := make(map[string]*httputil.ReverseProxy)
	targets := make(map[string]*url.URL)
	for _, methods := range config.Paths {
		for _, operation := range methods {
			if _, exists := proxies[operation.BackendURL]; !exists {
				targetURL, _ := url.Parse(operation.BackendURL)
				proxies[operation.BackendURL] = httputil.NewSingleHostReverseProxy(targetURL)
				targets[operation.BackendURL] = targetURL
			}
		}
	}
	return &Handler{config: config, proxies: proxies, targets: targets}
}

// ServeHTTP manipula as requisições HTTP roteando-as para o backend apropriado
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToLower(r.Method)

	// Buscar a configuração do caminho na especificação OpenAPI
	var operation Operation
	var found bool

	for route, methods := range h.config.Paths {
		if matches := matchRoute(route, path); matches {
			if op, ok := methods[method]; ok {
				operation = op
				found = true
				break
			}
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	proxy := h.proxies[operation.BackendURL]
	targetURL := h.targets[operation.BackendURL]
	if proxy == nil || targetURL == nil {
		http.Error(w, "Proxy não encontrado", http.StatusInternalServerError)
		return
	}

	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	r.Host = targetURL.Host
	proxy.ServeHTTP(w, r)
}

// matchRoute verifica se o caminho da requisição corresponde à rota definida
func matchRoute(route, path string) bool {
	routeParts := strings.Split(route, "/")
	pathParts := strings.Split(path, "/")

	if len(routeParts) != len(pathParts) {
		return false
	}

	for i, routePart := range routeParts {
		if strings.HasPrefix(routePart, "{") && strings.HasSuffix(routePart, "}") {
			continue
		}
		if routePart != pathParts[i] {
			return false
		}
	}

	return true
}
