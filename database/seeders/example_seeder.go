package seeders

import (
	"io/ioutil"
)

func (s Seed) RoleSeeder() {
	q, err := ioutil.ReadFile(GetSourcePath() + "/scripts/roles_seeder.sql")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec(string(q))
	if err != nil {
		panic(err)
	}
}
