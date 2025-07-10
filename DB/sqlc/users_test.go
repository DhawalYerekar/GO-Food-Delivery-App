package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		// Name:     "Test User",
		// Age:      25,
		// Email:    "testuser@example.com",
		// Password: "hashedpassword123", // typically a hash of a real password
		// Address:  "123 Test Street",
		// Contact:  "9876543210",
		// Gender:   "Other",
		Name:     "Test User",
		Age:      sql.NullInt32{Int32: 25, Valid: true},
		Email:    "testuser3@example.com",
		Password: "hashed-password",
		Address:  "123 Test Street",
		Contact:  "9876543210",
		Gender:   sql.NullString{String: "Male", Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	account1 := createRandomUser(t)

	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Age, account2.Age)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.Address, account2.Address)
	require.Equal(t, account1.Contact, account2.Contact)
	require.Equal(t, account1.Gender, account2.Gender)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateUser(t *testing.T) {
	account1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:      account1.ID,
		Name:    account1.Name,
		Address: account1.Address,
		Contact: account1.Contact,
	}

	account2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Name, account2.Name)
	require.Equal(t, account1.Address, account2.Address)
	require.Equal(t, account1.Contact, account2.Contact)

}
