package controller

import (
	"aBet/adapters/ports/incoming"
	"aBet/adapters/ports/outgoing"
	"aBet/model"
	"aBet/usecase/service"
	dservice "aBet/usecase/service/document"
	rService "aBet/usecase/service/report"
	"errors"
	"fmt"
	"net/http"
)

type reportController struct {
	createReportService       rService.CreateReportService
	getReportService          rService.GetReportService
	getDetailReportService    service.GetDetailReportService
	createDocumentService     dservice.CreateDocumentService
	getDocumentService        dservice.GetDocumentService
	createDetailReportService service.CreateDetailReportService
	editDetailReportService   service.EditDetailReportService
}

type ReportController interface {
	CreateNewReport(c *Context) error
	GetDetailDocument(c *Context) error
	GetAllDocument(c *Context) error
	GetAllDocumentById(c *Context) error
	SubmitReport(c *Context) error
	GetAllPIbySOId(c *Context) error
	DeleteDocument(c *Context) error
}

type reportByDocumentForm struct {
	reportTemplate []string
}

var documentForm = map[string]([]string){
	"PI.3.1": []string{"PI.3.1a", "PI.3.1b"},
	"PI.3.2": []string{"PI.3.2a", "PI.3.2b"},
	"PI.3.3": []string{"0"},
	"PI.3.4": []string{"PI.3.4a", "PI.3.4b"},
}
var detailReportTemplate = map[string]string{
	"PI.3.1_PI.3.1a": `{"Thành phần 1":"","Thành phần 2":"","Thành phần 3":""}`,
	"PI.3.1_PI.3.1b": `{"Thành phần 1":"","Thành phần 2":""}`,
	"PI.3.2_PI.3.2a": `{"Thành phần 1":""}`,
	"PI.3.2_PI.3.2b": `{"Thành phần 1":""}`,
	"PI.3.3_0":       `{"Thành phần 1":"","Thành phần 2":"","Thành phần 3":"","Thành phần 4":""}`,
	"PI.3.4_PI.3.4a": `{"Thành phần 1":"","Thành phần 2":""}`,
	"PI.3.4_PI.3.4b": `{"Thành phần 1":"","Thành phần 2":""}`,
}

func NewReportController(cRS rService.CreateReportService, cDS dservice.CreateDocumentService, cDRS service.CreateDetailReportService, gDS dservice.GetDocumentService, gRS rService.GetReportService, gDRS service.GetDetailReportService, eDRS service.EditDetailReportService) ReportController {
	return &reportController{
		createReportService:       cRS,
		getReportService:          gRS,
		createDocumentService:     cDS,
		getDocumentService:        gDS,
		createDetailReportService: cDRS,
		getDetailReportService:    gDRS,
		editDetailReportService:   eDRS,
	}
}
func (rC *reportController) GetAllDocument(c *Context) error {
	doc, e := rC.getDocumentService.GetAllDocument()
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, doc, nil)
}

func (rC *reportController) GetAllDocumentById(c *Context) error {
	userId := c.AuthObject.GetUserUserName()
	fmt.Println(userId)
	doc, e := rC.getDocumentService.GetAllDocumentByUserId(userId)
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, doc, nil)
}
func (rC *reportController) CreateNewReport(c *Context) error {
	var DocumentConfig incoming.CreateReport
	c.Bind(&DocumentConfig)
	if DocumentConfig.Name == "" {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params name"))
	}

	if len(DocumentConfig.AssessorId) == 0 {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least AssenorId"))
	}
	if DocumentConfig.VerifierId == "" {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least VerifierId"))
	}
	if DocumentConfig.EvaluateField == "" {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params EvaluateField"))
	}
	if DocumentConfig.IdentifierId == "" {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params IdentifierId"))
	}
	if len(DocumentConfig.ListStudent) == 0 {
		return c.Output(http.StatusBadRequest, nil, errors.New("Create Report need params the least student"))
	}

	ids, er := documentForm[DocumentConfig.EvaluateField]
	if !er {
		return c.Output(http.StatusBadRequest, nil, errors.New("invalid evaluateField"))
	}
	fmt.Println(ids)

	userId := c.AuthObject.GetUserId()
	document, e := rC.createDocumentService.CreateDocument(userId, DocumentConfig.Name, DocumentConfig.EvaluateField, DocumentConfig.AssessorId, DocumentConfig.VerifierId, DocumentConfig.SODocumentId)
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}

	if e != nil {
		c.Output(http.StatusBadRequest, nil, e)
	}
	for _, i := range ids {
		report, e := rC.createReportService.CreateReport(document.Id, i, DocumentConfig.ListStudent)
		if e != nil {
			return c.Output(http.StatusBadRequest, nil, e)
		}
		templateName := fmt.Sprint(document.EvaluteField, "_"+i)
		err := rC.createDetailReportService.CreateListDetailReport(DocumentConfig.ListStudent, report.Id, report.Field, detailReportTemplate[templateName])
		if err != nil {
			return c.Output(http.StatusBadRequest, nil, err)
		}
	}

	return c.Output(http.StatusCreated, document, nil)

}

func (rC *reportController) GetDetailDocument(c *Context) error {
	doc, e := rC.getDocumentService.GetDocument(c.Param("documentId"))
	if e != nil {
		return c.Output(http.StatusBadRequest, doc, e)
	}
	outgoingGetDetailDocument := outgoing.GetDetailDocument{
		Id:           doc.Id,
		CreatedBy:    doc.CreatedBy,
		CreatedAt:    doc.CreatedAt,
		UpDatedAt:    doc.UpDatedAt,
		AssessorId:   doc.AssessorId,
		VerifierId:   doc.VerifierId,
		EvaluteField: doc.EvaluteField,
		Data:         []outgoing.Report{},
	}
	report, e := rC.getReportService.GetAllReport(doc.Id)
	if e != nil {
		return c.Output(http.StatusBadRequest, doc, errors.New("can not find with documentId"))
	}
	fmt.Println(report)
	// reportIds := []string{}
	for i, r := range report {
		outgoingGetDetailDocument.Data = append(outgoingGetDetailDocument.Data, outgoing.Report{
			Id:     r.Id,
			Field:  r.Field,
			Detail: []model.DetailReport{},
		})
		result, e := rC.getDetailReportService.GetDetailReport(r.Id)
		if e != nil {
			return c.Output(http.StatusBadRequest, doc, e)
		}
		outgoingGetDetailDocument.Data[i].Detail = result

	}

	return c.Output(http.StatusOK, outgoingGetDetailDocument, nil)
}

func (rC *reportController) SubmitReport(c *Context) error {
	var detailReport []model.DetailReport

	c.Bind(&detailReport)
	for _, i := range detailReport {
		e := rC.editDetailReportService.EditDetailReport(i)
		if e != nil {
			return c.Output(http.StatusBadRequest, nil, e)
		}
	}
	return c.Output(http.StatusOK, detailReport, nil)
}

func (rC *reportController) GetAllPIbySOId(c *Context) error {
	doc, e := rC.getDocumentService.GetAllPIbySOId(c.QueryParam("sOId"))
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, doc, nil)
}

func (rC *reportController) DeleteDocument(c *Context) error {
	if c.AuthObject.GetUserRole() != 0 {
		return c.Output(http.StatusForbidden, "dont have permission", errors.New("you are not admin role"))
	}
	so := model.SODocument{}
	c.Bind(&so)
	fmt.Println(so.Id)
	e := rC.createReportService.DeleteDocument(so.Id)
	if e != nil {
		return c.Output(http.StatusBadRequest, nil, e)
	}
	return c.Output(http.StatusOK, nil, e)
}
