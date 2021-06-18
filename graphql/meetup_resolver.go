package graphql

import (
	"context"
	"github.com/shuza/meetmeup/models"
)

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return r.UsersRepo.GetUserByID(obj.UserID)
}

func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }
