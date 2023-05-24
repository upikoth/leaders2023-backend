package responses

import (
	ical "github.com/arran4/golang-ical"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
)

type convertCaledarResponseEvent struct {
	Date string `json:"date"`
}

type convertCaledarResponseData struct {
	Events []convertCaledarResponseEvent `json:"events"`
}

func ConvertCalendarResponseFromCalendarEvents(events []*ical.VEvent) convertCaledarResponseData {
	res := convertCaledarResponseData{}

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

type convertCaledarFromLinkResponseData struct {
	Events []convertCaledarFromLinkResponseEvent `json:"events"`
}

func ConvertCalendarFromLinkResponseFromCalendarEvents(events []*ical.VEvent) convertCaledarFromLinkResponseData {
	res := convertCaledarFromLinkResponseData{}

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
