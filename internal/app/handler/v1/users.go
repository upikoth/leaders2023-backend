package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
	"github.com/upikoth/leaders2023-backend/internal/app/store"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers godoc
// @Summary      Возвращает список пользователей
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getUsersResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/users [get].
func (h *HandlerV1) GetUsers(c *gin.Context) {
	users, err := h.store.GetUsers()
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUsersGetDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	if !isClaimsValid || userData.UserRole != model.RoleAdmin {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrUsersGetForbidden)
		return
	}

	responseData := responses.GetUsersResponseFromStoreData(users)
	c.Set("responseData", responseData)
}

// GetUser godoc
// @Summary      Возвращает информацию о пользователе
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getUserResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/users/:id [get].
func (h *HandlerV1) GetUser(c *gin.Context) {
	reqData, err := requests.GetUserDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserGetNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleAdmin && userData.UserId != reqData.Id {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrUserGetForbidden)
		return
	}

	user, err := h.store.GetUserById(reqData.Id)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserGetDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.GetUserResponseFromStoreData(user)
	c.Set("responseData", responseData)
}

// CreateUser godoc
// @Summary      Создание пользователя
// @Accept       json
// @Produce      json
// @Param        body body  requests.createUserRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createUserResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/user [post].
func (h *HandlerV1) CreateUser(c *gin.Context) {
	reqData, err := requests.CreateUserDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleAdmin {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrUserPostForbidden)
		return
	}

	saltedBytes := []byte(reqData.Password)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserPostCreateHash)
		c.Set("responseErrorDetails", err)
		return
	}

	user := store.User{
		Phone:        reqData.Phone,
		Role:         reqData.Role,
		PasswordHash: string(hashedBytes),
	}

	createdUserId, err := h.store.CreateUser(user)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserPostDbError)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.CreateUserResponseFromStoreData(createdUserId)
	c.Set("responseData", responseData)
}

// PatchUser godoc
// @Summary      Обновление информации о пользователе
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Param        body body  requests.patchUserRequestData true "Параметры запроса"
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/users/:id [patch].
func (h *HandlerV1) PatchUser(c *gin.Context) {
	reqData, err := requests.PatchUserDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserPatchNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleAdmin && userData.UserId != reqData.Id {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrUserPatchForbidden)
		return
	}

	user := store.User{
		Id:    reqData.Id,
		Phone: reqData.Phone,
	}

	storeErr := h.store.PatchUser(user)

	if storeErr != nil {
		c.Set("responseErrorCode", constants.ErrUserPatchDbError)
		c.Set("responseErrorDetails", storeErr)
		return
	}
}

// DeleteUser godoc
// @Summary      Удаление информации о пользователе
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "Id пользователя"
// @Success      200  {object}  model.ResponseSuccess
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/users/:id [delete].
func (h *HandlerV1) DeleteUser(c *gin.Context) {
	reqData, err := requests.DeleteUserDataFromRequest(c)
	userData, isClaimsValid := c.MustGet("userData").(model.JwtTokenUserData)

	if err != nil || !isClaimsValid {
		c.Set("responseCode", http.StatusBadRequest)
		c.Set("responseErrorCode", constants.ErrUserDeleteNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	if userData.UserRole != model.RoleAdmin {
		c.Set("responseCode", http.StatusForbidden)
		c.Set("responseErrorCode", constants.ErrUserDeleteForbidden)
		return
	}

	err = h.store.DeleteUser(reqData.Id)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrUserDeleteDbError)
		c.Set("responseErrorDetails", err)
		return
	}
}
