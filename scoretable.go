package main

import "log"

type Scoretable struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
}

var tankScoretable = &Scoretable{0, 0, 0}
var dpsScoretable = &Scoretable{0, 0, 0}
var supScoretable = &Scoretable{0, 0, 0}

func addWin(role string, inc int) {
	switch {
	case role == "tank":
		tankScoretable.Wins += inc
	case role == "dps":
		dpsScoretable.Wins += inc
	case role == "sup":
		supScoretable.Wins += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func addTie(role string, inc int) {
	switch {
	case role == "tank":
		tankScoretable.Ties += inc
	case role == "dps":
		dpsScoretable.Ties += inc
	case role == "sup":
		supScoretable.Ties += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func addLose(role string, inc int) {
	switch {
	case role == "tank":
		tankScoretable.Losses += inc
	case role == "dps":
		dpsScoretable.Losses += inc
	case role == "sup":
		supScoretable.Losses += inc
	default:
		log.Printf("Unknown role: %s\n", role)
	}
}

func resetStats() {
	tankScoretable = &Scoretable{0, 0, 0}
	dpsScoretable = &Scoretable{0, 0, 0}
	supScoretable = &Scoretable{0, 0, 0}
}
