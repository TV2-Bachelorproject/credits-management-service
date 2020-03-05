package model

type Season struct {
	ID       int
	Title    string
	Programs []Program
	Credits  []Credit
}
