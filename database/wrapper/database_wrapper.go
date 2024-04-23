package databasewrapper

import "database/sql"

type DatabaseWrapper struct {
	DatabaseInterface
}

type DatabaseConnectionWrapper struct {
	dbPq *sql.DB
	DatabaseConnectionInterface
}

func (dw *DatabaseWrapper) Open(driverName, dataSourceName string) (*sql.DB, error) {
	return sql.Open(driverName, dataSourceName)
}

func (dcw *DatabaseConnectionWrapper) Ping() error {
	return dcw.dbPq.Ping()
}

func (dcw *DatabaseConnectionWrapper) Close() error {
	return dcw.dbPq.Close()
}

func (dcw *DatabaseConnectionWrapper) SetMaxIdleConns(n int) {
	dcw.dbPq.SetMaxIdleConns(n)
}

func (dcw *DatabaseConnectionWrapper) SetMaxOpenConns(n int) {
	dcw.dbPq.SetMaxOpenConns(n)
}

func (dcw *DatabaseConnectionWrapper) Begin() (*sql.Tx, error) {
	return dcw.dbPq.Begin()
}
