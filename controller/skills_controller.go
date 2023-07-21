package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Me struct {
	Types string
	Skill string
	About string
}

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2008"
	dbname   = "postgres"
)

func GetAllData() []Me {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, lag := db.Query("SELECT * FROM skills")
	if lag != nil {
		panic(lag)
	}

	var names []Me

	defer rows.Close()

	for rows.Next() {
		var p Me
		lag = rows.Scan(&p.Types, &p.Skill, &p.About)
		if lag != nil {
			fmt.Print(p.Types, p.Skill, p.About)
			continue
		}
		names = append(names, p)
	}

	return names
}
