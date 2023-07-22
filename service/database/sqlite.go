package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLITEDatabase struct {
	db *sql.DB
}

func (s *SQLITEDatabase) execute(q string, args ...interface{}) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare(q)
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

func (s *SQLITEDatabase) Open(f string) error {
	database, err := sql.Open("sqlite3", f)
	if err != nil {
		return err
	}
	s.db = database
	return s.Initialize()
}

func (s *SQLITEDatabase) Close() {
	s.db.Close()
}

func (s *SQLITEDatabase) Initialize() error {
	return s.execute("CREATE TABLE IF NOT EXISTS users(name TEXT UNIQUE, status TEXT);")
}

func (s *SQLITEDatabase) GetUsers() ([]User, error) {
	rows, err := s.db.Query("SELECT * FROM users;")
	if err != nil {
		return []User{}, err
	}
	var users []User
	for rows.Next() {
		var name string
		var exists int
		if err := rows.Scan(&name, &exists); err != nil {
			continue
		}
		users = append(users, User{name, exists == 1})
	}
	return users, nil
}

func (s *SQLITEDatabase) AddUser(user User) bool {
	var userExists int
	if user.Exists {
		userExists = 1
	} else {
		userExists = 0
	}
	err := s.execute(`INSERT INTO users VALUES (?, ?);`, user.Name, userExists)
	return err == nil
}

func (s *SQLITEDatabase) RemoveUser(user User) error {
	return s.execute(`DELETE FROM users WHERE name=?;`, user.Name)
}
