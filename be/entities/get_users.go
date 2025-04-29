package entities

import (
    "legend_score/entities/db"
)

type GetUsersEntity struct {
    Code   string
    Users  []db.UserEntity
    UserID *int
    LoginID *string
    Name    *string
}

func (e *GetUsersEntity) SetFilters(userID *int, loginID *string, name *string) {
    e.UserID = userID
    e.LoginID = loginID
    e.Name = name
}