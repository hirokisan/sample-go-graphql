package graph

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hirokisan/sample-go-graphql/graph/generated"
	"github.com/hirokisan/sample-go-graphql/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	t.Run("", func(t *testing.T) {
		user := model.User{
			ID: "1",
		}
		todos := []*model.Todo{
			{
				ID:   "1",
				User: &user,
			},
			{
				ID:   "2",
				User: &user,
			},
		}
		resolvers := Resolver{todos: todos}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))

		var resp struct {
			Todos []model.Todo
		}
		q := `
    query findTodos {
      todos {
        id
      }
    }`
		c.MustPost(q, &resp)
		assert.Equal(t, len(todos), len(resp.Todos))
	})
}
