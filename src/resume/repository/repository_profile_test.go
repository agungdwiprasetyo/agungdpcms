package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/agungdwiprasetyo/agungdpcms/shared/mocking"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_profileRepo_FindByResumeID(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_profiles" WHERE ("resume_profiles"."resume_id" = $1)`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &profileRepo{mock.DB}
		result := <-repo.FindByResumeID(10)
		assert.NotNil(t, result)
	})
}

func Test_profileRepo_Save(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_profiles" ORDER BY "resume_profiles"."id" ASC LIMIT 1`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &profileRepo{mock.DB}
		result := repo.Save(&domain.Profile{})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_profiles" ORDER BY "resume_profiles"."id" ASC LIMIT 1`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &profileRepo{mock.DB}
		result := repo.Save(&domain.Profile{})
		assert.Error(t, result.Error)
	})
}

func Test_profileRepo_Remove(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectBegin()
		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_profiles" WHERE "resume_profiles"."id" = $1`)).
			WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.Mock.ExpectCommit()

		repo := &profileRepo{mock.DB}
		result := repo.Remove(&domain.Profile{ID: 1})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectBegin()
		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_profiles" WHERE "resume_profiles"."id" = $1`)).
			WithArgs(1).WillReturnError(fmt.Errorf("error"))
		mock.Mock.ExpectRollback()

		repo := &profileRepo{mock.DB}
		result := repo.Remove(&domain.Profile{ID: 1})
		assert.Error(t, result.Error)
	})
}
