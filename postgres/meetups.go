package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/shuza/meetmeup/models"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (r *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	if err := r.DB.Model(&meetups).Select(); err != nil {
		return nil, err
	}
	return meetups, nil
}

func (r *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := r.DB.Model(meetup).Returning("*").Insert()
	return meetup, err
}
