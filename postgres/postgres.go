package postgres

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
}

type User struct {
	Name string
	ID   int
	Age  int
}

func scanUserRows(rows *sql.Rows) []User {
	var r User
	users := []User{}

	for rows.Next() {
		err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
		)
		if err != nil {
			fmt.Println("Error while scanning User rows: ", err)
		}
		users = append(users, r)
	}

	return users
}

func (d *Db) GetRows(sqlStmt string, operation string, query string) *sql.Rows {
	stmt, err := d.Prepare(sqlStmt)
	if err != nil {
		fmt.Println(operation, " Preparation Error: ", err)
	}

	var rows *sql.Rows

	switch {
	case len(query) == 0:
		rows, err = stmt.Query()
	default:
		rows, err = stmt.Query(query)
	}
	if err != nil {
		fmt.Println(operation, " Query Error: ", err)
	}

	return rows
}

func (d *Db) GetUsers() []User {
	rows := d.GetRows(
		"SELECT * FROM users",
		"GetUsers",
		"",
	)

	return scanUserRows(rows)
}

func (d *Db) GetUsersByName(name string) []User {
	rows := d.GetRows(
		"SELECT * FROM users WHERE LOWER(name)=LOWER($1)",
		"GetUsersByName",
		name,
	)

	return scanUserRows(rows)
}

func (d *Db) GetUsersByID(id int) []User {
	rows := d.GetRows(
		"SELECT * FROM users WHERE id=$1",
		"GetUsersByID", strconv.Itoa(id),
	)

	return scanUserRows(rows)
}
