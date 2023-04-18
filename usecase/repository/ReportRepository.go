package repository

import "aBet/model"

type ReportRepository interface {
	CreateReport(*model.Report) error
}
