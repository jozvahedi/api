package routing

import (
	"api/Utility/middleware"
	"api/config"
	"api/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

/*var DefJwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(security.JwtClime)
	},
	SigningKey: []byte("JozVahedi"),
}*/

func SetRouting(e *echo.Echo) error {
	AccountController(e)
	Users(e)
	Groups(e)
	Permissions(e)
	return nil
}
func AccountController(e *echo.Echo) {
	accountController := controller.NewAccountController()
	AC := e.Group("/login")
	AC.POST("", accountController.LoginUsers)
}
func Users(e *echo.Echo) {
	userController := controller.NewUserController()
	users := e.Group("/users/")
	users.Use(echojwt.WithConfig(config.AppConfig.DefJwtConfig))
	users.GET("getList", userController.GetUsersList, middleware.PermissionChecker("/users/getList"))
	users.GET("get/:id", userController.GetUserById)
	users.POST("create", userController.InsertUser)
	users.PUT("edit/:id", userController.UpdateUserById)
	users.DELETE("delete/:id", userController.DeleteUserById)
	users.POST("uploadAvatar", userController.UploadAvatar)
}
func Groups(e *echo.Echo) {
	Groups := e.Group("/groups/")
	Groups.Use(echojwt.WithConfig(config.AppConfig.DefJwtConfig))
	Groups.GET("getList", controller.GetGroupList)
	Groups.POST("create", controller.CreateGroup)
}
func Permissions(e *echo.Echo) {
	perm := e.Group("/permission/")
	perm.Use(echojwt.WithConfig(config.AppConfig.DefJwtConfig))
	perm.GET("getList", controller.GetPermissionList)
	perm.POST("create", controller.CreatePermission)
}

