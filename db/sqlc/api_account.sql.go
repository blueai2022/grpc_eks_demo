// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: api_account.sql

package db

import (
	"context"
	"time"
)

const createApiAccount = `-- name: CreateApiAccount :one
INSERT INTO api_accounts (
  username,
  is_active,
  service_type,
  plan_name,
  credit_balance,
  last_use_at,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING id, username, is_active, is_auto_renewal, service_type, plan_name, credit_balance, active_at, last_use_at, expires_at, created_at
`

type CreateApiAccountParams struct {
	Username      string    `json:"username"`
	IsActive      bool      `json:"is_active"`
	ServiceType   string    `json:"service_type"`
	PlanName      string    `json:"plan_name"`
	CreditBalance int64     `json:"credit_balance"`
	LastUseAt     time.Time `json:"last_use_at"`
	ExpiresAt     time.Time `json:"expires_at"`
}

func (q *Queries) CreateApiAccount(ctx context.Context, arg CreateApiAccountParams) (ApiAccount, error) {
	row := q.db.QueryRowContext(ctx, createApiAccount,
		arg.Username,
		arg.IsActive,
		arg.ServiceType,
		arg.PlanName,
		arg.CreditBalance,
		arg.LastUseAt,
		arg.ExpiresAt,
	)
	var i ApiAccount
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.IsActive,
		&i.IsAutoRenewal,
		&i.ServiceType,
		&i.PlanName,
		&i.CreditBalance,
		&i.ActiveAt,
		&i.LastUseAt,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const debitApiAccountBalance = `-- name: DebitApiAccountBalance :one
UPDATE api_accounts
SET credit_balance = credit_balance - 1
WHERE id = $1
RETURNING id, username, is_active, is_auto_renewal, service_type, plan_name, credit_balance, active_at, last_use_at, expires_at, created_at
`

func (q *Queries) DebitApiAccountBalance(ctx context.Context, id int64) (ApiAccount, error) {
	row := q.db.QueryRowContext(ctx, debitApiAccountBalance, id)
	var i ApiAccount
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.IsActive,
		&i.IsAutoRenewal,
		&i.ServiceType,
		&i.PlanName,
		&i.CreditBalance,
		&i.ActiveAt,
		&i.LastUseAt,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getActiveApiAccount = `-- name: GetActiveApiAccount :one
SELECT id, username, is_active, is_auto_renewal, service_type, plan_name, credit_balance, active_at, last_use_at, expires_at, created_at FROM api_accounts
WHERE 
    username = $1 AND
    is_active = true AND
    service_type = $2
LIMIT 1
`

type GetActiveApiAccountParams struct {
	Username    string `json:"username"`
	ServiceType string `json:"service_type"`
}

func (q *Queries) GetActiveApiAccount(ctx context.Context, arg GetActiveApiAccountParams) (ApiAccount, error) {
	row := q.db.QueryRowContext(ctx, getActiveApiAccount, arg.Username, arg.ServiceType)
	var i ApiAccount
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.IsActive,
		&i.IsAutoRenewal,
		&i.ServiceType,
		&i.PlanName,
		&i.CreditBalance,
		&i.ActiveAt,
		&i.LastUseAt,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getApiAccount = `-- name: GetApiAccount :one
SELECT id, username, is_active, is_auto_renewal, service_type, plan_name, credit_balance, active_at, last_use_at, expires_at, created_at FROM api_accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetApiAccount(ctx context.Context, id int64) (ApiAccount, error) {
	row := q.db.QueryRowContext(ctx, getApiAccount, id)
	var i ApiAccount
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.IsActive,
		&i.IsAutoRenewal,
		&i.ServiceType,
		&i.PlanName,
		&i.CreditBalance,
		&i.ActiveAt,
		&i.LastUseAt,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}