package migration

import (
  "database/sql"
  "fmt"

  "github.com/pressly/goose"
)

func init() {
  goose.AddMigration(Up_20170301083810, Down_20170301083810)
}

func Up_20170301083810(tx *sql.Tx) error {
  res, err := tx.Exec(`CREATE TABLE MyGuests (
           id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
           firstname VARCHAR(30) NOT NULL,
           lastname VARCHAR(30) NOT NULL,
           email VARCHAR(50),
           reg_date TIMESTAMP
       )`)
  fmt.Println(res)
  return err
}

func Down_20170301083810(tx *sql.Tx) error {
  res, err := txn.Exec("DROP TABLE MyGuests;")
  fmt.Println(res)
  return err
}
