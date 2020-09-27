package database

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/polyhedronmc/daemon/config"
)

// Database Wrapper around the PostgreSQL database struct.
type Database struct {
	con *pg.DB
}

var db Database;

// Connect Connect to PostgreSQL.
func Connect(config config.DatabaseConfig) Database {
	postgres := pg.Connect(&pg.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})

	var n int
	_, err := postgres.QueryOne(pg.Scan(&n), "SELECT 1")
	if (err != nil) {
		panic(err)
	}
	
	db = Database {
		con: postgres,
	}

	return db
}
