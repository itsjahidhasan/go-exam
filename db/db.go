package db

import "database/sql"

func Open(host, port, user, pass, name string) (*sql.DB, error) {
	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + name + " sslmode=disable"
	return sql.Open("postgres", dsn)
}
