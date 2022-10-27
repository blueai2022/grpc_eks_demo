package graphql

import (
	"time"
)

type Account struct {
	ID            string    `json:"id"`
	IsActive      bool      `json:"is_active"`
	IsAutoRenewal bool      `json:"is_auto_renewal"`
	ServiceType   string    `json:"service_type"`
	PlanName      string    `json:"plan_name"`
	CreditBalance int64     `json:"credit_balance"`
	ExpiresAt     time.Time `json:"expires_at"`
	CreatedAt     time.Time `json:"created_at"`
	Owner         *User     `json:"owner"`
}

type Session struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
