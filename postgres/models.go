package postgres

import "time"

type Event struct {
	EventID        int       `json:"event_id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	GameName       string    `json:"game_name"`
	NumberEntrants int       `json:"number_entrants"`
	/* Other          interface{} */
}

type Entrant struct {
	EntrantID      string `json:"entrant_id"`
	EntrantTag     string `json:"entrant_tag"`
	InitialSeed    int    `json:"initial_seed"`
	FinalPlacement int    `json:"final_placement"`
	/* Other          interface{} */
}

type Set struct {
	SetID          int    `json:"set_id"`
	Entrant1ID     int    `json:"entrant1_id"`
	Entrant2ID     int    `json:"entrant2_id"`
	Entrant1Result string `json:"entrant1_result"`
	Entrant2Result string `json:"entrant2_result"`
	/* Other          interface{} */
}
