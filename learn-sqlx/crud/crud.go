package crud

import (
	"fmt"
	"github.com/diyor200/learn-sqlx/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Crud struct {
	db *sqlx.DB
}

type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Phone     string `db:"phone"`
	Birth     BirthInfo
}

type BirthInfo struct {
	Country     string `db:"country"`
	CountryCode string `db:"country_code"`
}

func NewCrud(db *sqlx.DB) *Crud {
	return &Crud{
		db: db,
	}
}

func (c *Crud) InsertUsers(users []User) {
	for _, user := range users {

		c.db.MustExec("INSERT INTO users (first_name, last_name, email, password, phone) "+
			"VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Password, user.Phone)

		c.db.MustExec("INSERT INTO birth_info (user_id, country, country_code)"+
			"VALUES ((SELECT id FROM users WHERE phone = $1), $2, $3, $4)", user.Phone, user.Birth.Country, user.Birth.CountryCode)
	}
}

func (c *Crud) InsertUser(user entity.UserInput) {
	c.db.MustExec("INSERT INTO users (first_name, last_name, email, password, phone) "+
		"VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Password, user.Phone)

	c.db.MustExec("INSERT INTO birth_info (user_id, country, country_code)"+
		"VALUES ((SELECT id FROM users WHERE phone = $1), $2, $3)", user.Phone, user.Birth.Country, user.Birth.CountryCode)

}

func (c *Crud) GetUsers() ([]User, error) {
	var users []User
	query := `SELECT u.id, u.first_name, u.last_name,u.email,u.password,u.phone,b.country AS "birth.country",
            b.country_code AS "birth.country_code" FROM users AS u
        	LEFT JOIN birth_info AS b ON b.user_id = u.id`

	err := c.db.Select(&users, query)
	return users, err
}

func (c *Crud) GetUser(phone string) ([]User, error) {
	var user []User

	query := fmt.Sprintf(`select users.first_name, users.last_name, users.email,
		users.password, users.phone, birth_info.country,
		birth_info.country_code from users left join birth_info on birth_info.user_id = users.id where users.phone=%s`, phone)
	err := c.db.Select(&user, query)
	return user, err
}

func (c *Crud) DeleteUser(id int) {
	c.db.MustExec("DELETE FROM birth_info where user_id=$1", id)
	c.db.MustExec("DELETE FROM users where id=$1", id)
}

func (c *Crud) Update(user entity.UserUpdateInput) {
	c.db.MustExec(`UPDATE users SET first_name=$1, last_name=$2,  phone=$3 where email=$4`,
		user.FirstName, user.LastName, user.Phone, user.Email)
}
