package db

import (
	"context"
	"testing"

	"github.com/blueai2022/appsubmission/secure"
	"github.com/blueai2022/appsubmission/test"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := secure.HashPassword(test.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       test.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       test.RandomOwner(),
		Email:          test.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// func createRandomAccount(t *testing.T) Account {
// 	user := createRandomUser(t)

// 	arg := CreateAccountParams{
// 		Owner:   user.Username,
// 		Balance: test.RandomMoney(),
// 	}

// 	account, err := testQueries.CreateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account)

// 	require.Equal(t, arg.Owner, account.Owner)
// 	require.Equal(t, arg.Balance, account.Balance)
// 	require.Equal(t, arg.Currency, account.Currency)

// 	require.NotZero(t, account.ID)
// 	require.NotZero(t, account.CreatedAt)

// 	return account
// }

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// func TestGetAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, account1.Balance, account2.Balance)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestUpdateAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)

// 	arg := UpdateAccountParams{
// 		ID:      account1.ID,
// 		Balance: test.RandomMoney(),
// 	}

// 	account2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, arg.Balance, account2.Balance)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	err := testQueries.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, account2)
// }

// func TestListAccounts(t *testing.T) {
// 	var lastAccount Account
// 	for i := 0; i < 10; i++ {
// 		lastAccount = createRandomAccount(t)
// 	}

// 	arg := ListAccountsParams{
// 		Owner:  lastAccount.Owner,
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	accounts, err := testQueries.ListAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, accounts)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 		require.Equal(t, lastAccount.Owner, account.Owner)
// 	}
// }
