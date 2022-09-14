package contacts

import (
	"database/sql"
)

var queryBase string = "select \"id\", \"name\", \"email\", \"active\", \"created_at\", \"user_id\" from users"

type Repository struct {
	DB *sql.DB
}

func (r *Repository) scan(rows *sql.Rows) (Contact, error) {
	contact := Contact{}
	err := rows.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Active, &contact.CreatedAt, &contact.UserId)
	if err != nil {
		return Contact{}, err
	}
	return contact, nil
}

func (r *Repository) GetBy(id string) (Contact, error) {
	rows, err := r.DB.Query(queryBase+` where id = $1`, id)
	if err != nil {
		return Contact{}, err
	}
	defer rows.Close()

	rows.Next()
	return r.scan(rows)
}

func (r *Repository) ListBy(userId string) ([]Contact, error) {

	contacts := make([]Contact, 0)
	rows, err := r.DB.Query(queryBase+` where user_id = $1`, userId)
	if err != nil {
		return contacts, err
	}
	defer rows.Close()

	for rows.Next() {
		contact, err := r.scan(rows)
		if err != nil {
			return contacts, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (r *Repository) Create(contact Contact) error {

	stmt, err := r.DB.Prepare("INSERT INTO contacts (id, name, email, active, created_at, user_id) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contact.Id, contact.Name, contact.Email, contact.Active, contact.CreatedAt, contact.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Save(contact Contact) error {

	stmt, err := r.DB.Prepare("UPDATE contacts name = $1, active = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contact.Name, contact.Active, contact.Id)
	if err != nil {
		return err
	}
	return nil
}
