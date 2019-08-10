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

func Test_achievementRepo_FindByResumeID(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_achievements" WHERE ("resume_achievements"."resume_id" = $1)`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &achievementRepo{mock.DB}
		result := <-repo.FindByResumeID(10)
		assert.NotNil(t, result)
	})
}

func Test_achievementRepo_Save(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_achievements" ORDER BY "resume_achievements`)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "resume_id"}).AddRow(1, 10))

		repo := &achievementRepo{mock.DB}
		result := repo.Save(&domain.Achievement{})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "resume_achievements" ORDER BY "resume_achievements`)).
			WillReturnError(fmt.Errorf("error"))

		repo := &achievementRepo{mock.DB}
		result := repo.Save(&domain.Achievement{})
		assert.Error(t, result.Error)
	})
}

func Test_achievementRepo_Remove(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_achievements" WHERE "resume_achievements"."id" = $1`)).
			WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := &achievementRepo{mock.DB}
		result := repo.Remove(&domain.Achievement{ID: 1})
		assert.NoError(t, result.Error)
	})
	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock := mocking.New()
		defer mock.Close()

		mock.Mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "resume_achievements" WHERE "resume_achievements"."id" = $1`)).
			WithArgs(1).WillReturnError(fmt.Errorf("error"))

		repo := &achievementRepo{mock.DB}
		result := repo.Remove(&domain.Achievement{ID: 1})
		assert.Error(t, result.Error)
	})
}
