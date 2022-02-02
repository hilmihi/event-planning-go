package router

import (
	"context"
	"database/sql"
	addmiddleware "sirclo/api/addMiddleware"
	"sirclo/api/controller"
	"sirclo/api/repository"
	"sirclo/api/service"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_graph "sirclo/api/graph"
)

func InitRoute(db *sql.DB) *echo.Echo {
	//new echo
	e := echo.New()
	//setting cors
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// e.Pre(middleware.RemoveTrailingSlash(), middleware.Logger())
	authService := addmiddleware.AuthService()

	//User
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewUserService(UserRepository)
	UserController := controller.NewUserController(authService, UserService)

	client := _graph.NewResolver(UserService)
	srv := NewGraphQLServer(client)

	e.POST("/login", UserController.AuthUserController)
	e.GET("/users", addmiddleware.AuthMiddleware(authService, UserService, UserController.GetUsersController))
	e.GET("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.GetUserController))
	e.GET("/users/myprofile", addmiddleware.AuthMiddleware(authService, UserService, UserController.GetMyUserController))
	e.POST("/users", UserController.CreateUserController)
	e.PUT("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.UpdateUserController))
	e.DELETE("/users/:id", addmiddleware.AuthMiddleware(authService, UserService, UserController.DeleteUserController))

	{
		e.POST("/query", func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c.Get("jwt_user_id"))
			c.SetRequest(c.Request().WithContext(ctx))
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		}, addmiddleware.JWTMiddlewareGraphQL())

		//for subscriptions
		// e.

		e.POST("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	return e
}
