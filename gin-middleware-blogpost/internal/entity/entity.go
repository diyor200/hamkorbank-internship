package entity

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Admin    int    `db:"isadmin"`
}

type Blog struct {
	Id         int    `db:"id"`
	AuthorID   int    `db:"user_id"`
	PostTittle string `db:"post_title"`
	PostBody   string `db:"post_body"`
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
	PostID     int    `json:"post_id"`
	PostTittle string `json:"post_title"`
	PostBody   string `json:"post_body"`
}

type CreateBlogInput struct {
	UserID     int
	PostTittle string
	PostBody   string
}
