package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/go-playground/validator/v10"
)

var (
    db       *gorm.DB
    validate *validator.Validate
)

func initDatabase() {
    var err error
    db, err = gorm.Open(sqlite.Open("wallet.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database")
    }
    db.AutoMigrate(&Wallet{})
}

func main() {
    initDatabase()
    validate = validator.New()

    r := gin.Default()

    api := r.Group("/api/wallets")
    {
        api.POST("", CreateWallet)
        api.GET("/:id", GetWallet)
        api.POST("/:id/deposit", Deposit)
        api.POST("/:id/withdraw", Withdraw)
    }

    r.Run(":8080")
}

