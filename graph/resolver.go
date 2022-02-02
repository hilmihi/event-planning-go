package graph

import (
	"context"
	"sirclo/api/service"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// db      *gorm.DB
	userService service.ServiceUser
	// tmpList  []*_graphModel.User
	// Observer map[string]chan []*model.Person
	// Observer map[string]chan *_graphModel.User
	// mu sync.Mutex
}

func NewResolver(ur service.ServiceUser) *Resolver {
	return &Resolver{
		userService: ur,
		// tmpList:  []*_graphModel.User{},
		// Observer: map[string]chan *_graphModel.User{},
		// mu:       sync.Mutex{},
	}
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
