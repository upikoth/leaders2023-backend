package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
)

// CreateScore godoc
// @Summary      Бронирование креативной площадки
// @Accept       json
// @Produce      json
// @Param        body body  requests.createScoreRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createScoreResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/score [post].
func (h *HandlerV1) CreateScore(c *gin.Context) {
	reqData, err := requests.CreateScoreDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrScorePostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleTenant {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrScorePostForbidden)
		return
	}

	score := store.Score{
		ID:              uuid.New().String(),
		UserID:          userData.UserID,
		CreativeSpaceID: reqData.CreativeSpaceID,
		BookingID:       reqData.BookingID,
		Rating:          reqData.Rating,
		Comment:         reqData.Comment,
	}

	scoreID, storeErr := h.store.CreateScore(score)

	if storeErr != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrScorePostDbError)
		c.Set("responseErrorDetails", storeErr)
		return
	}

	responseData := responses.CreateScoreResponseFromStoreData(scoreID)
	c.Set("responseData", responseData)
}
