package campaigns

import (
	"database/sql"
)

var queryBase string = `select id, name, "from", subject, body, status, created_at, user_id from campaigns`

type Repository struct {
	DB *sql.DB
}

func (r *Repository) scan(rows *sql.Rows) (Campaign, error) {
	campaign := Campaign{}
	err := rows.Scan(&campaign.Id, &campaign.Name, &campaign.From, &campaign.Subject, &campaign.Body, &campaign.Status, &campaign.CreatedAt, &campaign.UserId)
	if err != nil {
		return Campaign{}, err
	}
	return campaign, nil
}

func (r *Repository) GetBy(id string) (Campaign, error) {
	rows, err := r.DB.Query(queryBase+` where id = $1`, id)
	if err != nil {
		return Campaign{}, err
	}
	defer rows.Close()
	for rows.Next() {
		return r.scan(rows)
	}
	return Campaign{}, nil
}

func (r *Repository) GetEmailsReport(id string) (EmailsReport, error) {
	rows, err := r.DB.Query(`select status from subscribers where campaign_id = $1`, id)
	if err != nil {
		return EmailsReport{}, err
	}
	defer rows.Close()
	emailsReport := EmailsReport{}
	for rows.Next() {
		var status string
		err := rows.Scan(&status)
		if err != nil {
			return EmailsReport{}, err
		}
		emailsReport.BaseOfSubscribers++
		if status == "Delivery" {
			emailsReport.Sent++
		} else if status == "Open" {
			emailsReport.Opened++
		} else if status == "Bounce" {
			emailsReport.NotSent++
		}
	}
	return emailsReport, nil
}

func (r *Repository) ListBy(userId string) ([]Campaign, error) {

	rows, err := r.DB.Query(queryBase+` where user_id = $1`, userId)
	if err != nil {
		return []Campaign{}, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		campaign, err := r.scan(rows)
		if err != nil {
			return campaigns, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

func (r *Repository) Create(campaign Campaign) error {

	stmt, err := r.DB.Prepare(`
		INSERT INTO campaigns (id, name, "from", subject, body, status, created_at, user_id) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(campaign.Id, campaign.Name, campaign.From, campaign.Subject, campaign.Body, campaign.Status, campaign.CreatedAt, campaign.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Save(campaign Campaign) error {

	stmt, err := r.DB.Prepare(`UPDATE campaigns SET name = $1, "from" = $2, subject = $3, body = $4, status = $5  WHERE id = $6`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(campaign.Name, campaign.From, campaign.Subject, campaign.Body, campaign.Status, campaign.Id)
	if err != nil {
		return err
	}
	return nil
}
