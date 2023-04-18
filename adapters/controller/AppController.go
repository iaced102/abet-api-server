package controller

import (
	"net/http"
)

type AppController struct {
	AuthController   AuthController
	ReportController ReportController
}

func (app *AppController) HelthCheck(e *Context) error {
	return e.Output(http.StatusOK, map[string]interface{}{"status": "Running"}, nil)
}
