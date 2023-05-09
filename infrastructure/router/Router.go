package router

/*
create by: Hoangnd
create at: 2023-01-01
des: Xử lý router & authen
*/

import (
	aAuth "aBet/adapters/auth"
	"aBet/adapters/controller"
	"aBet/crypt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	var authObject aAuth.AuthObject
	// crypt.CreateAndSaveKeyPair()
	config := getMiddleWareConfig(&authObject)
	// e.Use(middleware.JWTWithConfig(config))
	group := e.Group("")
	// group.Use(middleware.CORS())
	// group.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowHeaders:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
	// 	AllowCredentials: true,
	// }))
	group.Use(middleware.JWTWithConfig(config))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ct := &controller.Context{
				Context: c,
			}
			return next(ct.Context)
		}
	})

	crypt.CreateAndSaveKeyPairV2()
	e.Static("/", "static/index.html")
	group.GET("/testJWT", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.TestJWT)
	})
	e.GET("/users-account/get-detail", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.GetDetailUsers)
	})
	e.POST("/users-account/login", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.LoginUserAccount)
	})
	e.POST("/users-account/create", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.AddUsers)
	})
	group.POST("/users-account/update", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.UpdateUsers)
	})
	e.DELETE("/users-account/delete", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.DeleteUsers)
	})
	e.POST("/users-account/reset-password", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.ResetPassword)
	})

	//__________________________________________________
	group.POST("document/create-report", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.CreateNewReport)
	})

	group.GET("document/get-document/:documentId", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.GetDetailDocument)
	})

	group.GET("document/get-document", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.GetAllDocument)
	})
	group.GET("document/get-document-by-user", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.GetAllDocumentById)
	})

	group.PUT("document/submit-report", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.SubmitReport)
	})
	group.GET("document/get-all-document-by-soId/:id", func(context echo.Context) error {
		return forward(context, authObject, c.ReportController.GetAllPIbySOId)
	})
	group.GET("SO/get-detail-so/:id", func(context echo.Context) error {
		return forward(context, authObject, c.SOController.GetDetailSODocument)
	})
	group.DELETE("SO/delete-SO", func(context echo.Context) error {
		return forward(context, authObject, c.SOController.DeleteSODocument)
	})

	group.POST("SO/create-so", func(context echo.Context) error {
		return forward(context, authObject, c.SOController.CreateNewSODocument)
	})

	group.GET("SO/get-all-so", func(context echo.Context) error {
		return forward(context, authObject, c.SOController.GetAllSODocument)
	})

	return e
}
func forward(context echo.Context, authObject aAuth.AuthObject, f func(*controller.Context) error) error {
	ct := &controller.Context{}
	ct.Context = context
	ct.AuthObject = authObject
	return f(ct)
}
