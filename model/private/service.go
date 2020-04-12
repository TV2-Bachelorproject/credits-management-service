package private

import (
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	Name  string
	Token string
}

func (s Service) Find(id uint) Service {
	db.Model(s).Where("id = ?", id).First(&s)
	return s
}

func (s Service) FindServiceWithToken(token string) Service {
	db.Model(s).Where("token = ?", token).First(&s)
	return s
}

type Services []Service

func (s Services) Find() Services {
	db.Model(Service{}).Find(&s)
	return s
}
