package model

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleRenter   Role = "renter"
	RoleLandlord Role = "landlord"
)
