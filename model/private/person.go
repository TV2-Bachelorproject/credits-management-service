package private

import (
	"errors"
	"regexp"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type ValidationError []error

func (errs ValidationError) Error() (str string) {
	for _, err := range errs {
		str += err.Error() + "\n\t"
	}

	return str
}

type Person struct {
	gorm.Model
	Name    string
	Email   string
	Address string
	City    string
	Postal  string
	Country string
}

func (p Person) Find(id uint) Person {
	db.Model(p).Where("id = ?", id).First(&p)
	return p
}

func (p Person) Invalid() ValidationError {
	e := ValidationError{}
	set := func(msg string) {
		e = append(e, errors.New(msg))
	}

	if p.Name == "" {
		set("missing name")
	}

	if p.Email == "" {
		set("missing email")
	}

	validEmail, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", p.Email)

	if err != nil {
		set(err.Error())
	}

	if !validEmail {
		set("invalid email")
	}

	if p.Address == "" {
		set("missing address")
	}

	if p.City == "" {
		set("missing city")
	}

	if p.Country == "" {
		set("missing country")
	}

	if p.Postal == "" {
		set("missing postal")
	}

	if len(e) > 0 {
		e = append(ValidationError{errors.New("invalid person:")}, e...)
	}

	return e
}

type People []Person

func (s People) Find() People {
	db.Model(&Person{}).Find(&s)
	return s
}
