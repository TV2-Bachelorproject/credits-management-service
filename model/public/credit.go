package public

type CreditGroup struct {
	ID    int
	Title string
}

type Credit struct {
	ID      int
	Persons []Person
	Group   CreditGroup
}
