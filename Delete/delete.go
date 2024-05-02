package delete

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)



func DeleteProduct(db *sql.DB){
	p, err := db.Begin()
	if err != nil {
		log.Fatal("failed to begin transaction", err)
	}

	sql := `
	DELETE  FROM products WHERE product_id = 78;
	`
	res, err := p.Exec(sql)
	if err != nil {
		log.Fatal("failed to exec delete transaction", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal("failed to get rows affected", err)
	}

	err = p.Commit()
	if err != nil {
		log.Fatal("failed to commit", err)
	}

	fmt.Println(n,"\n\nDATA DELETION COMPLETED!!!")

}