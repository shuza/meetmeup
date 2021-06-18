package graphql

import (
	"context"
	"github.com/shuza/meetmeup/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }
