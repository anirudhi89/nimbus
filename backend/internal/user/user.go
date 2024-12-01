package user

import (
	t "github.com/anirudhi89/nimbus/internal/tenant"
	"github.com/labstack/gommon/email"
)

// User model and methods (e.g., user registration)

type User struct {
	name        string
	id          int
	email       email.Email
	role        string
	companyName t.Tenant
}

func GetIDTestCall() {
	return
}
