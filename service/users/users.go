package users

import (
	"database/sql"
	"fmt"
	"server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserById(id int) (*types.User, error) {

	rows, err := s.db.Query("select id, login, password from users where id=$1", id)
	if err != nil {
		return nil, err
	}
	var u *types.User
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	return u, nil
}

func (s *Store) GetUserBylogin(login string) (*types.User, error) {

	rows, err := s.db.Query("select id, login, password from users where login=$1", login)
	if err != nil {
		return nil, err
	}

	var u *types.User
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u == nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(&u.ID, &u.Login, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) CreateUser(u types.User) error {

	_, err := s.db.Exec("insert into users (login,password) values ($1,$2)", u.Login, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateUser(u types.User) error {

	_, err := s.db.Exec("update users set (login,password) = ($1,$2) where id=$3", u.Login, u.Password, u.ID)
	if err != nil {
		return err
	}
	return nil

}
