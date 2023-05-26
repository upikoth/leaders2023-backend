package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

// CreateBooking godoc
// @Summary      Бронирование креативной площадки
// @Accept       json
// @Produce      json
// @Param        body body  requests.createBookingRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createBookingResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/booking [post].
func (h *HandlerV1) CreateBooking(c *gin.Context) {
	reqData, err := requests.CreateBookingDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrBookingPostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleTenant {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostForbidden)
		return
	}

	creativeSpace, err := h.store.GetCreativeSpaceById(reqData.CreativeSpaceId)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrBookingPostDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	bookingCalendarEvents := []*store.CalendarEvent{}

	for _, event := range reqData.CalendarEvents {
		bookingCalendarEvents = append(bookingCalendarEvents, &store.CalendarEvent{
			CreativeSpaceId: creativeSpace.Id,
			Date:            event.Date,
		})
	}

	booking := store.Booking{
		TenantId:        userData.UserId,
		LandlordId:      creativeSpace.LandlordId,
		CreativeSpaceId: creativeSpace.Id,
		Status:          model.BookingStausConfirmationByLandlord,
		FullPrice:       creativeSpace.PricePerDay * len(reqData.CalendarEvents),
		CalendarEvents:  bookingCalendarEvents,
	}

	bookingId, storeErr := h.store.CreateBooking(booking)

	if storeErr != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrBookingPostDbError)
		c.Set("responseErrorDetails", storeErr)
		return
	}

	responseData := responses.CreateBookingResponseFromStoreData(bookingId)
	c.Set("responseData", responseData)
}
