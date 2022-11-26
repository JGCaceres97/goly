package model

import "time"

type Goly struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Redirect  string    `json:"redirect" gorm:"not null"`
	Goly      string    `json:"goly" gorm:"unique;not null"`
	Clicked   uint64    `json:"clicked"`
	IsRandom  bool      `json:"isRandom"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
}
