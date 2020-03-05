package model

type Address struct {
	ID      int
	Lines   []string
	City    string
	Postal  string
	Country string
}

type Person struct {
	ID      int
	Name    string
	Email   string
	Address Address
}
