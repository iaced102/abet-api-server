package model

type Document struct {
	Id           string   `json:"id" param:"id" db:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	CreatedBy    string   `pg:"created_by"`
	Name         string   `pg:"name"`
	CreatedAt    string   `pg:"created_at"`
	UpDatedAt    string   `pg:"updated_at"`
	AssessorId   []string `pg:"assessor_id"`
	VerifierId   string   `pg:"verifier_id"`
	tableName    struct{} `pg:"document"`
	SODocumentId string   `pg:"so_document_id"`
	EvaluteField string   `pg:"evaluate_field"`
}
