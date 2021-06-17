package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/shuza/meetmeup/graph/generated"
	"github.com/shuza/meetmeup/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
