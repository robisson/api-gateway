package gateway

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// OpenAPI define a estrutura para a configuração OpenAPI
type OpenAPI struct {
	Paths map[string]map[string]Operation `yaml:"paths"`
}

// Operation representa uma operação HTTP em um caminho
type Operation struct {
	BackendURL string `yaml:"x-backend-url"`
}

// LoadConfig carrega e desserializa a configuração OpenAPI do arquivo fornecido
func LoadConfig(filepath string) (*OpenAPI, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("falha ao ler arquivo: %w", err)
	}

	var config OpenAPI
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("falha ao deserializar configuração: %w", err)
	}

	return &config, nil
}
