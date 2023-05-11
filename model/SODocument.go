package model

type SODocument struct {
	Id           string   `pg:"id" json:"id" form:"id"`
	Name         string   `pg:"name" json:"name"`
	CreatedAt    string   `pg:"created_at" json:"createAt"`
	CreatedBy    string   `pg:"created_by" json:"createBy"`
	IdentifierId string   `pg:"identifier_id" json:"identifierId"`
	Desscription string   `pg:"description" json:"description"`
	tableName    struct{} `pg:"so_document"`
}
