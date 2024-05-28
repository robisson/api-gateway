#!/bin/bash

go test ./... -v 
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html 

# Teste de criação de pagamento
echo "Criando pagamento..."
response=$(curl -s -X POST http://localhost:8080/payments \
    -H "Content-Type: application/json" \
    -d '{"account_id": "123456", "amount": 100.50}')

payment_id=$(echo $response | jq -r '.payment_id')
echo "Pagamento criado com ID: $payment_id"

# Pequeno delay para garantir que o pagamento foi processado
#sleep 1

# Teste de obtenção de detalhes do pagamento
echo "Obtendo detalhes do pagamento..."
details=$(curl -s -X GET http://localhost:8080/payments/$payment_id)
echo $details | jq .