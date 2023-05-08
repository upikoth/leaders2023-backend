package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
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
	}

	user, err := h.store.GetUserByEmail(reqData.Email)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(reqData.Password))

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostUserNotExist)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
	})

	jwtToken, err := token.SignedString(h.jwtSecret)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrSessionPostCreateJwtToken)
		return
	}

	responseData := responses.CreateSessionResponseFromStoreData(user)
	c.Set("responseData", responseData)
	c.SetCookie("Authorization", jwtToken, int(constants.Month/time.Second), "", "", true, true)
}

// DeleteSession godoc
// @Summary      Удаление сессии
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/session [delete].
func (h *HandlerV1) DeleteSession(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", true, true)
}

// GetSession godoc
// @Summary      Получение сессии
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/session [get].
func (h *HandlerV1) GetSession(_ *gin.Context) {}
