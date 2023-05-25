package store

type CreativeSpaceMetroStation struct {
	tableName         struct{}      `pg:"creative_space_metro_station"` //nolint:unused // Имя таблицы
	MetroStationId    int           `pg:"metro_station_id"`
	CreativeSpaceId   int           `pg:"creative_space_id"`
	DistanceInMinutes int           `pg:"distance_in_minutes"`
	MetroStation      *MetroStation `pg:"rel:has-one"`
}
