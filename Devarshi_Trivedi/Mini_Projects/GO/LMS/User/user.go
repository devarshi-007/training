package user

import "time"

type User struct {
	UserId             int
	Name               string
	CurrentBookHolding int
	UserValidity       time.Time
}
