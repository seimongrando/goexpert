# **Load Tester CLI**

Uma aplicação CLI escrita em Go para realizar testes de carga em serviços web. Com ela, você pode definir o número de requisições, nível de concorrência e gerar relatórios detalhados sobre o desempenho do serviço.

---

### **Como Funciona**

O `Load Tester CLI` realiza múltiplas requisições HTTP paralelamente, analisando os tempos de resposta e os códigos de status retornados. É possível configurar a URL, o número de requisições totais e a quantidade de chamadas simultâneas.

---

### **Pré-requisitos**

- **Docker**: Certifique-se de que o Docker está instalado.
- **Go** (opcional): Para rodar ou testar localmente.

---

### **Como Buildar**

#### **Com Docker**
1. Build da imagem Docker:
   ```bash
   docker build -t load-tester .
   ```

2. Rodar a aplicação:
   ```bash
   docker run load-tester test --url=https://www.google.com --requests=100 --concurrency=10
   ```

#### **Localmente**
1. Clone o repositório e instale as dependências:
   ```bash
   go mod tidy
   ```

2. Compile o binário:
   ```bash
   go build -o load-tester .
   ```

3. Execute:
   ```bash
   ./load-tester test --url=https://www.google.com --requests=100 --concurrency=10
   ```

---

### **Como Testar**

#### **Testes Unitários**
Os testes verificam a funcionalidade do sistema de teste de carga (`stress`).
1. Execute os testes com:
   ```bash
   go test ./stress -v
   ```

---

### **Estratégia do `entrypoint.sh`**

Para aceitar e repassar corretamente os argumentos fornecidos ao contêiner, utilizamos um script chamado `entrypoint.sh`. Ele reenvia todos os argumentos para o binário `load-tester`.

**Conteúdo do `entrypoint.sh`:**
```sh
#!/bin/sh
./load-tester "$@"
```

**Como funciona:**
- O `ENTRYPOINT` no Dockerfile é configurado para usar o `entrypoint.sh`.
- Quando você executa o contêiner com argumentos, eles são enviados para o script, que os repassa para o binário.

---

### **Exemplo de Uso**

Rodar um teste de carga com 100 requisições e 10 chamadas simultâneas:
```bash
docker run load-tester test --url=https://www.google.com --requests=100 --concurrency=10
```

---

### **Saída do Relatório**
Ao final do teste, será exibido um relatório como este:
```plaintext
--- Teste de Carga Concluído ---
Tempo total: 5.12s
Total de requests: 100
Requests com status 200: 98
Distribuição de outros códigos de status:
  Status 500: 2
```
