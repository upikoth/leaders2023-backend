package store

type CreativeSpaceMetroStation struct {
	MetroStationID    string `gorm:"primarykey"`
	CreativeSpaceID   string
	DistanceInMinutes int
	MetroStation      *MetroStation `gorm:"foreignKey:ID;references:MetroStationID"`
}
