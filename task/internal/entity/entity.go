package entity

type Response struct {
	Code         int            `json:"code"`
	Msg          string         `json:"msg"`
	ResponseBody []ResponseData `json:"responseBody"`
}

type ResponseData struct {
	ArmId          string `json:"armId"`
	BirthDate      string `json:"birthDate"`
	BranchId       int    `json:"branchId"`
	DepName        string `json:"depName"`
	DepartmentCode string `json:"departmentCode"`
	EMail          string `json:"eMail"`
	EmpState       string `json:"empState"`
	EmpStateName   string `json:"empStateName"`
	Gender         string `json:"gender"`
	Grade          string `json:"grade"`
	LocalCode      string `json:"localCode"`
	Login          string `json:"login"`
	Name           string `json:"name"`
	PersId         int    `json:"persId"`
	PhoneNumber    int    `json:"phoneNumber"`
	PostId         int    `json:"postId"`
	PostName       string `json:"postName"`
	State          string `json:"state"`
	TabNum         string `json:"tabNum"`
	UserId         int    `json:"userId"`
	UserState      string `json:"userState"`
}

type User struct {
	ID             int    `db:"id"`
	ArmId          string `db:"armid"`
	BirthDate      string `db:"birthdate"`
	BranchId       int    `db:"branchid"`
	DepName        string `db:"depname"`
	DepartmentCode string `db:"departmentcode"`
	EMail          string `db:"email"`
	EmpState       string `db:"empstate"`
	EmpStateName   string `db:"empstatename"`
	Gender         string `db:"gender"`
	Grade          string `db:"grade"`
	LocalCode      string `db:"localcode"`
	Login          string `db:"login"`
	Name           string `db:"name"`
	PersId         int    `db:"persid"`
	PhoneNumber    int    `db:"phonenumber"`
	PostId         int    `db:"postid"`
	PostName       string `db:"postname"`
	State          string `db:"state"`
	TabNum         string `db:"tabnum"`
	UserId         int    `db:"userid"`
	UserState      string `db:"userstate"`
}

type Token struct {
	TokenName string `db:"token"`
	ExpiresAt int    `db:"expires_at"`
}
