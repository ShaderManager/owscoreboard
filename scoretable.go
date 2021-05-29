package main

import (
	"encoding/json"
	"log"
	"os"
)

type Scoretable struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
}

type RoleScoretable struct {
	Tank    *Scoretable `json:"tank"`
	Dps     *Scoretable `json:"dps"`
	Support *Scoretable `json:"sup"`
}

var table = &RoleScoretable{
	Tank:    &Scoretable{0, 0, 0},
	Dps:     &Scoretable{0, 0, 0},
	Support: &Scoretable{0, 0, 0},
}

func addWin(role string, inc int) {
	switch {
	case role == "tank":
		table.Tank.Wins += inc
	case role == "dps":
		table.Dps.Wins += inc
	case role == "sup":
		table.Support.Wins += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func addTie(role string, inc int) {
	switch {
	case role == "tank":
		table.Tank.Ties += inc
	case role == "dps":
		table.Dps.Ties += inc
	case role == "sup":
		table.Support.Ties += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func addLose(role string, inc int) {
	switch {
	case role == "tank":
		table.Tank.Losses += inc
	case role == "dps":
		table.Dps.Losses += inc
	case role == "sup":
		table.Support.Losses += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func resetStats() {
	table.Tank = &Scoretable{0, 0, 0}
	table.Dps = &Scoretable{0, 0, 0}
	table.Support = &Scoretable{0, 0, 0}
}

func loadJSON(filename string, v interface{}) error {
	fileObject, err := os.Open(filename) // For read access.
	if err == nil {
		decoder := json.NewDecoder(fileObject)

		if err := decoder.Decode(&v); err != nil {
			return err
		}
	}

	return err
}

func saveJSON(filename string, v interface{}) error {
	fileObject, err := os.Create(filename)

	if err == nil {
		encoder := json.NewEncoder(fileObject)

		// Write to the file
		if err := encoder.Encode(&v); err != nil {
			return err
		}
	}

	return err
}

func saveTable(filename string) error {
	return saveJSON(filename, table)
}

func loadTable(filename string) error {
	return loadJSON(filename, table)
}
