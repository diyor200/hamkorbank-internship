package repository

import (
	"errors"
	"fmt"
	"github.com/diyor200/gin-middleware-blogpost/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Repo struct {
	DB *sqlx.DB
}

func NewRepo(DB *sqlx.DB) *Repo {
	return &Repo{DB}
}

// User methods
func (r *Repo) CreateUser(user entity.UserInput) error {
	_, err := r.DB.Exec("INSERT INTO users(username, email, password) VALUES ($1,$2,$3)", user.Username, user.Email, user.Password)
	//res := r.db.Table("users").Create(&user)
	return err
}

func (r *Repo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	//res := r.db.Table("users").Scan(&users)
	err := r.DB.Select(&users, "select * from users")
	return users, err
}

func (r *Repo) GetUserID(email string) (int, error) {
	var user_id int
	err := r.DB.Select(&user_id, "select id from users where email=$1", email)
	return user_id, err
}

func (r *Repo) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	//res := r.db.Table("users").First(&user, "email=?", email)
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	fmt.Println("GetUserByEmail - error:", err)
	return user, err
}

func (r *Repo) DeleteUser(id int) error {
	//err := r.db.Table("users").Where("id=?", id).Delete(&entity.User{})
	_, err := r.DB.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

// CreatePost methods
func (r *Repo) CreatePost(input entity.CreateBlogInput) error {
	_, err := r.DB.Exec("INSERT INTO posts(user_id, post_title, post_body) VALUES ($1,$2,$3)",
		input.UserID, input.PostTittle, input.PostBody)
	return err
}

func (r *Repo) GetPosts() ([]entity.Blog, error) {
	var blogs []entity.Blog
	//res := r.db.Table("blogs").Scan(&blogs)
	err := r.DB.Select(&blogs, "SELECT * FROM posts")
	return blogs, err
}

func (r *Repo) GetPost(postId int) ([]entity.Blog, error) {
	var blog = []entity.Blog{}
	//res := r.db.Table("blogs").Scan(&blog)
	err := r.DB.Select(&blog, "SELECT * FROM posts WHERE id=$1", postId)
	if len(blog) == 0 {
		return nil, errors.New("not exists")
	}
	return blog, err
}

func (r *Repo) EditPost(input entity.BlogInput, userID int) error {
	var blog []entity.Blog
	fmt.Println("incomed post to edit: ", input)

	isAdmin := r.CheckAdmin(userID)

	log.Println("isAdmin := r.CheckAdmin(userID)", isAdmin)
	if isAdmin {
		err := r.DB.Select(&blog, "select * from posts where id=$1", input.PostID)
		fmt.Println("pg - EditPost - blog = ", blog)
		if len(blog) == 0 {
			return errors.New("not exists")
		}
		_, err = r.DB.Exec("UPDATE posts SET post_title=$1, post_body=$2 WHERE id=$3",
			input.PostTittle, input.PostBody, input.PostID)
		return err
	}
	err := r.DB.Select(&blog, "select * from posts where id=$1 and user_id=$2", input.PostID, userID)
	fmt.Println("pg - EditPost - blog = ", blog)
	if len(blog) == 0 {
		return errors.New("not exists")
	}
	_, err = r.DB.Exec("UPDATE posts SET post_title=$1, post_body=$2 WHERE id=$3 and user_id=$4",
		input.PostTittle, input.PostBody, input.PostID, userID)
	return err
}

func (r *Repo) DeletePost(postID, userID int) error {
	fmt.Println("postID, userID = ", postID, userID)
	var err error
	var blog []entity.Blog
	err = r.DB.Select(&blog, "select * from posts where id=$1", postID)
	fmt.Println("pg - EditPost - blog = ", blog)
	if len(blog) == 0 {
		return errors.New("not exists")
	}
	isAdmin := r.CheckAdmin(userID)
	log.Println("isAdmin := r.CheckAdmin(userID)", isAdmin)
	if isAdmin {
		_, err = r.DB.Exec("DELETE FROM posts WHERE id=$1", postID)
		return err
	}
	_, err = r.DB.Exec("DELETE FROM posts WHERE id=$1 and user_id=$2", postID, userID)
	err = r.DB.Select(&blog, "select * from posts where id=$1", postID)
	fmt.Println("blog == ", blog)
	if len(blog) > 0 {
		return errors.New("you are not owner")
	}
	log.Println(err)
	return err
}

func (r *Repo) CheckAdmin(userID int) bool {
	var user []entity.User
	err := r.DB.Select(&user, "SELECT * from users where id=$1", userID)
	fmt.Println("checkAdmin - error = ", err)
	fmt.Println("CheckAdmin - user = ", user)
	if user[0].Admin == 1 {
		return true
	}
	return false
}
