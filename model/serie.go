package model

type Serie struct {
	ID      int
	Title   string
	Seasons []Season
	Credits []Credit
}
