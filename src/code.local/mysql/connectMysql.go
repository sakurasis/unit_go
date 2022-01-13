package main

import "database/sql"

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	open, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {

		}
	}(open)
}
