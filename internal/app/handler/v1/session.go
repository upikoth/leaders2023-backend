package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

// CreateSession godoc
// @Summary      Создание сессии пользователя
// @Accept       json
// @Produce      json
// @Param        body body  requests.createSessionRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createSessionResponseData}
// @Router       /api/v1/session [post].
func (h *HandlerV1) CreateSession(c *gin.Context) {
	reqData, err := requests.CreateSessionDataFromRequest(c)

	if err != nil {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrSessionPostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	user, err := h.store.GetUserByPhone(reqData.Phone)

	if user.ID == "" || err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(reqData.Password))

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	tokenUserData := model.JwtTokenUserData{
		UserID:   user.ID,
		UserRole: model.Role(user.Role),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userData": tokenUserData,
	})

	jwtToken, err := token.SignedString(h.env.JwtSecret)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostCreateJwtToken)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.CreateSessionResponseFromStoreData(user, jwtToken)
	c.Set("responseData", responseData)
}

// GetSession godoc
// @Summary      Получение сессии
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/session [get].
func (h *HandlerV1) GetSession(c *gin.Context) {
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrSessionGetNotValidRequestData)
		return
	}

	user, err := h.store.GetUserByID(userData.UserID)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserGetDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.GetSessionResponseFromStoreData(user)
	c.Set("responseData", responseData)
}
