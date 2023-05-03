package model

type SODocument struct {
	Id           string   `pg:"id"`
	Name         string   `pg:"name" json:"name"`
	CreatedAt    string   `pg:"created_at"`
	CreatedBy    string   `pg:"created_by"`
	IdentifierId string   `pg:"identifier_id" json:"identifierId"`
	Desscription string   `pg:"description" json:"description"`
	tableName    struct{} `pg:"so_document"`
}
