package db

import (
	"context"
	"github.com/arbeeorlar/simplebank/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func creatAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: pgtype.Text{Valid: true, String: util.RandomCurrency()},
	}

	account, err := testQuery.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	creatAccount(t)
}
func TestGetAccount(t *testing.T) {
	account1 := creatAccount(t)
	account2, err := testQuery.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)

	require.Equal(t, account1.ID, account2.ID)
	require.NotEmpty(t, account1.ID)
	require.NotZero(t, account1.ID)
	require.NotZero(t, account2.CreatedAt)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)

	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := creatAccount(t)
	arg := UpdateAccountParams{
		Owner: util.RandomOwner(),
		ID:    account1.ID,
	}
	account2, err := testQuery.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEqual(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
	require.Equal(t, account1.ID, account2.ID)

}

func TestDeleteAccount(t *testing.T) {
	account1 := creatAccount(t)
	err := testQuery.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQuery.GetAccount(context.Background(), account1.ID)

	require.Empty(t, account2)
	require.Error(t, err)

}

//func TestDeleteAccounts(t *testing.T) {
//
//}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		creatAccount(t)
	}

	arg := ListAccountsParams{
		Offset: 5,
		Limit:  5,
	}
	accounts, err := testQuery.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
