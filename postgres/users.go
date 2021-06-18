package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/shuza/meetmeup/models"
)

type UsersRepo struct {
	DB *pg.DB
}

func (r *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Model(&user).Where("id = ?", id).First(); err != nil {
		return nil, err
	}
	return &user, nil
}
