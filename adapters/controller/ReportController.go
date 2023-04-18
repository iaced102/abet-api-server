package controller

import (
	"aBet/adapters/ports/incoming"
	"aBet/usecase/service"
	dservice "aBet/usecase/service/document"
	rService "aBet/usecase/service/report"
	"errors"
	"fmt"
	"net/http"
)

type reportController struct {
	createReportService       rService.CreateReportService
	createDocumentService     dservice.CreateDocumentService
	createDetailReportService service.CreateDetailReportService
}

type ReportController interface {
	CreateNewReport(c *Context) error
}

func NewReportController(cRS rService.CreateReportService, cDS dservice.CreateDocumentService, cDRS service.CreateDetailReportService) ReportController {
	return &reportController{
		createReportService:       cRS,
		createDocumentService:     cDS,
		createDetailReportService: cDRS,
	}
}

func (rC *reportController) CreateNewReport(c *Context) error {
	var DocumentConfig incoming.CreateReport
	fmt.Println(DocumentConfig, "___________________________")
	c.Bind(&DocumentConfig)
	fmt.Println(DocumentConfig)
	if DocumentConfig.DocumentId == "" {
		if DocumentConfig.Name == "" {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params name"))
		}
		if DocumentConfig.SuperviserId == "" {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params SuperViderId"))
		}
		if len(DocumentConfig.AssessorId) == 0 {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least AssenorId"))
		}
		if len(DocumentConfig.VerifierId) == 0 {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least VerifierId"))
		}
		if DocumentConfig.EvaluateField == "" {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params EvaluateField"))
		}
		if DocumentConfig.IdentifierId == "" {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params IdentifierId"))
		}
		if len(DocumentConfig.ListStudent) == 0 {
			c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least student"))
		}
		userId := c.AuthObject.GetUserId()
		document, e := rC.createDocumentService.CreateDocument(userId, DocumentConfig.Name, DocumentConfig.AssessorId, DocumentConfig.VerifierId, DocumentConfig.SuperviserId)
		if e != nil {
			return c.Output(http.StatusBadRequest, nil, e)
		}

		if e != nil {
			c.Output(http.StatusBadRequest, nil, e)
		}
		report, e := rC.createReportService.CreateReport(document.Id, DocumentConfig.EvaluateField, DocumentConfig.ListStudent)
		if e != nil {
			return c.Output(http.StatusBadRequest, nil, e)
		}
		err := rC.createDetailReportService.CreateListDetailReport(DocumentConfig.ListStudent, report.Id, report.Field)
		if err != nil {
			return c.Output(http.StatusBadRequest, nil, err)
		}
		return c.Output(http.StatusCreated, nil, nil)
		// for _,student := range DocumentConfig.ListStudent{

		// }

	} else {
		// document :=
	}
	return nil
}
