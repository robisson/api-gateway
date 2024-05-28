package gateway

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfig verifica se a configuração é carregada corretamente do arquivo
func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../config/openapi.yaml")
	assert.NoError(t, err, "Erro ao carregar configuração")
	assert.NotNil(t, config, "Configuração não deve ser nula")
	assert.Contains(t, config.Paths, "/payments", "Configuração deve conter /payments")
}

// TestLoadConfigFileNotFound verifica o comportamento quando o arquivo de configuração não é encontrado
func TestLoadConfigFileNotFound(t *testing.T) {
	_, err := LoadConfig("invalid/path/to/openapi.yaml")
	assert.Error(t, err, "Deve retornar um erro quando o arquivo não é encontrado")
}

// TestLoadConfigInvalidYAML verifica o comportamento quando o conteúdo do arquivo YAML é inválido
func TestLoadConfigInvalidYAML(t *testing.T) {
	invalidYAMLPath := "invalid.yaml"
	invalidYAMLContent := []byte(`
openapi: 3.0.0
info:
  title: "Invalid YAML"
  version: 1.0.0
paths:
  /payments:
    post:
      summary: "Create a new payment"
      x-backend-url: "http://localhost:8081
`) // Intencionalmente malformado

	// Cria um arquivo YAML inválido temporariamente
	err := ioutil.WriteFile(invalidYAMLPath, invalidYAMLContent, 0644)
	assert.NoError(t, err, "Deve ser possível criar um arquivo YAML inválido temporariamente")

	_, err = LoadConfig(invalidYAMLPath)
	assert.Error(t, err, "Deve retornar um erro ao tentar desserializar um YAML inválido")

	// Limpeza: Remove o arquivo YAML inválido
	err = os.Remove(invalidYAMLPath)
	assert.NoError(t, err, "Deve ser possível remover o arquivo YAML inválido temporariamente")
}
