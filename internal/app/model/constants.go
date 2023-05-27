package model

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleTenant   Role = "tenant"
	RoleLandlord Role = "landlord"
)

type BookingStatus string

const (
	BookingStatusConfirmationByLandlord BookingStatus = "confirmation"
	BookingStatusConfirmedByLandlord    BookingStatus = "confirmed"
)

type CreativeSpaceStatus string

const (
	CreativeSpaceStatusConfirmationByAdmin CreativeSpaceStatus = "confirmation"
	CreativeSpaceStatusConfirmedByAdmin    CreativeSpaceStatus = "confirmed"
)
