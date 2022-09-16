package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func SaveSubscriber(db *sql.DB, subscriber Subscriber) {
	stmt, err := db.Prepare("INSERT INTO subscribers VALUES($1, Now(), $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(subscriber.Id, subscriber.CampaignID, subscriber.ContactID, subscriber.Email, subscriber.Status, subscriber.ProviderEmailKey)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateStatusSubscriber(db *sql.DB, newStatus, providerEmailKey string) (bool, error) {
	var status string
	err := db.QueryRow(`select status from subscribers where provider_email_key = $1`, providerEmailKey).Scan(&status)
	if err != nil {
		log.Fatal("GetCampaignBy: ", err.Error())
	}

	if newStatus == "Sent" && (status == "Delivery" || status == "Bounce") {
		return true, nil
	} else if newStatus == "Delivery" && status == "Open" {
		return true, nil
	}

	stmt, err := db.Prepare("UPDATE subscribers set status = $1 WHERE provider_email_key = $2")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(newStatus, providerEmailKey)
	if err != nil {
		return false, err
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected > 0, err
}
