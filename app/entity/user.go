package entity

type User struct {
	Id        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Name      string `json:"name" db:"name"`
	Password  string `json:"password" db:"password"`
	Status    string `json:"status" db:"status"`
	Role      string `json:"role" db:"role"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
