package dbrepo

import "database/sql"

type PostgresDBRepo struct {
	DB *sql.DB
}

// Connection returns the database connection, which is used to perform database operations
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}