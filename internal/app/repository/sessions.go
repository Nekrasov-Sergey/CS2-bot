package repository

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

func AddSession(db *sqlx.DB, VkID int, output string) error {
	_, err := db.Exec("INSERT INTO sessions(user_vk_id, sessions_date, solutions) VALUES($1, $2, $3)", VkID, time.Now().UTC().Add(time.Hour*3), output)
	if err != nil {
		log.WithError(err).Error("Failed to add session")
		return err
	}
	return nil
}
