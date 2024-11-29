FROM golang:1.23-alpine

# Instala dependências necessárias para build e runtime
RUN apk add --no-cache build-base ca-certificates

# Define o diretório de trabalho no container
WORKDIR /app

# Copia todos os arquivos do projeto para o contêiner
COPY . .

# Baixa as dependências e compila o binário
RUN go mod tidy && go build -o load-tester .

# Copiar o script de entrada
COPY entrypoint.sh /app/entrypoint.sh

# Dar permissão de execução ao script
RUN chmod +x /app/entrypoint.sh

# Usar o script como ponto de entrada
ENTRYPOINT ["/app/entrypoint.sh"]
