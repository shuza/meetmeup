package graphql

import (
	"context"
	"github.com/shuza/meetmeup/models"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() UserResolver { return &userResolver{r} }

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup
	for _, v := range meetups {
		if v.UserID == obj.ID {
			m = append(m, v)
		}
	}
	return m, nil
}
