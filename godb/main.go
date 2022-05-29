package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type City struct {
	ID           int
	Name         string
	ConutryCodde string
	District     string
	Population   int
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/world")
	if err != nil {
		panic(err.Error())
	}
	cities, err := GetCities(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v", cities[:10])
}

func GetCities(db *sql.DB) ([]City, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := "select ID, Name from world.city"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	cities := []City{}
	for rows.Next() {
		city := City{}
		err = rows.Scan(
			&city.ID,
			&city.Name,
		)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}
