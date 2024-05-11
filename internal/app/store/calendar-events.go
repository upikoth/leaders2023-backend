package store

type CalendarEvent struct {
	ID              string `gorm:"primarykey"`
	Date            string
	CreativeSpaceID string
	BookingID       string
}
