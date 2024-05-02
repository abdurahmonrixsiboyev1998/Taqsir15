package price

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)




func NewPrice(db *sql.DB){
		new, err := db.Begin()
		if err != nil {
			log.Fatal("failed to begin transaction", err)
		}
		sql := `
		UPDATE products
		SET price = 11000
		WHERE product_id = 78;
		`

		res, err := new.Exec(sql)
		if err != nil {
			log.Fatal("failed to exec price transaction", err)
		}

		sql = `
		SELECT price FROM products 
		WHERE product_id = 78;
		`
	
		var price string
		err = new.QueryRow(sql).Scan(&price)
		if err != nil {
			log.Fatal("failed to query transaction", err)
		}
	
		sql = `
		SELECT description FROM categories 
		WHERE category_id = 1;
		`
	
		var category string
		err = new.QueryRow(sql).Scan(&category)
		if err != nil {
			log.Fatal("failed to query transaction", err)
		}
	
	
		x, err := res.RowsAffected()
		if err != nil {
			log.Fatal("failed to get rows affected", err)
		}
	
	
		err = new.Commit()
		if err != nil {
			log.Fatal("failed to commit", err)
		}
	
		fmt.Println(x,"\n\nNew price: ", price,"\nCategory: ",category)
	
	}