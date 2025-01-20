### README - Weather Service

---

#### **Descrição do Projeto**

#### **Objetivo:** Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

#### **Requisitos:**

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
  - Em caso de sucesso:
  Código HTTP: 200
  Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
  Em caso de falha, caso o CEP não seja válido (com formato correto):
  Código HTTP: 422
  Mensagem: invalid zipcode
  ​​​Em caso de falha, caso o CEP não seja encontrado:
  Código HTTP: 404
  Mensagem: can not find zipcode
  Deverá ser realizado o deploy no Google Cloud Run.
  Dicas:

- Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
- Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
  - Sendo F = Fahrenheit
  - Sendo C = Celsius
  - Sendo K = Kelvin

---

#### **Validações**

- **Entrada**: CEP válido de 8 dígitos.
- **Saída**:
    - Em caso de sucesso (HTTP 200):
      ```json
      {
          "temp_C": 28.5,
          "temp_F": 83.3,
          "temp_K": 301.65
      }
      ```
    - CEP inválido (HTTP 422):
      ```json
      {
          "message": "invalid zipcode"
      }
      ```
    - CEP não encontrado (HTTP 404):
      ```json
      {
          "message": "can not find zipcode"
      }
      ```

---

#### **Tecnologias Utilizadas**

- **Linguagem**: Go
- **Framework**: Gin Gonic
- **APIs**:
    - [ViaCEP](https://viacep.com.br/) - Consulta de CEPs.
    - [WeatherAPI](https://www.weatherapi.com/) - Consulta de dados climáticos.
- **Docker**: Para containerização.
- **Google Cloud Run**: Para deploy.

---

#### **Passos para Configuração Local**

1. **Clonar o Repositório**
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd weather-service
   ```

2. **Configurar Dependências**
    - Certifique-se de ter o Go instalado (versão 1.21 ou superior).
    - Baixe as dependências:
      ```bash
      go mod tidy
      ```

3. **Executar o Serviço**
   ```bash
   go run main.go
   ```
   O serviço ficará disponível em: `http://localhost:8080`.

4. **Testar Localmente**
    - Envie uma requisição:
      ```bash
      curl http://localhost:8080/weather/01001000
      ```

---

#### **Testes Automatizados**

1. Execute os testes com o comando:
   ```bash
   go test ./...
   ```

---

#### **Passos para Deploy no Google Cloud Run**

1. **Build da Imagem Docker**
   ```bash
   docker build -t gcr.io/PROJECT-ID/weather-service .
   ```

2. **Push da Imagem para o Google Container Registry**
   ```bash
   docker push gcr.io/PROJECT-ID/weather-service
   ```

3. **Deploy no Google Cloud Run**
   ```bash
   gcloud run deploy weather-service \
       --image gcr.io/PROJECT-ID/weather-service \
       --platform managed \
       --region REGION \
       --allow-unauthenticated
   ```

4. **Acessar o Endpoint**
   Após o deploy, o Google Cloud Run fornecerá um URL público. Você pode usar esse URL para realizar requisições, como:
   ```bash
   curl https://<URL_FORNECIDA>/weather/01001000
   ```

---

#### **Testando com Docker**

1. **Build da Imagem**
   ```bash
   docker build -t weather-service .
   ```

2. **Executar o Contêiner**
   ```bash
   docker run -p 8080:8080 weather-service
   ```

3. **Testar a Aplicação**
   ```bash
   curl http://localhost:8080/weather/01001000
   ```

---

#### **Notas**

- Substitua `PROJECT-ID` pelo ID do projeto no Google Cloud.
- Certifique-se de ter uma chave válida da [WeatherAPI](https://www.weatherapi.com/) e configure-a no código em `getWeather`.
