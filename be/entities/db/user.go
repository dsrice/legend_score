package db

import (
	"legend_score/infra/database/models"
	"time"
)

type UserEntity struct {
	ID             int
	LoginID        string
	Name           string
	Password       string
	ChangePassFlag bool
	ErrorCount     int
	ErrorDateTime  *time.Time
	LockDateTime   *time.Time
}

func (u *UserEntity) SetEntity(user *models.User) {
	u.ID = user.ID
	u.LoginID = user.LoginID
	u.Name = user.Name
	u.Password = user.Password
	u.ChangePassFlag = user.ChangePassFlag
	u.ErrorCount = user.ErrorCount
	if user.ErrorDatetime.Valid {
		u.ErrorDateTime = &user.ErrorDatetime.Time
	}

	if user.LockDatetime.Valid {
		u.LockDateTime = &user.LockDatetime.Time
	}
}