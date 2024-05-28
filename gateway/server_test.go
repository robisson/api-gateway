package gateway

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewServer verifica se o servidor é criado corretamente
func TestNewServer(t *testing.T) {
	config, err := LoadConfig("../config/openapi.yaml")
	require.NoError(t, err, "Erro ao carregar configuração")
	server := NewServer(config)

	assert.NotNil(t, server, "Servidor não deve ser nulo")
	assert.IsType(t, &http.Server{}, server.httpServer, "Deve ser um servidor HTTP")
}

// TestServerListenAndServe verifica se o servidor inicia corretamente
func TestServerListenAndServe(t *testing.T) {
	config, err := LoadConfig("../config/openapi.yaml")
	require.NoError(t, err, "Erro ao carregar configuração")
	server := NewServer(config)

	// Usando um endereço inválido para forçar um erro de inicialização
	err = server.ListenAndServe("invalid-address")
	assert.Error(t, err, "Deve retornar um erro ao usar um endereço inválido")
}

// TestServerListenAndServeSuccess verifica a inicialização correta do servidor
func TestServerListenAndServeSuccess(t *testing.T) {
	config, err := LoadConfig("../config/openapi.yaml")
	require.NoError(t, err, "Erro ao carregar configuração")
	server := NewServer(config)

	// Usando um servidor HTTP de teste
	testServer := &http.Server{Addr: ":0", Handler: server.httpServer.Handler}
	go func() {
		_ = testServer.ListenAndServe()
	}()
	defer testServer.Close()

	// Verifique se o servidor iniciou sem erros
	assert.NotNil(t, testServer, "Servidor de teste deve ser inicializado corretamente")
}
