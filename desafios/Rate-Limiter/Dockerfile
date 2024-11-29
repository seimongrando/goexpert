# Usa a imagem oficial do Go como base
FROM golang:1.23-alpine

# Instala dependências necessárias para build e runtime
RUN apk add --no-cache build-base ca-certificates

# Define o diretório de trabalho no container
WORKDIR /app

# Copia todos os arquivos do projeto para o contêiner
COPY . .

# Baixa as dependências e compila o binário
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/rate-limiter

# Exponha a porta 8080
EXPOSE 8080

# Define o comando padrão para rodar o executável
CMD ["/app/rate-limiter"]
