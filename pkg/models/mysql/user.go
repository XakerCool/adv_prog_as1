package mysql

import (
	"1_assignment/pkg/models"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Register(fullName, email, role, password string, approved bool) (int, error) {
	stmt := `INSERT INTO users (full_name, email, role, password, approved) VALUES (?, ?, ?, ?, ?)`

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	result, err := m.DB.Exec(stmt, fullName, email, role, hashedPwd, approved)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err

	}
	return int(id), nil
}

func (m *UserModel) Login(email, password string) (*models.User, error) {
	stmt := `SELECT id, full_name, email, password, role, approved FROM users WHERE email = ?`

	row := m.DB.QueryRow(stmt, email)
	user := &models.User{}

	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Approved)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, models.ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (m *UserModel) GetTeachers() ([]*models.User, error) {
	stmt := `SELECT id, full_name, email, role, approved FROM users WHERE role="teacher"`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	users := []*models.User{}
	for rows.Next() {
		u := &models.User{}
		err := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.Role, &u.Approved)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
