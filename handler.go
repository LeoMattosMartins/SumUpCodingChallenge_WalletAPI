package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type AmountRequest struct {
    Amount float64 `json:"amount" validate:"required,gt=0"`
}

type WalletResponse struct {
    ID      uint    `json:"id"`
    Balance float64 `json:"balance"`
}

func CreateWallet(c *gin.Context) {
    w := Wallet{}
    db.Create(&w)
    c.JSON(http.StatusCreated, WalletResponse{w.ID, w.Balance})
}

func GetWallet(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var w Wallet
    if err := db.First(&w, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
        return
    }
    c.JSON(http.StatusOK, WalletResponse{w.ID, w.Balance})
}

func Deposit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req AmountRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be > 0"})
        return
    }

    var w Wallet
    if err := db.First(&w, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
        return
    }

    w.Balance += req.Amount
    db.Save(&w)
    c.JSON(http.StatusOK, WalletResponse{w.ID, w.Balance})
}

func Withdraw(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req AmountRequest
    if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be > 0"})
        return
    }

    var w Wallet
    if err := db.First(&w, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "wallet not found"})
        return
    }

    if w.Balance < req.Amount {
        c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient funds"})
        return
    }

    w.Balance -= req.Amount
    db.Save(&w)
    c.JSON(http.StatusOK, WalletResponse{w.ID, w.Balance})
}

