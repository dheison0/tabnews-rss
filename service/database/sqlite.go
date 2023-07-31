package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLITEDatabase struct {
	db *sql.DB
}

func (s *SQLITEDatabase) execute(q string, args ...any) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	_, err = tx.Exec(q, args...)
	return err
}

func (s *SQLITEDatabase) Open(f string) error {
	database, err := sql.Open("sqlite3", f)
	if err != nil {
		return err
	}
	s.db = database
	return s.Setup()
}

func (s *SQLITEDatabase) Close() {
	s.db.Close()
}

func (s *SQLITEDatabase) Setup() error {
	s.db.SetMaxOpenConns(1)
	return s.execute("CREATE TABLE IF NOT EXISTS users(name TEXT UNIQUE, status TEXT);")
}

func (s *SQLITEDatabase) GetUsers() ([]User, error) {
	var users []User
	rows, err := s.db.Query("SELECT * FROM users;")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.Name, &user.Exists); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *SQLITEDatabase) AddUser(user User) bool {
	err := s.execute(`INSERT INTO users VALUES (?, ?);`, user.Name, user.Exists)
	return err == nil
}

func (s *SQLITEDatabase) RemoveUser(user User) error {
	return s.execute(`DELETE FROM users WHERE name=?;`, user.Name)
}

func (s *SQLITEDatabase) SetUserExists(user User, exists bool) error {
	return s.execute(`UPDATE users SET status=? WHERE name=?;`, exists, user.Name)
}
