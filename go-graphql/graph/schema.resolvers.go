package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-graphql/graph/loader"
	"go-graphql/graph/model"
	"math/big"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))

	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", randNumber),
		Text:   input.Text,
		UserID: input.UserID,
		//		Done: false,
	}

	// dependency, memory as db
	r.todos = append(r.todos, todo)

	return todo, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserStorage.Put(input)
	return user, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user, err := loader.GetUser(ctx, obj.UserID)

	//	user, err := r.UserStorage.Get(obj.UserID)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// Name2 is the resolver for the name2 field.
func (r *userResolver) Name2(ctx context.Context, obj *model.User) (string, error) {
	return obj.Name + "2", nil
}

// NameFull is the resolver for the nameFull field.
func (r *userResolver) NameFull(ctx context.Context, obj *model.User) (string, error) {
	return obj.FullName(), nil
}

// Todo is the resolver for the todo field.
func (r *userResolver) Todo(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	todos, err := r.GetTodosByUserId(obj)
	if err != nil {
		return nil, err
	}
	return todos, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
