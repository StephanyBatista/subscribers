package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetContactsBy(db *sql.DB, userId string) []Contact {
	rows, err := db.Query(`select "id", "name", "email" from contacts where user_id = $1`, userId)
	if err != nil {
		log.Fatal("GetContactsBy: ", err.Error())
	}
	defer rows.Close()

	contacts := make([]Contact, 0)
	for rows.Next() {
		contact := Contact{}
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Email)
		if err != nil {
			log.Fatal(err)
		}
		contacts = append(contacts, contact)
	}
	return contacts
}
