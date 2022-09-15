package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetCampaignBy(db *sql.DB, id string) Campaign {
	rows, err := db.Query(`select "id", "from", "subject", "body", "user_id" from campaigns where id = $1`, id)
	if err != nil {
		log.Fatal("GetCampaignBy: ", err.Error())
	}
	defer rows.Close()

	campaign := Campaign{}
	for rows.Next() {
		err := rows.Scan(&campaign.Id, &campaign.From, &campaign.Subject, &campaign.Body, &campaign.CreatedById)
		if err != nil {
			log.Fatal(err)
		}
	}
	return campaign
}
