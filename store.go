package main

import "database/sql"

type Store interface {
	// users
	CreateUser(u *User) (*User, error)
	GetUser(id string) (User, error)
	// tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser(u *User) (*User, error) {
	res, err := s.db.Exec(`
		INSERT INTO users (email, firstName, lastName, password)
		VALUES (?, ?, ?, ?)`, u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	u.ID = id
	return u, nil
}

func (s *Storage) GetUser(id string) (User, error) {
	var u User
	err := s.db.QueryRow(`
		SELECT id, email, firstName, lastName, createdAt 
		FROM users
		WHERE id = ?
	`, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)
	return u, err
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	res, err := s.db.Exec(`
		INSERT INTO tasks (name, status, project_id, assigned_to) 
		VALUES (?, ?, ?, ?)`, t.Name, t.Status, t.ProjectId, t.AssignedTo)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow(`
		SELECT id, name, status, project_id, assigned_to, createdAt 
		FROM tasks
		WHERE id = ?
	`, id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectId, &t.AssignedTo, &t.CreatedAt)
	return &t, err
}
