package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
)

// GetMetroStations godoc
// @Summary      Возвращает полный список станций метро
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.getMetroStationsResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/metro-stations [get].
func (h *HandlerV1) GetMetroStations(c *gin.Context) {
	metroStations, err := h.store.GetMetroStations()

	if err != nil {
		c.Set("responseErrorCode", err)
		return
	}

	responseData := responses.GetMetroStationsResponseFromStoreData(metroStations)
	c.Set("responseData", responseData)
}
