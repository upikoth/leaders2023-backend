package v1

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

// GetCreativeSpaces godoc
// @Summary      Возвращает список креативных площадок
// @Accept       json
// @Produce      json
// @Param        MyAuthorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getCreativeSpacesResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creativeSpaces [get].
func (h *HandlerV1) GetCreativeSpaces(c *gin.Context) {
	creativeSpaces, err := h.store.GetCreativeSpaces()

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpacesGetDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.GetCreativeSpacesResponseFromStoreData(creativeSpaces)
	c.Set("responseData", responseData)
}

// GetCreativeSpace godoc
// @Summary      Возвращает информацию о пользователе
// @Produce      json
// @Param        id  path  string  true  "Id креативной площадки"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getCreativeSpaceResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creativeSpaces/:id [get].
func (h *HandlerV1) GetCreativeSpace(c *gin.Context) {
	reqData, err := requests.GetCreativeSpaceDataFromRequest(c)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpaceGetNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	creativeSpace, err := h.store.GetCreativeSpaceByID(reqData.ID)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpaceGetDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.GetCreativeSpaceResponseFromStoreData(creativeSpace)
	c.Set("responseData", responseData)
}

// CreateCreativeSpace godoc
// @Summary      Создание креативной площадки
// @Accept       json
// @Produce      json
// @Param        body body  requests.createCreativeSpaceRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createCreativeSpaceResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creativeSpace [post].
func (h *HandlerV1) CreateCreativeSpace(c *gin.Context) {
	reqData, err := requests.CreateCreativeSpaceDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleLandlord {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostForbidden)
		return
	}

	creativeSpaceCalendarEvents := []*store.CalendarEvent{}

	for _, event := range reqData.Calendar.Events {
		creativeSpaceCalendarEvents = append(creativeSpaceCalendarEvents, &store.CalendarEvent{
			Date: event.Date,
		})
	}

	creativeSpaceMetroStations := []*store.CreativeSpaceMetroStation{}

	for _, station := range reqData.MetroStations {
		creativeSpaceMetroStations = append(creativeSpaceMetroStations, &store.CreativeSpaceMetroStation{
			MetroStationID:    station.ID,
			DistanceInMinutes: station.DistanceInMinutes,
		})
	}

	creativeSpace := store.CreativeSpace{
		ID:                     uuid.New().String(),
		LandlordID:             userData.UserID,
		Title:                  reqData.Title,
		SpaceType:              reqData.SpaceType,
		Area:                   reqData.Area,
		Capacity:               reqData.Capacity,
		Address:                reqData.Address,
		Status:                 string(model.CreativeSpaceStatusConfirmationByAdmin),
		Photos:                 reqData.Photos,
		PricePerDay:            reqData.PricePerDay,
		Latitude:               reqData.Coordinate.Latitude,
		Longitude:              reqData.Coordinate.Longitude,
		Description:            reqData.Description,
		CalendarLink:           reqData.Calendar.Link,
		CalendarWorkDayIndexes: reqData.Calendar.WorkDayIndexes,
		CalendarEvents:         creativeSpaceCalendarEvents,
		MetroStations:          creativeSpaceMetroStations,
	}

	creativeSpaceID, storeErr := h.store.CreateCreativeSpace(creativeSpace)

	if storeErr != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePostDbError)
		c.Set("responseErrorDetails", storeErr)
		return
	}

	responseData := responses.CreateCreativeSpaceResponseFromStoreData(creativeSpaceID)
	c.Set("responseData", responseData)
}

// PatchCreativeSpace godoc
// @Summary      Обновление информации о креативном пространстве
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id креативного пространства"
// @Param        body body  requests.patchCreativeSpaceRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creativeSpaces/:id [patch].
func (h *HandlerV1) PatchCreativeSpace(c *gin.Context) {
	reqData, err := requests.PatchCreativeSpaceDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpacePatchNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	creativeSpaceToUpdate, err := h.store.GetCreativeSpaceByID(reqData.ID)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpacePatchDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserID != creativeSpaceToUpdate.LandlordID && userData.UserRole != model.RoleAdmin {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrPatchSpacePostForbidden)
		return
	}

	creativeSpaceCalendarEvents := []*store.CalendarEvent{}

	for _, event := range reqData.Calendar.Events {
		creativeSpaceCalendarEvents = append(creativeSpaceCalendarEvents, &store.CalendarEvent{
			Date:            event.Date,
			CreativeSpaceID: reqData.ID,
		})
	}

	creativeSpaceMetroStations := []*store.CreativeSpaceMetroStation{}

	for _, station := range reqData.MetroStations {
		creativeSpaceMetroStations = append(creativeSpaceMetroStations, &store.CreativeSpaceMetroStation{
			CreativeSpaceID:   reqData.ID,
			MetroStationID:    station.ID,
			DistanceInMinutes: station.DistanceInMinutes,
		})
	}

	creativeSpace := store.CreativeSpace{
		ID:                     reqData.ID,
		SpaceType:              reqData.SpaceType,
		Area:                   reqData.Area,
		Capacity:               reqData.Capacity,
		Title:                  reqData.Title,
		Address:                reqData.Address,
		Status:                 string(reqData.Status),
		Photos:                 reqData.Photos,
		PricePerDay:            reqData.PricePerDay,
		Latitude:               reqData.Coordinate.Latitude,
		Longitude:              reqData.Coordinate.Longitude,
		Description:            reqData.Description,
		CalendarLink:           reqData.Calendar.Link,
		CalendarWorkDayIndexes: reqData.Calendar.WorkDayIndexes,
		MetroStations:          creativeSpaceMetroStations,
	}

	if reqData.Calendar.Events != nil {
		creativeSpace.CalendarEvents = creativeSpaceCalendarEvents
	} else {
		creativeSpace.CalendarEvents = creativeSpaceToUpdate.CalendarEvents
	}

	storeErr := h.store.PatchCreativeSpace(creativeSpace)

	if storeErr != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpacePatchDbError)
		c.Set("responseErrorDetails", storeErr)
		return
	}
}

// DeleteCreativeSpace godoc
// @Summary      Удаление информации о креативной площадке
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id креативной площадки"
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/creativeSpaces/:id [delete].
func (h *HandlerV1) DeleteCreativeSpace(c *gin.Context) {
	reqData, err := requests.DeleteCreativeSpaceDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrCreativeSpaceDeleteNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	creativeSpaceToDelete, err := h.store.GetCreativeSpaceByID(reqData.ID)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpaceDeleteDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserID != creativeSpaceToDelete.LandlordID && userData.UserRole != model.RoleAdmin {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrPatchSpacePostForbidden)
		return
	}

	err = h.store.DeleteCreativeSpace(reqData.ID)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrCreativeSpaceDeleteDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	photos := []*s3.ObjectIdentifier{}

	for _, photo := range strings.Split(creativeSpaceToDelete.Photos, ";") {
		photos = append(photos, &s3.ObjectIdentifier{
			Key: aws.String(photo),
		})
	}

	//nolint:errcheck //Если ошибка при удалении с s3, ничего не делаем.
	h.s3.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(constants.FilesBucketName),
		Delete: &s3.Delete{
			Objects: photos,
		},
	})
}
