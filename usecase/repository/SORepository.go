package repository

import "aBet/model"

type CreateSORepository interface {
	CreateSO(so *model.SODocument) error
}

type GetSORepository interface {
	GetAllSODocument(*[]model.SODocument) error
	GetDetailSODocument(*model.SODocument) error
}
