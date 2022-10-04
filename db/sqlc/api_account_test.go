package db

import (
	"context"
	"testing"

	"github.com/blueai2022/appsubmission/crypt"
	"github.com/stretchr/testify/require"
)

func randomCreateAccountParams(t *testing.T, username string) CreateApiAccountParams {
	params := CreateApiAccountParams{
		Username:      username,
		IsActive:      true,
		ServiceType:   "ICD",
		PlanName:      "DEMO",
		CreditBalance: 50,
	}
	return params
}

func TestCreateApiAccount(t *testing.T) {
	arg := randomCreateUserParams(t)
	// fixed password for api account testing in Postman
	hashedPassword, err := crypt.HashPassword("secret")
	require.NoError(t, err)

	arg.HashedPassword = hashedPassword
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	argAcct := randomCreateAccountParams(t, arg.Username)
	acct, err := testQueries.CreateApiAccount(context.Background(), argAcct)
	// user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acct)

	require.Equal(t, arg.Username, acct.Username)
	require.NotZero(t, acct.CreatedAt)
}
