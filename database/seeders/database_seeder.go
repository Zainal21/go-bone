package seeders

import (
	"log"
	"reflect"

	"github.com/jmoiron/sqlx"
)

// Seed type
type Seed struct {
	db *sqlx.DB
}

func Execute(db *sqlx.DB, seedMethodNames ...string) {
	s := Seed{db}

	seederTable := []map[string]interface{}{
		{"name": "UserSeeder"},
	}

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeders in order...")
		for _, seeder := range seederTable {
			seedName, ok := seeder["name"].(string)
			if !ok {
				log.Println("Invalid seeder entry in seederTable")
				continue
			}
			seed(s, seedName)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}
