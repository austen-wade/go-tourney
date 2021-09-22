package postgres

import (
	"database/sql"
	"fmt"

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
	ID   int
	Name string
	Age  int
}

func (d *Db) GetUsersByName(name string) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUsersByName Preparation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUsersByName Query Err: ", err)
	}

	var r User
	users := []User{}

	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}

func (d *Db) GetUsersByID(id int) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE id=$1")
	if err != nil {
		fmt.Println("GetUsersByID Preparation Err: ", err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("GetUsersByID Query Err: ", err)
	}

	var r User
	users := []User{}

	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}
