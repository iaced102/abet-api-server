package controller

import (
	"aBet/model"
	"aBet/usecase/service"
	"errors"
	"fmt"
	"net/http"
)

type sOController struct {
	createSOService service.CreateSOService
	getSOService    service.GetAllSODocumentService
}

type SOController interface {
	CreateNewSODocument(c *Context) error
	GetAllSODocument(c *Context) error
	GetDetailSODocument(c *Context) error
}

func NewSOController(cSS service.CreateSOService, gASS service.GetAllSODocumentService) SOController {
	return &sOController{
		createSOService: cSS,
		getSOService:    gASS,
	}
}

func (sOC *sOController) CreateNewSODocument(c *Context) error {
	fmt.Println("aaaaaaaaaaaaaaaa")
	if c.AuthObject.GetUserRole() != 0 {
		return c.Output(http.StatusForbidden, "dont have permission", errors.New("you are not admin role"))
	}
	sODocument := model.SODocument{}
	c.Bind(&sODocument)
	s, e := sOC.createSOService.CreateSO(sODocument.Name, sODocument.Desscription, sODocument.IdentifierId, c.AuthObject.GetUserId())
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusCreated, s, e)
}

func (sOC *sOController) GetAllSODocument(c *Context) error {
	so, e := sOC.getSOService.GetAllSO()
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, so, e)
}

func (sOC *sOController) GetDetailSODocument(c *Context) error {
	id := c.Param("id")
	so, e := sOC.getSOService.GetDetailSODocument(id)
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, so, e)
}
