package model

import (
	"log"

	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func Seed() {
	users()
	creditGroups()
	people()
	categories()
	productions()
	series()
	seasons()
	programs()
}

func categories() {
	categories := []public.Category{
		{Name: "Test category"},
	}

	for _, category := range categories {
		db.Create(&category)
	}
}

func productions() {
	productions := []public.Production{
		{Editor: "Test editor"},
	}

	for _, production := range productions {
		db.Create(&production)
	}
}

func series() {
	series := []public.Serie{
		{
			Title: "Test serie",
		},
	}

	for _, serie := range series {
		db.Create(&serie)
	}
}

func seasons() {
	seasons := []public.Season{
		{
			Title:   "Test season 1",
			SerieID: 1,
		},
	}

	for _, season := range seasons {
		db.Create(&season)
	}
}

func programs() {
	programs := []public.Program{
		{
			Title:        "Test program 1",
			SerieID:      1,
			SeasonID:     1,
			CategoryID:   1,
			ProductionID: 1,
		},
	}

	for _, program := range programs {
		db.Create(&program)
	}
}

func people() {
	people := []private.Person{
		{
			Name:    "John Doe",
			Email:   "john@example.com",
			Address: "Testing street 3",
			City:    "Testing city",
			Postal:  "5000",
			Country: "Testing Country",
		},
		{
			Name:    "Jane Doe",
			Email:   "Jane@example.com",
			Address: "Testing street 3",
			City:    "Testing city",
			Postal:  "5000",
			Country: "Testing Country",
		},
	}

	for _, person := range people {
		db.Create(&person)
	}
}

func creditGroups() {
	groups := []public.CreditGroup{
		{Name: "TV 2 Redaktør"},
		{Name: "Producent"},
		{Name: "Redaktionschef"},
		{Name: "Postproducer"},
		{Name: "Produktionsleder"},
		{Name: "Tilrettelægger og Redaktion"},
		{Name: "Research"},
		{Name: "Fotograf"},
		{Name: "Drone"},
		{Name: "Klipper"},
		{Name: "Grafik"},
		{Name: "Grade"},
	}

	for _, group := range groups {
		db.Create(&group)
	}
}

func users() {
	u1, err := user.New("admin", "admin@example.com", "123456", user.Admin)

	if err != nil {
		log.Fatal(err)
	}

	u2, err := user.New("producer", "producer@example.com", "123456", user.Producer)

	if err != nil {
		log.Fatal(err)
	}

	db.Create(&u1)
	db.Create(&u2)
}
