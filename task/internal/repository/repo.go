package repository

import (
	"absapi/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// User operations

func (r *Repository) GetUser(persID int) ([]entity.User, error) {
	var user []entity.User
	query := fmt.Sprintf("Select * from users where persid=%d", persID)
	err := r.db.Select(&user, query)
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, ErrNotFound
	}
	return user, nil
}

func (r *Repository) InsertUser(user entity.Response) {
	u := user.ResponseBody
	_, err := r.db.Exec(`INSERT INTO  users(armid, birthdate, branchid, depname, departmentcode, email, empstate, empstatename,
                   gender, grade, localcode, login, name, persid, phonenumber, postid, postname, state, tabnum, userid, userstate)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`,
		u[0].ArmId, u[0].BirthDate, u[0].BranchId, u[0].DepName, u[0].DepartmentCode, u[0].EMail, u[0].EmpState, u[0].EmpStateName,
		u[0].Gender, u[0].Grade, u[0].LocalCode, u[0].Login, u[0].Name, u[0].PersId, u[0].PhoneNumber, u[0].PostId,
		u[0].PostName, u[0].State, u[0].TabNum, u[0].UserId, u[0].UserState)
	if err != nil {
		log.Println("repo - InsertUser - error", err)
	}
}

func (r *Repository) InsertUserFromMap(user map[string]interface{}) (entity.User, error) {

	aa := fmt.Sprintf("%v", user["brmId"])
	fmt.Println("11111 ", aa)
	_, err := r.db.Exec(`INSERT INTO  users(armid, birthdate, branchid, depname, departmentcode, email, empstate, empstatename,
                   gender, grade, localcode, login, name, persid, phonenumber, postid, postname, state, tabnum, userid, userstate)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`,
		aa,
		user["birthDate"].(string),
		int(user["branchId"].(float64)),
		user["depName"].(string),
		user["departmentCode"].(string),
		user["eMail"].(string),
		user["empState"].(string),
		user["empStateName"].(string),
		user["gender"].(string),
		user["grade"].(string),
		user["localCode"].(string),
		user["login"].(string),
		user["name"].(string),
		int(user["persId"].(float64)),
		int(user["phoneNumber"].(float64)),
		int(user["postId"].(float64)),
		user["postName"].(string),
		user["state"].(string),
		user["tabNum"].(string),
		int(user["userId"].(float64)),
		user["userState"].(string),
	)
	u, err := r.GetUser(int(user["persId"].(float64)))
	if err != nil {
		return entity.User{}, err
	}
	return u[0], err
}

// token operations

func (r *Repository) GetToken() (*entity.Token, error) {
	var t []entity.Token
	err := r.db.Select(&t, "select * from tokens")
	if err != nil {
		return &entity.Token{}, err
	}
	if len(t) == 0 {
		log.Println(ErrNotFound)
		return &entity.Token{}, ErrNotFound
	}
	return &t[0], nil
}

func (r *Repository) InsertToken(t *entity.Token) {
	r.db.MustExec("INSERT INTO  tokens VALUES ($1, $2)", t.TokenName, t.ExpiresAt)
}

func (r *Repository) UpdateToken(token *entity.Token) {
	r.db.MustExec("UPDATE tokens set token=$1 and expires_at=$2", token.TokenName, token.ExpiresAt)
}
