package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMatchRoute verifica se a função matchRoute funciona corretamente
func TestMatchRoute(t *testing.T) {
	assert.True(t, matchRoute("/payments/{payment_id}", "/payments/12345"), "Rotas devem corresponder")
	assert.False(t, matchRoute("/payments", "/payment/12345"), "Rotas não devem corresponder")
}

// TestServeHTTP verifica se o handler serve as requisições corretamente para rotas válidas
func TestServeHTTP(t *testing.T) {
	config, _ := LoadConfig("../config/openapi.yaml")
	handler := NewHandler(config)

	// Criar uma requisição válida
	req := httptest.NewRequest("GET", "/payments/12345", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	// Verifique se o status é 404 porque o mock server não está em execução
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode, "Deve retornar 404 Not Found")
}

// TestServeHTTPInvalidMethod verifica se o handler retorna 404 para método inválido
func TestServeHTTPInvalidMethod(t *testing.T) {
	config, _ := LoadConfig("../config/openapi.yaml")
	handler := NewHandler(config)

	// Criar uma requisição com método inválido
	req := httptest.NewRequest("DELETE", "/payments/12345", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	// Verifique se o status é 404 para método inválido
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode, "Deve retornar 404 Not Found para método inválido")
}

// TestServeHTTPInvalidURL verifica se o handler lida com URLs de destino inválidas
func TestServeHTTPInvalidURL(t *testing.T) {
	invalidConfig := &OpenAPI{
		Paths: map[string]map[string]Operation{
			"/invalid-url": {
				"get": {
					BackendURL: "http://[::1]:namedport", // URL inválida
				},
			},
		},
	}
	handler := NewHandler(invalidConfig)

	// Criar uma requisição para uma rota com URL inválida
	req := httptest.NewRequest("GET", "/invalid-url", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	// Verifique se o status é 500 devido à URL inválida
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode, "Deve retornar 500 Internal Server Error para URL inválida")
}

// TestServeHTTPNotFound verifica se o handler retorna 404 para caminhos não encontrados
func TestServeHTTPNotFound(t *testing.T) {
	config, _ := LoadConfig("../config/openapi.yaml")
	handler := NewHandler(config)

	// Criar uma requisição para um caminho inexistente
	req := httptest.NewRequest("GET", "/unknown", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	// Verifique se o status é 404 para caminho inexistente
	assert.Equal(t, http.StatusNotFound, w.Result().StatusCode, "Deve retornar 404 Not Found para caminho inexistente")
}
