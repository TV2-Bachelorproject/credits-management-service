package user

import (
	"errors"
	"regexp"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type ValidationError []error

func (errs ValidationError) Error() (str string) {
	for _, err := range errs {
		str += err.Error() + "\n\t"
	}

	return str
}

type Type int

const (
	Admin Type = iota
	Producer
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string `json:"-"`
	Type     Type
	Token    string `json:"-"`
}

//UserType for graphql
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.Int},
			"name":     &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"password": &graphql.Field{Type: graphql.String},
			"type":     &graphql.Field{Type: graphql.Int},
			"token":    &graphql.Field{Type: graphql.String},
		},
	},
)

func New(name, email, password string, t Type) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return User{}, err
	}

	return User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		Type:     t,
	}, nil
}

func Find(id uint) User {
	u := User{}
	db.Where("id = ?", id).First(&u)
	return u
}

type Users []User

func All() Users {
	u := Users{}
	db.Model(User{}).Find(&u)
	return u
}

func Authorize(email, password string) (User, error) {
	user := User{Email: email}
	invalid := errors.New("invalid credentials")

	db.Where(user).First(&user)

	if user.ID == 0 {
		return User{}, invalid
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return User{}, invalid
	}

	return user, nil
}

func (u User) Invalid() ValidationError {
	e := ValidationError{}
	set := func(msg string) {
		e = append(e, errors.New(msg))
	}

	if u.Name == "" {
		set("missing name")
	}

	if u.Email == "" {
		set("missing email")
	}

	validEmail, err := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", u.Email)

	if err != nil {
		set(err.Error())
	}

	if !validEmail {
		set("invalid email")
	}

	if u.Password == "" {
		set("invalid password")
	}

	if u.Type != Admin && u.Type != Producer {
		set("invalid user type")
	}

	if len(e) > 0 {
		e = append(ValidationError{errors.New("invalid user:")}, e...)
	}

	return e
}
