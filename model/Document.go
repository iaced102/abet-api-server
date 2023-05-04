package model

type Document struct {
	Id           string   `json:"id" json:"id" param:"id" db:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	CreatedBy    string   `pg:"created_by" json:"createdBy"`
	Name         string   `pg:"name" json:"name"`
	CreatedAt    string   `pg:"created_at" json:"createdAt"`
	UpDatedAt    string   `pg:"updated_at" json:"updateAt"`
	AssessorId   []string `pg:"assessor_id" json:"assessorId"`
	VerifierId   string   `pg:"verifier_id" json:"verifierId"`
	tableName    struct{} `pg:"document"`
	SODocumentId string   `pg:"so_document_id" json:"sODocumentId"`
	EvaluteField string   `pg:"evaluate_field" json:"evaluteField"`
}
