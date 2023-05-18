package model

type Report struct {
	Id         string   `json:"id" param:"id" pg:"id" query:"id" form:"id" pg:"id"   primary:"true"`
	DocumentId string   `pg:"document_id"`
	Field      string   `pg:"field"`
	tableName  struct{} `pg:"report"`
}

type DetailReport struct {
	Id         string   `json:"id" param:"id" pg:"id" query:"id" form:"id" pg:"id"   primary:"true" form:"id"`
	StudentId  string   `pg:"student_id" json:"studentId" form:"studentId"`
	FirstName  string   `pg:"first_name" json:"firstName" form:"firstName"`
	LastName   string   `pg:"last_name" json:"lastName" form:"lastName"`
	ClassId    string   `pg:"class_id" json:"classId" form:"classId"`
	ReportId   string   `pg:"report_id" json:"reportId" form:"reportId"`
	Value      string   `pg:"value" json:"value" form:"value"`
	Major      string   `pg:"major" json:"major" form:"major"`
	Course     string   `pg:"course" json:"course" form:"course"`
	AssessorId string   `pg:"assessor_id" json:"assesserId" form:"assesserId"`
	tableName  struct{} `pg:"detail_report"`
}

type Student struct {
	Major      string `json:"major"`
	Course     string `json:"course"`
	AssessorId string `json:"assessorId" form:"assesserId"`
	StudentId  string `json:"studentId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	ClassId    string `json:"classId"`
}
