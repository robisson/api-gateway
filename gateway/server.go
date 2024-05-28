package gateway

import (
	"net/http"
)

// Server representa o servidor HTTP do API Gateway
type Server struct {
	httpServer *http.Server
}

// NewServer cria um novo servidor com a configuração fornecida
func NewServer(config *OpenAPI) *Server {
	handler := NewHandler(config)
	return &Server{
		httpServer: &http.Server{
			Handler: handler,
		},
	}
}

// ListenAndServe inicia o servidor HTTP na porta especificada
func (s *Server) ListenAndServe(addr string) error {
	s.httpServer.Addr = addr
	return s.httpServer.ListenAndServe()
}
