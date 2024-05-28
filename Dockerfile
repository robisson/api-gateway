# Utiliza a imagem oficial do Golang com Alpine como base para a construção
FROM golang:1.20-alpine AS builder

# Instala as dependências necessárias
RUN apk add --no-cache git

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos go.mod e go.sum e baixa as dependências
ENV GOPROXY=direct 
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte da aplicação
COPY . .

# Compila a aplicação
RUN go build -o /api-gateway main.go

# Utiliza uma imagem mínima do Alpine para executar o binário compilado
FROM alpine:3.14

# Copia o binário compilado da imagem de construção
COPY --from=builder /api-gateway /api-gateway

# Copia o arquivo de configuração
COPY config/openapi.yaml /config/openapi.yaml

# Define o comando de entrada para o contêiner
ENTRYPOINT ["/api-gateway"]

# Expõe a porta que a aplicação utilizará
EXPOSE 8080
