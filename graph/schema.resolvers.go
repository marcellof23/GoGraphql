package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marcellof23/GoGraphql/graph/generated"
	"github.com/marcellof23/GoGraphql/graph/model"
	"github.com/marcellof23/GoGraphql/internal/handlers"
	"github.com/marcellof23/GoGraphql/internal/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*models.User, error) {
	return handlers.CreateUserHandler(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, input *model.NewUser) (*models.User, error) {
	return handlers.UpdateUserHandler(ctx, id, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (bool, error) {
	return handlers.DeleteUserHandler(ctx, id)
}

func (r *queryResolver) User(ctx context.Context) ([]*models.User, error) {
	return handlers.GetAllUserHandler(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
