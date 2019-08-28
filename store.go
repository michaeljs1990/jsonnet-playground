package main

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DatastoreNoSuchID = errors.New("The ID you request to lookup doesn't exist in our backend store")
)

type PersistJsonnet interface {
	Get(id string) (string, error)
	Store(id string, code string) error
}

// InMemory is for testing or a quick deploy
type InMemory struct {
	sync.RWMutex
	store map[string]string
}

func (i InMemory) Get(id string) (string, error) {
	i.RLock()
	defer i.RUnlock()
	val, ok := i.store[id]
	if ok {
		return val, nil
	}
	return val, DatastoreNoSuchID
}

func (i InMemory) Store(id string, code string) error {
	i.Lock()
	i.store[id] = code
	i.Unlock()
	return nil
}

// SQL backed storage engine using the database/sql functions only
type JSQL struct {
	Conn      *sql.DB
	GetStmt   *sql.Stmt
	StoreStmt *sql.Stmt
}

// Setup whatever is needed before calling Get/Store such as ensuring the
// DB is initialized and that we have prepare statements setup
func NewJSQL(conn *sql.DB) JSQL {
	m := JSQL{
		Conn: conn,
	}

	m.GetStmt, _ = conn.Prepare("FROM jsonnet SELECT code WHERE id = ? LIMIT 1")
	m.StoreStmt, _ = conn.Prepare("INSERT INTO jsonnet (id, code) VALUES (?, ?)")

	return m
}

func (m JSQL) Get(id string) (string, error) {
	row := m.GetStmt.QueryRow(id)
	var code string
	if err := row.Scan(&code); err == nil {
		return code, nil
	}
	return code, DatastoreNoSuchID
}

func (m JSQL) Store(id string, code string) error {
	row := m.StoreStmt.QueryRow(id, code)
	if err := row.Scan(); err != nil {
		return err
	}
	return nil
}
