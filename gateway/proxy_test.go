package gateway

import (
	"net/http/httputil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewProxy verifica se o proxy é criado corretamente para uma URL válida
func TestNewProxy(t *testing.T) {
	proxy, err := NewProxy("http://localhost:8081")
	assert.NoError(t, err, "Não deve haver erro ao criar um proxy válido")
	assert.NotNil(t, proxy, "O proxy não deve ser nulo")
	assert.IsType(t, &httputil.ReverseProxy{}, proxy, "Deve ser um proxy reverso do tipo httputil.ReverseProxy")
}

// TestNewProxyInvalidURL verifica se o proxy retorna um erro para uma URL inválida
func TestNewProxyInvalidURL(t *testing.T) {
	proxy, err := NewProxy("http://[::1]:namedport")
	assert.Error(t, err, "Deve haver erro ao criar um proxy com URL inválida")
	assert.Nil(t, proxy, "O proxy deve ser nulo para uma URL inválida")
}
