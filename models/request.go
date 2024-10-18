package models

type CreateRuleRequest struct {
	Rule string `json:"rule"`
}

type EvaluateRuleRequest struct {
	Salary     int    `json:"salary"`
	Age        int    `json:"age"`
	Experience int    `json:"experience"`
	Department string `json:"department"`
}
