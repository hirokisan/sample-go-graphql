package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/hirokisan/sample-go-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver :
type Resolver struct {
	todos []*model.Todo
}
