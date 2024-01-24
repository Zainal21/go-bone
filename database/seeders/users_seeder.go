package seeders

import (
	"io/ioutil"
)

func (s Seed) UserSeeder() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/users_seeder.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
