package mocking

import (
	"database/sql"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var lock sync.Mutex

// MockDB model
type MockDB struct {
	db   *sql.DB
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

// New construct db mocking
func New() *MockDB {
	lock.Lock()
	defer lock.Unlock()
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
	lock.Lock()
	defer lock.Unlock()
	m.db.Close()
	m.DB.Close()
}
