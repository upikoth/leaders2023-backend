package v1

import (
	"context"

	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
)

// GetAddresses godoc
// @Summary      Возвращает список пользователей
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getAddressesResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/addresses [get].
func (h *HandlerV1) GetAddresses(c *gin.Context) {
	reqData, err := requests.GetAddressesDataFromRequest(c)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrAddressesGetNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	addresses, err := h.dadataSuggestApi.Address(context.Background(), &suggest.RequestParams{
		Query: reqData.Query,
	})

	if err != nil {
		c.Set("responseErrorCode", constants.ErrAddressesGetNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	responseData := responses.GetAddressesResponseFromStoreData(addresses)
	c.Set("responseData", responseData)
}
