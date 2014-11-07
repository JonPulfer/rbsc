package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20141107092742(txn *sql.Tx) {
	sql := `
	CREATE TABLE article
(
   id integer, 
   title character varying, 
   body character varying, 
   created_at timestamp with time zone, 
   updated_at timestamp with time zone, 
   deleted_at timestamp with time zone, 
   CONSTRAINT article_id_pkey PRIMARY KEY (id)
) 
WITH (
  OIDS = FALSE
);`
	_, err := txn.Exec(sql)
	if err != nil {
		log.Printf("Failed to execute %s: %v\n", sql, err)
		return
	}
}

// Down is executed when this migration is rolled back
func Down_20141107092742(txn *sql.Tx) {
	sql := `drop table article;`
	_, err := txn.Exec(sql)
	if err != nil {
		log.Printf("Failed to execute %s: %v\n", sql, err)
		return
	}
}
