package entity

import (
	"time"
)

type UserSession struct {
	ID        int
	UserId    int
	Token     string
	ExpiresAt time.Time
}
