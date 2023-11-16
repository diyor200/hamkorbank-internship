package entity

type User struct {
	Id       int    `db:"ID"`
	Username string `db:"USERNAME"`
	Password string `db:"PASSWORD"`
	Email    string `db:"EMAIL"`
	Admin    string `db:"ISADMIN"`
}

type Blog struct {
	Id         int    `db:"ID"`
	AuthorID   int    `db:"USER_ID"`
	PostTittle string `db:"POST_TITLE"`
	PostBody   string `db:"POST_BODY"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BlogInput struct {
	PostTittle string `json:"post_title"`
	PostBody   string `json:"post_body"`
}

type CreateBlogInput struct {
	UserID     int
	PostTittle string
	PostBody   string
}
