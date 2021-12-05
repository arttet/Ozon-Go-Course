package product_service

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setup_repo(t *testing.T) (*repository, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB,"sqlmock")


	repo := &repository{
		DB: sqlxDB,
	}

	return repo, mock
}

func TestDeleteProduct_Success(t *testing.T) {
	r, dbMock := setup_repo(t)
	ctx := context.Background()

	dbMock.ExpectExec("DELETE FROM products").WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 2))

	err := r.DeleteProduct(ctx, []int64{1, 2})

	require.NoError(t, err)
}