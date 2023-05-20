package requests

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type createFileRequestData struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func CreateFileDataFromRequest(c *gin.Context) (createFileRequestData, error) {
	data := createFileRequestData{}

	err := c.Bind(&data)

	if err != nil {
		return createFileRequestData{}, err
	}

	return data, nil
}

type deleteFileRequestData struct {
	FileName string `uri:"fileName" binding:"required"`
}

func DeleteFileDataFromRequest(c *gin.Context) (deleteFileRequestData, error) {
	data := deleteFileRequestData{}

	err := c.BindUri(&data)

	if err != nil {
		return deleteFileRequestData{}, err
	}

	return data, nil
}
