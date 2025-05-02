# Wallet API (Golang)

A simple RESTful API to create wallets, check balances, deposit, and withdraw money.

## 🚀 Tech Stack

- Go
- Gin (HTTP framework)
- GORM (ORM)
- SQLite (Database)

## ⚙️ Setup

```bash
git clone https://github.com/YOUR_USERNAME/wallet-api-go
cd wallet-api-go
go run .
```

## 🔍 API Endpoints

| Method | Endpoint                     | Description           |
|--------|------------------------------|-----------------------|
| POST   | /api/wallets                 | Create a wallet       |
| GET    | /api/wallets/:id             | Get wallet balance    |
| POST   | /api/wallets/:id/deposit     | Deposit money         |
| POST   | /api/wallets/:id/withdraw    | Withdraw money        |

## 🧪 Testing

You can test the endpoints using curl or Postman:

```bash
curl -X POST http://localhost:8080/api/wallets
curl -X POST http://localhost:8080/api/wallets/1/deposit -H "Content-Type: application/json" -d '{"amount": 50}'
curl -X POST http://localhost:8080/api/wallets/1/withdraw -H "Content-Type: application/json" -d '{"amount": 20}'
curl http://localhost:8080/api/wallets/1
```

