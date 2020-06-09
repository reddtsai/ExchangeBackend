package model

import (
	"database/sql"
)

type mysqlDB struct {
	db *sql.DB
}

func newMySQL() (*mysqlDB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(0.0.0.0:3306)/EXCHANGE?tls=skip-verify&autocommit=true")
	if err != nil {
		return nil, err
	}
	// db.SetMaxIdleConns()

	return &mysqlDB{
		db: db,
	}, nil
}

func (m *mysqlDB) insert(sql string) error {
	_, err := m.db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlDB) close() {
	m.db.Close()
}
