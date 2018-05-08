package gen

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func sqlOpen(dbName, dbUser, dbPass string) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
}

func sqlExec(conn *sql.DB, q string) error {
	log.Printf("Will execute query:\n%s\n", q)
	_, err := conn.Exec(q)
	return err
}

// Convert mysql date (string) into yy, mm and dd.
// The date input is assumed to be in right format. Will panic if error.
// func Date(s string) (int, int, int) {
// 	ss := strings.Split(s, "-")
// 	return atoi(ss[0]), atoi(ss[1]), atoi(ss[2])
// }
//
// func atoi(s string) int {
// 	v, err := strconv.Atoi(s)
// 	panicIf(err)
// 	return v
// }
