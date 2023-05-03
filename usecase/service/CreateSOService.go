package service

import (
	"aBet/model"
	"aBet/usecase/repository"
	"time"

	"github.com/google/uuid"
)

type CreateSOService interface {
	CreateSO(name string, description string, userId string) (model.SODocument, error)
}

type createSOService struct {
	soRepository repository.CreateSORepository
}

func NewCreateSOService(s repository.CreateSORepository) CreateSOService {
	return &createSOService{
		soRepository: s,
	}
}

func (cSS *createSOService) CreateSO(name string, description string, userId string) (model.SODocument, error) {
	sODocument := model.SODocument{
		Id:           uuid.NewString(),
		Name:         name,
		Desscription: description,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:    userId,
	}
	e := cSS.soRepository.CreateSO(&sODocument)
	return sODocument, e
}
