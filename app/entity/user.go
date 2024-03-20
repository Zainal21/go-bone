package entity

type User struct {
	ID   int    `json:"id" db:"user_id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}
