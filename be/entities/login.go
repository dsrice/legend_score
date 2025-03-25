package entities

import (
	"legend_score/entities/db"
)

type LoginEntity struct {
	LoginID  string
	Password string

	Code string

	User db.UserEntity
}