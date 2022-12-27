package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func NewDb(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error ", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table ", err)
	}

	log.Infoln("okay")
}
