package store

import "time"

type Event struct {
	tableName       struct{}  `pg:"events"` //nolint:unused // Имя таблицы
	Id              int       `pg:"id"`
	CreativeSpaceId int       `pg:"creative_space_id"`
	StartAt         time.Time `pg:"start_at_date"`
	EndAt           time.Time `pg:"end_at_date"`
}
