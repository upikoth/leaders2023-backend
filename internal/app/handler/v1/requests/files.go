package requests

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type createFileRequestData struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func CreateFilesDataFromRequest(c *gin.Context) (createFileRequestData, error) {
	data := createFileRequestData{}

	err := c.Bind(&data)

	if err != nil {
		return createFileRequestData{}, err
	}

	return data, nil
}
