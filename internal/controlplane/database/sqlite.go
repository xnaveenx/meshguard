package database

import (
	"log"
	"database/sql"
	_ "modernc.org/sqlite"
	)

type Database struct {
	db *sql.DB
}

func InitDB(filepath string) (*Database, error) {
	conn, err := sql.Open("sqlite", filepath)
	if err != nil{
		return nil, err
	}
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS devices (
		machine_key TEXT PRIMARY KEY,
		hostname TEXT NOT NULL,
		os TEXT
	);`

	_, err = conn.Exec(createTableQuery)
	if err != nil{
		return nil, err
	}
	log.Println("Database initialized successfully!")
	return &Database{db: conn}, nil
}

func (d *Database) SaveDevice(machinekey, hostname, os string) error{
	query := `
	INSERT INTO devices (machine_key, hostname, os)
	VALUES (?, ?, ?)
	ON CONFLICT(machine_key) DO UPDATE SET
		hostname=excluded.hostname,
		os=excluded.os;
	`

	_, err:= d.db.Exec(query, machinekey, hostname, os)
	if err!= nil{
		return err
	}
	return nil
}
