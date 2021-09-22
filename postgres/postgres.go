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

func scanEventRows(rows *sql.Rows) []Event {
	var event Event
	events := []Event{}

	for rows.Next() {
		err := rows.Scan(
			&event.Name,
			&event.Date,
			&event.GameName,
			&event.NumberEntrants,
		)
		if err != nil {
			fmt.Println("Error while scanning Event rows: ", err)
		}
		events = append(events, event)
	}

	return events
}

func (d *Db) GetRows(sqlStmt string, operation string, query string) *sql.Rows {
	stmt, err := d.Prepare(sqlStmt)
	if err != nil {
		fmt.Println(operation, " Preparation Error: ", err)
	}

	var rows *sql.Rows

	rows, err = stmt.Query(query)
	if err != nil {
		fmt.Println(operation, " Query Error: ", err)
	}

	return rows
}

func (d *Db) GetEventsByID(id int) []Event {
	rows := d.GetRows(
		"SELECT * FROM tourney_events WHERE event_id=$1",
		"GetEventsByID", strconv.Itoa(id),
	)

	var event Event
	events := []Event{}

	for rows.Next() {
		err := rows.Scan(
			&event.EventID,
			&event.Name,
			&event.Date,
			&event.GameName,
			&event.NumberEntrants,
		)
		if err != nil {
			fmt.Println("Error while scanning Event rows: ", err)
		}
		events = append(events, event)
	}

	return events
}

func (d *Db) GetEntrantsByID(id int) []Entrant {
	rows := d.GetRows(
		"SELECT * FROM tourney_entrants WHERE entrant_id=$1",
		"GetEntrantsByID", strconv.Itoa(id),
	)

	var entrant Entrant
	entrants := []Entrant{}

	for rows.Next() {
		err := rows.Scan(
			&entrant.EntrantID,
			&entrant.EntrantTag,
			&entrant.InitialSeed,
			&entrant.FinalPlacement,
		)
		if err != nil {
			fmt.Println("Error while scanning Entrant rows: ", err)
		}
		entrants = append(entrants, entrant)
	}

	return entrants
}

func (d *Db) GetSetsByID(id int) []Set {
	rows := d.GetRows(
		"SELECT * FROM tourney_sets WHERE set_id=$1",
		"GetSetsByID", strconv.Itoa(id),
	)

	var set Set
	sets := []Set{}

	for rows.Next() {
		err := rows.Scan(
			&set.SetID,
			&set.Entrant1ID,
			&set.Entrant2ID,
			&set.Entrant1Result,
			&set.Entrant2Result,
		)
		if err != nil {
			fmt.Println("Error while scanning Set rows: ", err)
		}
		sets = append(sets, set)
	}

	return sets
}

func (d *Db) SaveSet(setParams SetParams) int {
	stmt := `
		INSERT INTO tourney_sets (entrant1_id, entrant2_id, entrant1_result, entrant2_result)
		WHERE ($1, $2, $3, $4)
		RETURNING entrant_id`
	var entrantID int
	err := d.QueryRow(
		stmt,
		setParams.Entrant1ID,
		setParams.Entrant2ID,
		setParams.Entrant1Result,
		setParams.Entrant2Result,
	).Scan(&entrantID)
	if err != nil {
		fmt.Println("Error while inserting Set record: ", err)
	}

	return entrantID
}

type SetParams struct {
	Entrant1ID     int    `json:"entrant1_id"`
	Entrant2ID     int    `json:"entrant2_id"`
	Entrant1Result string `json:"entrant1_result"`
	Entrant2Result string `json:"entrant2_result"`
}
