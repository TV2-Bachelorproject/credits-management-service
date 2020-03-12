package private

type Address struct {
	Address []string `gorm:"type:varchar(100)[]"`
	City    string
	Postal  string
	Country string
}
