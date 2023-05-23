package responses

import (
	"time"

	ical "github.com/arran4/golang-ical"
)

type convertCaledarResponseEvent struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

type convertCaledarResponseData struct {
	Events []convertCaledarResponseEvent `json:"events"`
}

func ConvertCalendarResponseFromCalendarEvents(events []*ical.VEvent) convertCaledarResponseData {
	res := convertCaledarResponseData{}

	for _, event := range events {
		startAt, _ := event.GetStartAt()
		endAt, _ := event.GetEndAt()

		res.Events = append(res.Events, convertCaledarResponseEvent{
			StartAt: startAt.UTC().Format(time.RFC1123),
			EndAt:   endAt.UTC().Format(time.RFC1123),
		})
	}

	return res
}
