package main

import (
	"api-gateway/gateway"
	"log"
)

func main() {
	// Carrega a configuração do arquivo OpenAPI
	config, err := gateway.LoadConfig("config/openapi.yaml")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Inicia o servidor com a configuração carregada
	server := gateway.NewServer(config)
	log.Println("Iniciando API Gateway na porta 8080...")
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
