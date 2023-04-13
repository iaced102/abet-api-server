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
	"fmt"

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
	e.POST("/user-account/create", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.AddUsers)
	})
	e.POST("/users-account/update", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.UpdateUsers)
	})
	e.DELETE("/users-account/delete", func(context echo.Context) error {
		return forward(context, authObject, c.AuthController.DeleteUsers)
	})

	return e
}
func forward(context echo.Context, authObject aAuth.AuthObject, f func(*controller.Context) error) error {
	ct := &controller.Context{}
	ct.Context = context
	ct.AuthObject = authObject
	fmt.Println(authObject)
	return f(ct)
}
