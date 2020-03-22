package model

import (
	"time"

	"github.com/TV2-Bachelorproject/server/model/public"

	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

var tables = []interface{}{
	&private.Person{},
	&public.Category{},
	&public.Genre{},
	&public.Credit{},
	&public.CreditGroup{},
	&public.Production{},
	&public.Program{},
	&public.Season{},
	&public.Serie{},
	&user.User{},
}

func Migrate() {
	time.Sleep(500 * time.Millisecond)
	db.Migrate(tables...)
}

func Reset() {
	time.Sleep(500 * time.Millisecond)
	db.Reset(tables...)
}
