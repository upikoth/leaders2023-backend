package responses

import (
	ical "github.com/arran4/golang-ical"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type convertCaledarResponseEvent struct {
	Date string `json:"date"`
}

type convertCalendarResponseData struct {
	Events []convertCaledarResponseEvent `json:"events"`
}

func ConvertCalendarResponseFromCalendarEvents(events []*ical.VEvent) convertCalendarResponseData {
	res := convertCalendarResponseData{
		Events: []convertCaledarResponseEvent{},
	}

	for _, event := range events {
		startAt, _ := event.GetStartAt()
		endAt, _ := event.GetEndAt()

		for startAt.Before(endAt) {
			res.Events = append(res.Events, convertCaledarResponseEvent{
				Date: startAt.Format(constants.DateLayout),
			})

			startAt = startAt.Add(constants.Day)
		}
	}

	return res
}

type convertCaledarFromLinkResponseEvent struct {
	Date string `json:"date"`
}

type convertCalendarFromLinkResponseData struct {
	Events []convertCaledarFromLinkResponseEvent `json:"events"`
}

func ConvertCalendarFromLinkResponseFromCalendarEvents(events []*ical.VEvent) convertCalendarFromLinkResponseData {
	res := convertCalendarFromLinkResponseData{
		Events: []convertCaledarFromLinkResponseEvent{},
	}

	for _, event := range events {
		startAt, _ := event.GetStartAt()
		endAt, _ := event.GetEndAt()

		for startAt.Before(endAt) {
			res.Events = append(res.Events, convertCaledarFromLinkResponseEvent{
				Date: startAt.Format(constants.DateLayout),
			})

			startAt = startAt.Add(constants.Day)
		}
	}

	return res
}
