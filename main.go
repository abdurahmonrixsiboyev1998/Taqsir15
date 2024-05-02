package main

import (
	delete "Taqsir15/Delete"
	new "Taqsir15/Price"
	prodakt "Taqsir15/Prodakt"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "14022014"
	dbname   = "demo"
)

func main() {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	z, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer z.Close()

	err = z.Ping()
	if err != nil {
		log.Fatal("failed to ping", err)
	}

	fmt.Println("successfully connected")
	prodakt.Product(z)
	new.NewPrice(z)
	delete.DeleteProduct(z)

}
