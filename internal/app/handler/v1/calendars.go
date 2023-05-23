package v1

import (
	ical "github.com/arran4/golang-ical"
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
)

// ConverCalendar godoc
// @Summary      Возвращает события календаря
// @Accept       mpfd
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.convertCaledarResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/calendar/convert [post].
func (h *HandlerV1) ConverCalendar(c *gin.Context) {
	reqData, err := requests.ConvertCalendarDataFromRequest(c)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCalendarConvertNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	file, errOpenFile := reqData.Calendar.Open()

	if errOpenFile != nil {
		c.Set("responseErrorCode", constants.ErrCalendarConvertNotValidRequestData)
		c.Set("responseErrorDetails", errOpenFile)
		return
	}

	calendar, errCalendarParse := ical.ParseCalendar(file)

	if errCalendarParse != nil {
		c.Set("responseErrorCode", constants.ErrCalendarConvertNotValidRequestData)
		c.Set("responseErrorDetails", errCalendarParse)
		return
	}

	events := calendar.Events()

	responseData := responses.ConvertCalendarResponseFromCalendarEvents(events)
	c.Set("responseData", responseData)
}
