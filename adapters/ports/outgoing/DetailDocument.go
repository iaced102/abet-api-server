package outgoing

import "aBet/model"

/*
create by: Hoangnd
create at: 2023-01-01
des: Khai báo và xử lý thông tin nhận từ request
*/

type GetDetailDocument struct {
	Id           string   `json:"id"`
	CreatedBy    string   `json:"createdBy"`
	Name         string   `json:"name"`
	CreatedAt    string   `json:"createdAt"`
	UpDatedAt    string   `json:"updatedAt"`
	AssessorId   []string `json:"assessorId"`
	VerifierId   []string `json:"verifierId"`
	SuperviserId string   `json:"superviserId"`
	// IdentifierId string `pg:"identifier_id"`
	EvaluteField string   `json:"evaluateField"`
	Data         []Report `json:"data"`
}

type Report struct {
	Id     string               `json:"id"`
	Field  string               `json:"field"`
	Detail []model.DetailReport `json:"Detail"`
}

// type DetailReport struct {
// 	Id        string `json:"id"`
// 	StudentId string `json:"studentId"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	ClassId   string `json:"classIdd"`
// 	ReportId  string `json:"reportId"`
// 	Value     string `json:"value"`
// }
