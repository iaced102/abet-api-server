package model

type Report struct {
	Id         string   `json:"id" param:"id" pg:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	DocumentId string   `pg:"document_id"`
	Field      string   `pg:"field"`
	tableName  struct{} `pg:"report"`
}

type DetailReport struct {
	Id        string   `json:"id" param:"id" pg:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	StudentId string   `pg:"student_id"`
	FirstName string   `pg:"first_name"`
	LastName  string   `pg:"last_name"`
	ClassId   string   `pg:"class_id"`
	ReportId  string   `pg:"report_id"`
	Value     string   `pg:"value"`
	tableName struct{} `pg:"detail_report"`
}

type Student struct {
	StudentId string `json:"studentId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ClassId   string `json:"classId"`
}
