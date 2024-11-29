# **Rate Limiter**

Este projeto implementa um **Rate Limiter** em Go, configurado para limitar o número de requisições com base no **IP** ou em um **Token de Acesso**. Ele utiliza armazenamento em memória para simplificar a gestão de requisições e é construído para ser escalável e fácil de configurar.

---

## **Funcionalidades**
- **Limitação por IP**: Controla o número de requisições por segundo por endereço IP.
- **Limitação por Token**: Permite definir limites diferentes para tokens de acesso.
- **Sobreposição de Limites**: As configurações de limite do token têm precedência sobre as do IP.
- **Configuração Simples**: Utiliza variáveis de ambiente para definir os limites e o tempo de bloqueio.
- **Middleware**: Implementado como um middleware para fácil integração com servidores HTTP.
- **Armazenamento em Memória**: Gerencia as requisições usando um mapa em memória com `sync.Mutex` para segurança em concorrência.
- **Armazenamento Redis**: Gerencia as requisições usando o redis para melhor performance.
- **Docker**: Inclui suporte para build e execução em containers Docker.

---

## **Pré-requisitos**
- **Go** (versão 1.23 ou superior)
- **Docker** e **Docker Compose**

---

## **Configurações**
As configurações são feitas por meio de variáveis de ambiente, definidas no arquivo `.env` ou diretamente no `docker-compose.yml`. As variáveis disponíveis são:

| Variável                   | Descrição                                   | Valor Padrão     |
|----------------------------|---------------------------------------------|------------------|
| `RATE_LIMITER_IP_LIMIT`    | Limite de requisições por segundo por IP    | `10`             |
| `RATE_LIMITER_TOKEN_LIMIT` | Limite de requisições por segundo por Token | `100`            |
| `RATE_LIMITER_BLOCK_TIME`  | Tempo de bloqueio em segundos               | `300`            |
| `SERVER_PORT`              | Porta onde o servidor será executado        | `8080`           |
| `REDIS_ADDR`               | Endereço do redis                           | `localhost:6379` |
| `REDIS_PASSWORD`           | Senha do redis                              | ``               |
| `REDIS_DB`                 | Nome do servidor redis                      | `0`              |

Exemplo do `.env`:
```
RATE_LIMITER_IP_LIMIT=10
RATE_LIMITER_TOKEN_LIMIT=100
RATE_LIMITER_BLOCK_TIME=300
SERVER_PORT=8080
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

---

## **Como Buildar Localmente**
1. **Clone o Repositório**:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd rate-limiter
   ```

2. **Instale as Dependências**:
   ```bash
   go mod tidy
   ```

3. **Compile o Binário**:
   ```bash
   go build -o rate-limiter .
   ```

4. **Execute o Servidor**:
   ```bash
   ./rate-limiter
   ```

O servidor estará disponível na porta **80**.

---

## **Como Construir e Executar com Docker**
1. **Build da Imagem**:
   Execute o comando abaixo para criar a imagem do Docker:
   ```bash
   docker-compose build
   ```

2. **Subir o Container**:
   ```bash
   docker-compose up
   ```

3. **Testar a Aplicação**:
   Acesse `http://localhost:8080` ou use `curl`:
   ```bash
   curl http://localhost:8080
   ```

---

## **Testando o Rate Limiter**
### **Cenário 1: Limitação por Token**
1. Use um cabeçalho com o token:
   ```bash
   curl -H "API_KEY: token123" http://localhost:8080
   ```
2. Envie requisições rápidas consecutivas. Após exceder o limite configurado, a resposta será:
   ```
   HTTP/1.1 429 Too Many Requests
   you have reached the maximum number of requests or actions allowed within a certain time frame
   ```

### **Cenário 2: Limitação por IP**
1. Teste sem incluir o cabeçalho `API_KEY`:
   ```bash
   curl http://localhost:8080
   ```
2. O comportamento será o mesmo do token, mas o limite será baseado no IP.

---

## **Estrutura do Projeto**
```
rate-limiter/
├── config/                # Pacote de configuração do projeto
│   └── config.go
├── limiter/               # Lógica do Rate Limiter
│   ├── limiter.go         # Gerenciador principal do Rate Limiter
│   └── interfaces.go      # Mapeia as dependencies para utilizaz o limiter
│   └── types.go           # Armazena os tipos
├── middleware/            
│   └── rate_limiter.go    # Middleware para integração com servidores HTTP
│   └── interfaces.go      # Mapeia as dependencies para utilizar o middleware  
├── strategy/              # Algoritmos para salvar as requisições
│   ├── memory.go          # Implementação em memória com mapa
│   ├── redis.go           # Implementação utilizando redis
│   └── types.go           # Armazena os tipos
├── Dockerfile             # Configuração para build do container
├── docker-compose.yml     # Configuração do Docker Compose
├── main.go                # Ponto de entrada da aplicação
├── go.mod                 # Dependências do projeto
├── go.sum                 # Checksum das dependências
└── README.md              # Documentação do projeto
```

---

## **Testes Unitários / Integrado**

Este projeto possui testes para validar as funcionalidades principais, incluindo o middleware, o Rate Limiter, e o armazenamento em memória. Para executar todos os testes, siga os passos abaixo:

1. Certifique-se de que você tem o Go instalado (versão 1.21 ou superior).
2. No diretório raiz do projeto, execute o seguinte comando:
   ```bash
   go test ./... -v
   ```
   
