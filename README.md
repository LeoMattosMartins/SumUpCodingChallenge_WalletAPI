# Wallet API (Go)

A simple RESTful API to create wallets, check balances, deposit, and withdraw money. /
Uma API RESTful simples para criar carteiras, consultar saldos, depositar e sacar dinheiro.

---

## Tech Stack / Tecnologias Utilizadas

- Go  
- Gin (HTTP framework / framework HTTP)  
- GORM (ORM)  
- SQLite (Database / Banco de dados)  

---

## ⚙️ Setup / Configuração

```bash
git clone https://github.com/LeoMattosMartins/SumUpCodingChallenge_WalletAPI
cd wallet-api-go
go run .
```

---

## API Endpoints / Endpoints da API

| Method / Método | Endpoint                  | Description / Descrição           |
|------------------|---------------------------|-----------------------------------|
| POST             | /api/wallets              | Create a wallet / Criar carteira |
| GET              | /api/wallets/:id          | Get balance / Consultar saldo     |
| POST             | /api/wallets/:id/deposit  | Deposit money / Depositar dinheiro |
| POST             | /api/wallets/:id/withdraw | Withdraw money / Sacar dinheiro   |

---

## Testing / Testes

Use `curl` or Postman / Use `curl` ou Postman para testar os endpoints abaixo:

```bash
curl -X POST http://localhost:8080/api/wallets
curl -X POST http://localhost:8080/api/wallets/1/deposit -H "Content-Type: application/json" -d '{"amount": 50}'
curl -X POST http://localhost:8080/api/wallets/1/withdraw -H "Content-Type: application/json" -d '{"amount": 20}'
curl http://localhost:8080/api/wallets/1
```
