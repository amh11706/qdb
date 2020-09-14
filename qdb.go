package qdb

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var lock = sync.Mutex{}

func init() {
	lock.Lock()
}

func Set(db *sqlx.DB) {
	DB = db
	lock.Unlock()
}

func Get() *sqlx.DB {
	if DB != nil {
		return DB
	}
	lock.Lock()
	defer lock.Unlock()
	return DB
}
