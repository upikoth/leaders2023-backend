package store

type CalendarEvent struct {
	tableName       struct{} `pg:"calendar_events"` //nolint:unused // Имя таблицы
	Id              int      `pg:"id"`
	Date            string   `pg:"date"`
	CreativeSpaceId int      `pg:"creative_space_id"`
}
