package model

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleTenant   Role = "tenant"
	RoleLandlord Role = "landlord"
)

type BookingStaus string

const (
	BookingStausConfirmationByLandlord BookingStaus = "confirmation"
	BookingStausConfirmedByLandlord    BookingStaus = "confirmed"
	BookingStausFinishedByTenant       BookingStaus = "finished"
)
