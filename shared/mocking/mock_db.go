package mocking

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// MockDB model
type MockDB struct {
	db   *sql.DB
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

// New construct db mocking
func New() *MockDB {
	db, mock, err := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic(err)
	}
	return &MockDB{
		db: db, DB: gormDB, Mock: mock,
	}
}

// Close open db
func (m *MockDB) Close() {
	m.db.Close()
	m.DB.Close()
}
