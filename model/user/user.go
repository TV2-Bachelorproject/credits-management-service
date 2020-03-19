package user

import (
	"errors"

	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

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
