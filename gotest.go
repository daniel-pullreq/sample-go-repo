package database

import (
	"database/sql"
	"log"
)

func Status(username string) string {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	sql := "SELECT * FROM user WHERE username='" + username + "'"
	row := db.QueryRow(sql)

	var isLogged bool
	row.Scan(&isLogged)
	if isLogged {
		return "online"
	}

	return "offline"
}


func (p TagPauseParams) Valid() bool {
	empty := 0
	if p.Hashes == "" {
		empty += 1
	}
	if p.AdTagIDs == "" {
		empty += 1
	}
	if p.Name == "" {
		empty += 1
	}
	return empty == 2
}
