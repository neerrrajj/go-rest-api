package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage (cfg mysql.Config) *MySQLStorage {
	log.Println("connecting to MySQL...")
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println("connection FAILED")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println("ping FAILED")
		log.Fatal(err)
	}
	log.Println("CONNECTED to MySQL!")
	return &MySQLStorage{db: db}
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	err := s.createProjectsTable()
	if err != nil {
		return nil, err
	}
	err = s.createUsersTable()
	if err != nil {
		return nil, err
	}
	err = s.createTasksTable()
	if err != nil {
		return nil, err
	}
	return s.db, nil
}

// DOUBT : how do we know we can use .Exec() on a sql.db ?

func (s *MySQLStorage) createProjectsTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}

func (s *MySQLStorage) createUsersTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			firstName VARCHAR(255) NOT NULL,
			lastName VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}

func (s *MySQLStorage) createTasksTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			projectId INT UNSIGNED NOT NULL,
			assignedTo INT UNSIGNED NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (assignedTo) REFERENCES users(id),
			FOREIGN KEY (projectId) REFERENCES projects(id)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}