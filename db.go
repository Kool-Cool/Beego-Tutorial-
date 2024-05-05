package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/beegodb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM material")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var material_id int
		var material_name string
		var material_subtype string
		if err := rows.Scan(&material_id, &material_name, &material_subtype); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Material ID: %d, Material Name: %s, Material Subtype: %s\n", material_id, material_name, material_subtype)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
