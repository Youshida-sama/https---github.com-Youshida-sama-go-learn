package requests

import (
	"main/models"
	"time"
)

type UserActionRequest struct {
	ID      int    `json:"id" validate:"required,gte=0"`
	Name    string `json:"name" validate:"req=Surname"`
	Surname string `json:"surname" validate:"req=Name"`
	Time    string `json:"time" validate:"isoTime"`
}

//время обязательное, только если есть имя и наоборот

func (uar UserActionRequest) Map() (ua *models.UserAction, err error) {
	t, err := time.Parse(time.RFC3339, uar.Time)

	if err != nil {
		return
	}

	ua = &models.UserAction{
		ID:      uar.ID,
		Name:    uar.Name,
		Surname: uar.Surname,
		Time:    t,
	}

	return
}
