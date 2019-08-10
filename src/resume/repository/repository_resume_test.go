package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/agungdwiprasetyo/agungdpcms/shared/filter"
	"github.com/agungdwiprasetyo/agungdpcms/shared/mocking"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_resumeRepo_FindAll(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" ORDER BY LIMIT 0 OFFSET 0`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "resumes"`)).
			WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(10))

		repo := &resumeRepo{mock.DB}
		result := repo.FindAll(&filter.Filter{})
		assert.NoError(t, result.Error)

		// counting
		c := repo.Count(&domain.Resume{})
		assert.Equal(t, 10, c)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" ORDER BY LIMIT 0 OFFSET 0`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &resumeRepo{mock.DB}
		result := repo.FindAll(&filter.Filter{})
		assert.Error(t, result.Error)
	})
}

func Test_resumeRepo_FindBySlug(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" WHERE ("resumes"."slug" = $1)`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &resumeRepo{mock.DB}
		result := repo.FindBySlug("agungdp")
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" WHERE ("resumes"."slug" = $1)`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &resumeRepo{mock.DB}
		result := repo.FindBySlug("agungdp")
		assert.Error(t, result.Error)
	})
}

func Test_resumeRepo_Save(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" ORDER BY "resumes"."id" ASC LIMIT 1`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &resumeRepo{mock.DB}
		result := repo.Save(&domain.Resume{})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resumes" ORDER BY "resumes"."id" ASC LIMIT 1`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &resumeRepo{mock.DB}
		result := repo.Save(&domain.Resume{})
		assert.Error(t, result.Error)
	})
}
