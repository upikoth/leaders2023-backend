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

// CreateCreativeSpace godoc
// @Summary      Создание креативной площадки
// @Accept       json
// @Produce      json
// @Param        body body  requests.createCreativeSpaceRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createCreativeSpaceResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creative-space [post].
func (h *HandlerV1) CreateCreativeSpace(c *gin.Context) {
	reqData, err := requests.CreateCreativeSpaceDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostNotValidRequestData)
	}

	creativeSpace := store.CreativeSpace{
		LandlordId:          userData.UserId,
		Photos:              reqData.Photos,
		PricePerHour:        reqData.PricePerHour,
		Latitude:            reqData.Coordinate.Latitude,
		Longitude:           reqData.Coordinate.Longitude,
		WorkingHoursStartAt: reqData.WorkingHours.StartAt,
		WorkingHoursEndAt:   reqData.WorkingHours.EndAt,
		Description:         reqData.Description,
	}

	creativeSpaceMetroStations := []store.CreativeSpaceMetroStation{}

	for _, station := range reqData.MetroStations {
		creativeSpaceMetroStations = append(creativeSpaceMetroStations, store.CreativeSpaceMetroStation{
			MetroStationId:    station.Id,
			DistanceInMinutes: station.DistanceInMinutes,
		})
	}

	creativeSpaceId, storeErr := h.store.CreateCreativeSpace(creativeSpace, creativeSpaceMetroStations)

	if storeErr != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostDbError)
	}

	responseData := responses.CreateCreativeSpaceResponseFromStoreData(creativeSpaceId)
	c.Set("responseData", responseData)
}
