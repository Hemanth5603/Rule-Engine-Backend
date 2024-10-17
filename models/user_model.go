package models

type UserModel struct {
	Age        int    `json:"age"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
	Experience int    `json:"experience"`
}
