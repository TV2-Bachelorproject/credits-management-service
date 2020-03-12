package model

import (
	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func Migrate() {
	db.Migrate(
		&private.Person{},
		&public.Credit{},
		&public.CreditGroup{},
		&public.Program{},
		&public.Season{},
		&public.Serie{},
		&public.Serie{},
	)
}

func Reset() {
	db.Reset(
		&private.Person{},
		&public.Credit{},
		&public.CreditGroup{},
		&public.Program{},
		&public.Season{},
		&public.Serie{},
		&public.Serie{},
	)
}
