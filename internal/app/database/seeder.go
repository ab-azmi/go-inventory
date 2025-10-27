package database

import "service/internal/app/database/seeder"

type DatabaseSeeder interface {
	Seed()
}

type data struct {
	DatabaseSeeder
}

func Seeder() {
	seeders := []data{
		{&seeder.TestingSeeder{}},
		{&seeder.ItemSeeder{}},
	}

	for _, seed := range seeders {
		seed.Seed()
	}
}
