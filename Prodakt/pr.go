package prodakt

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)



func Product(db *sql.DB) {
		test, err := db.Begin()
		if err != nil {
			log.Fatal("failed to begin transaction", err)
		}

		sql := `
		INSERT INTO products(product_id,product_name,category_id,unit,price) values(78,'Milliy Coca-Cola',1,'1 litr',9000);
	`
	
		res, err := test.Exec(sql)
		if err != nil {
			log.Fatal("failed to exec prodakt transaction", err)
		}
	
		sql = `
		SELECT product_name FROM products 
		WHERE product_id = 78;
	`
	
		var product string
		err = test.QueryRow(sql).Scan(&product)
		if err != nil {
			log.Fatal("failed to query transaction", err)
		}

		sql = `
		SELECT description FROM categories 
		WHERE category_id = 1;
		`
	
		var category string
		err = test.QueryRow(sql).Scan(&category)
		if err != nil {
			log.Fatal("failed to query transaction", err)
		}
	
		n, err := res.RowsAffected()
		if err != nil {
			log.Fatal("failed to get rows affected", err)
		}
	
	//Commit qilish
		err = test.Commit()
		if err != nil {
			log.Fatal("failed to commit", err)
		}
	
		fmt.Println(n,"\nProduct name: ",product,"\nCategory: ",category)
	}
	