package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func randomCreateAccountParams(t *testing.T) CreateApiAccountParams {
	params := CreateApiAccountParams{
		Username:      "Alicee",
		IsActive:      true,
		ServiceType:   "ICD",
		PlanName:      "DEMO",
		CreditBalance: 100,
	}
	return params
}

func TestCreateApiAccount(t *testing.T) {
	arg := randomCreateAccountParams(t)
	acct, err := testQueries.CreateApiAccount(context.Background(), arg)
	// user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acct)

	require.Equal(t, arg.Username, acct.Username)
	require.NotZero(t, acct.CreatedAt)
}
