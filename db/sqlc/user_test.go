package db

import (
	"context"
	"go-gprc-project/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	createdUser := CreateRandomUser(t)
	fetchedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedUser)
	require.Equal(t, createdUser.Username, fetchedUser.Username)
	require.Equal(t, createdUser.HashedPassword, fetchedUser.HashedPassword)
	require.Equal(t, createdUser.FullName, fetchedUser.FullName)
	require.Equal(t, createdUser.Email, fetchedUser.Email)

	require.WithinDuration(t, createdUser.PasswordChangedAt, fetchedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, createdUser.CreatedAt, fetchedUser.CreatedAt, time.Second)
}
