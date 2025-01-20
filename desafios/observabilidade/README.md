### README - Sistema de Temperatura por CEP com Observabilidade

---

### **Descrição do Projeto**

Este projeto consiste em dois serviços desenvolvidos em Go que, juntos, permitem consultar a cidade e a temperatura atual de um CEP fornecido. Ele também implementa **observabilidade** com **OpenTelemetry** e **Zipkin** para rastreamento distribuído entre os serviços.

---

### **Serviços**

#### **Serviço A - Entrada**
1. Recebe um CEP via **POST** no formato:
   ```json
   { "cep": "29902555" }
   ```
2. Valida se o CEP contém **8 dígitos** e é uma **string**.
3. Encaminha o CEP ao **Serviço B** via HTTP.
4. Retorna a resposta do Serviço B ao cliente.

- **Endpoints**:
    - **POST /cep**:
        - Exemplo de Requisição:
          ```bash
          curl -X POST http://localhost:8081/cep \
               -H "Content-Type: application/json" \
               -d '{"cep": "01001000"}'
          ```

#### **Serviço B - Orquestração**
1. Recebe o CEP do Serviço A.
2. Consulta a cidade correspondente ao CEP via [ViaCEP](https://viacep.com.br/).
3. Consulta a temperatura atual da cidade via [WeatherAPI](https://www.weatherapi.com/).
4. Retorna as informações ao Serviço A no formato:
   ```json
   {
     "city": "São Paulo",
     "temp_C": 28.5,
     "temp_F": 83.3,
     "temp_K": 301.65
   }
   ```

- **Endpoints**:
    - **GET /weather/:zipcode**:
        - Exemplo de Requisição:
          ```bash
          curl -X GET http://localhost:8080/weather/01001000
          ```

---

### **Recursos de Observabilidade**

1. **OpenTelemetry**:
    - Implementado em ambos os serviços para rastreamento de spans.
    - Mede o tempo de execução das consultas de CEP e temperatura.

2. **Zipkin**:
    - Coletor e painel para visualização dos spans.
    - Acesse o painel em: [http://localhost:9411](http://localhost:9411).

---

### **Configuração e Execução**

#### **Pré-requisitos**
1. [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/).
2. Chave válida para a [WeatherAPI](https://www.weatherapi.com/).

#### **Passos**
1. Clone o repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd observabilidade
   ```

2. Insira sua chave da WeatherAPI no arquivo de ambiente do **Serviço B**:
    - Edite `docker-compose.yml`:
      ```yaml
      service-b:
        environment:
          - WEATHER_API_TOKEN=<SUA_CHAVE_API>
      ```

3. Execute os serviços:
   ```bash
   docker-compose up --build
   ```

4. Teste os serviços:
    - **Serviço A**:
      ```bash
      curl -X POST http://localhost:8081/cep \
           -H "Content-Type: application/json" \
           -d '{"cep": "01001000"}'
      ```
    - **Serviço B** (diretamente):
      ```bash
      curl -X GET http://localhost:8080/weather/01001000
      ```

5. Acesse o painel do Zipkin para visualizar os rastreamentos:
    - [http://localhost:9411](http://localhost:9411).

---

### **Estrutura de Diretórios**

```plaintext
observabilidade/
│
├── docker-compose.yml       # Orquestração dos serviços e Zipkin
├── service-a/               # Código-fonte do Serviço A
│   ├── Dockerfile
│   ├── main.go
│   └── ...
├── service-b/               # Código-fonte do Serviço B
│   ├── Dockerfile
│   ├── main.go
│   └── ...
```

---

### **Recursos e Dependências**

#### **Bibliotecas Utilizadas**
1. **Gin Gonic**: Framework web para APIs.
2. **OpenTelemetry**:
    - `go.opentelemetry.io/otel`
    - `go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp`
3. **Zipkin**: Exportador de spans.
4. **Resty**: Cliente HTTP para requisições.

---

### **Possíveis Problemas e Soluções**

1. **Erro: `connection refused` ao enviar spans para o Zipkin**:
    - Certifique-se de que o serviço `zipkin` está ativo.
    - Verifique se o endpoint `http://zipkin:9411/api/v2/spans` está correto.

2. **Chave da WeatherAPI inválida**:
    - Obtenha uma nova chave em [WeatherAPI](https://www.weatherapi.com/).

---
