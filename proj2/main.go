package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/libp2p/go-libp2p"
)

func main() {
	db,  err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306/store")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var(ccnum,  date, cvv, exp string
		 amount float32
		)
	rows, err := db.Query("SELECT ccnum, date, amount, cvv, exp FROM transactions")
	if err != nil {
		log.Panic(err)
	}	
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ccnum, &date, &amount, &cvv, &exp)
		if err != nil{
			log.Panic(err)
		}
		fmt.Println(ccnum, date, amount, cvv, exp)
	}
	if rows.Err() != nil {
		log.Panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	go func ()  {
		
	}()
	
}