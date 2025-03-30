package entities

import "legend_score/controllers/request"

type CreateUserEntity struct {
	LoginID  string
	Name     string
	Password string

	Code string
}

func (e *CreateUserEntity) SetEntity(req *request.CreateUserRequest) {
	e.LoginID = req.LoginID
	e.Name = req.Name
	e.Password = req.Password
}