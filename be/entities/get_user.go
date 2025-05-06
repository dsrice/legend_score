package entities

import (
	"legend_score/entities/db"
)

type GetUserEntity struct {
	Code   string
	User   db.UserEntity
	UserID int
}