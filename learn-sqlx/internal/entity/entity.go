package entity

type UserInput struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Birth     BirthInfo `json:"birth_info"`
}

type BirthInfo struct {
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

type UserUpdateInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
