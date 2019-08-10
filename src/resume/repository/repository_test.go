package repository

import (
	"fmt"
	"testing"

	"github.com/agungdwiprasetyo/agungdpcms/shared/mocking"
	"github.com/stretchr/testify/assert"
)

func TestRepository_WithTransaction(t *testing.T) {
	t.Run("Test Start Transaction (Panic recovered)", func(t *testing.T) {
		mock := mocking.New()
		mock.Mock.ExpectBegin()
		defer mock.Close()

		repo := NewRepository(mock.DB)
		err := repo.WithTransaction(func(repo *Repository) error {
			panic("Error")
		})
		mock.Mock.ExpectRollback()
		assert.Error(t, err)
	})
	t.Run("Test error happened when transaction", func(t *testing.T) {
		mock := mocking.New()
		mock.Mock.ExpectBegin()
		defer mock.Close()

		repo := NewRepository(mock.DB)
		mock.Mock.ExpectRollback()
		err := repo.WithTransaction(func(repo *Repository) error {
			return fmt.Errorf("error")
		})
		assert.Error(t, err)
	})
	t.Run("Test positive transaction", func(t *testing.T) {
		mock := mocking.New()
		mock.Mock.ExpectBegin()
		defer mock.Close()

		repo := NewRepository(mock.DB)
		mock.Mock.ExpectCommit()
		err := repo.WithTransaction(func(repo *Repository) error {
			return nil
		})
		assert.NoError(t, err)
	})
}
