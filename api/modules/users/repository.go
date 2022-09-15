package users

import (
	"database/sql"
)

var queryBase string = "select \"id\", \"name\", \"email\", \"password_hash\", \"created_at\" from users"

type Repository struct {
	DB *sql.DB
}

func (r *Repository) scan(rows *sql.Rows) (User, error) {
	user := User{}
	err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *Repository) GetByEmail(email string) (User, error) {

	rows, err := r.DB.Query(queryBase+` where email = $1`, email)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		return r.scan(rows)
	}
	return User{}, nil
}

func (r *Repository) Create(user User) error {

	stmt, err := r.DB.Prepare("INSERT INTO users (id, name, email, password_hash, created_at) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Id, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Save(user User) error {
	stmt, err := r.DB.Prepare("UPDATE users set name = $1, passwordHash = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.PasswordHash, user.Id)
	if err != nil {
		return err
	}
	return nil
}
