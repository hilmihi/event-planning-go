package router

import (
	"context"
	"database/sql"
	addmiddleware "sirclo/api/addMiddleware"
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

	e.Pre(middleware.RemoveTrailingSlash())

	//User
	UserRepository := repository.NewRepositoryUser(db)
	UserService := service.NewUserService(UserRepository)

	//Comment
	CommentRepository := repository.NewRepositoryComment(db)
	CommentService := service.NewCommentService(CommentRepository)

	//Participant
	ParticipantRepository := repository.NewRepositoryParticipant(db)
	ParticipantService := service.NewParticipantService(ParticipantRepository)

	// Event
	EventRepository := repository.NewRepositoryEvent(db)
	EventService := service.NewEventService(EventRepository)

	client := _graph.NewResolver(UserService, CommentService, ParticipantService, EventService)
	srv := NewGraphQLServer(client)

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
