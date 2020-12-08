# ebanx-challenge

Este projeto é a implementação do desafio da Ebanx. Uma API para registro de transações entre contas.

### Ferramentas
- Linguagem: [Go](https://golang.org/ "Go")
- Criação de mocks: [mockery](https://github.com/vektra/mockery "mockery")

### Dependências
- Gin (go get github.com/gin-gonic/gin)
- Viper (go get github.com/spf13/viper)
- Testify (go get github.com/stretchr/testify)

### Arquivo de configuração: *config.yml*
Exemplo:
```yaml
server:
  address: 127.0.0.1:8080
```
- **server**: configurações do servidor da API
	- **address**: endereço e porta de acesso à API

#### Adicionar um evento de transação

> Método: POST
> Endpoint: /v1/event

- **Campos do corpo**:
  - **type**: tipo de transação para executar [_obrigatório_]
    - **deposit**: faz um depósito na conta indicada no campo "destination"
    - **withdraw**: faz uma retirada da conta indicada no campo "origin"
    - **transfer**: faz uma transferência da conta do campo "origin" para a conta do campo "destination"
  - **amount**: valor da transação [_obrigatório_]
  - **destination**: conta de destino da transação [obrigatório _se o campo "type" for "deposit" ou "transfer"_]
  - **origin**: conta de destino da transação [obrigatório _se o campo "type" for "withdraw" ou "transfer"_]

##### Exemplo requisições:
> POST /v1/event
```json
{
    "type": "deposit",
    "destination": "100",
    "amount": 10
}
```
> POST /v1/event
```json
{
    "type": "withdraw",
    "origin": "100",
    "amount": 5
}
```
> POST /v1/event
```json
{
    "type": "transfer",
    "origin": "100",
    "amount": 15,
    "destination": "300"    
}
```

##### Exemplo respostas:
> 201 - Created
```json
{
    "destination": {
        "id": "100",
        "balance": 10
    }
}
```
> 201 - Created
```json
{
    "origin": {
        "id": "100",
        "balance": 15
    }
}
```
> 201 - Created
```json
{
    "origin": {
        "id": "100",
        "balance": 0
    },
    "destination": {
        "id": "300",
        "balance": 15
    }
}
```
> 404 - Not Found
```json
0
```

#### Buscar balanço de uma conta

> Método: GET
> Endpoint: /v1/balance?account_id={conta}

##### Exemplo requisição:
> GET /v1/balance?account_id=100

##### Exemplo resposta:
> 200 - OK
```json 
20
```

#### Resetar estado da API

> Método: POST
> Endpoint: /v1/reset

------------

#### Host remoto

https://ebanx-challenge.herokuapp.com

#### Usando localmente:
Para rodar a aplicação localmente é necessário executar os seguintes passos:
1. Instalar as ferramentas abaixo na máquina local:
	- Go v1.14.6+
2. Clonar esse repositório em qualquer diretório
3. Alterar o arquivo *config.yml* com as configurações desejadas
4. No diretório clonado, rodar a aplicação usando: **go run api/main.go**
5. Se desejar, executar os testes com o comando: **go test ./...**