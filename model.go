package main

import "gorm.io/gorm"

type Wallet struct {
    gorm.Model
    Balance float64 `gorm:"not null;default:0"`
}

