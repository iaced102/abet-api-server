package incoming

import "aBet/model"

/*
create by: Hoangnd
create at: 2023-01-01
des: Khai báo và xử lý thông tin nhận từ request
*/

type CreateReport struct {
	DocumentId    string          `json:"documentId"`
	Name          string          `json:"name" param:"name"`
	AssessorId    []string        `json:"assessorId"`
	VerifierId    []string        `json:"verifierId"`
	SuperviserId  string          `json:"superviserId"`
	EvaluateField string          `json:"evaluateField"`
	IdentifierId  string          `json:"identifierId"`
	ListStudent   []model.Student `json:"listStudent"`
}
