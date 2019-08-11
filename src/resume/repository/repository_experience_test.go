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

func Test_experienceRepo_FindByResumeID(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_experiences" WHERE ("resume_experiences"."resume_id" = $1)`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &experienceRepo{mock.DB}
		result := <-repo.FindByResumeID(10)
		assert.NotNil(t, result)
	})
}

func Test_experienceRepo_Save(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_experiences" ORDER BY "resume_experiences"."id" ASC LIMIT 1`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &experienceRepo{mock.DB}
		result := repo.Save(&domain.Experience{})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_experiences" ORDER BY "resume_experiences"."id" ASC LIMIT 1`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &experienceRepo{mock.DB}
		result := repo.Save(&domain.Experience{})
		assert.Error(t, result.Error)
	})
}

func Test_experienceRepo_Remove(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectBegin()
		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_experiences" WHERE "resume_experiences"."id" = $1`)).
			WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.Mock.ExpectCommit()

		repo := &experienceRepo{mock.DB}
		result := repo.Remove(&domain.Experience{ID: 1})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectBegin()
		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_experiences" WHERE "resume_experiences"."id" = $1`)).
			WithArgs(1).WillReturnError(fmt.Errorf("error"))
		mock.Mock.ExpectRollback()

		repo := &experienceRepo{mock.DB}
		result := repo.Remove(&domain.Experience{ID: 1})
		assert.Error(t, result.Error)
	})
}
