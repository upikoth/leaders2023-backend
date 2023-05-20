package v1

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/requests"
	"github.com/upikoth/leaders2023-backend/internal/app/handler/v1/responses"
	"github.com/upikoth/leaders2023-backend/internal/app/utils"
)

// CreateFile godoc
// @Summary      Возвращает ссылку на файл
// @Accept       mpfd
// @Produce      json
// @Param        Authorization  header  string  true  "Authentication header"
// @Success      200  {object}  model.ResponseSuccess{data=responses.createFileStationsResponseData}
// @Failure      403  {object}  model.ResponseError "Коды ошибок: [1100]"
// @Router       /api/v1/file [post].
func (h *HandlerV1) CreateFile(c *gin.Context) {
	reqData, err := requests.CreateFilesDataFromRequest(c)

	if err != nil {
		c.Set("responseErrorCode", constants.ErrFilePostNotValidRequestData)
		c.Set("responseErrorDetails", err)
		return
	}

	f, fileOpenError := reqData.File.Open()

	if fileOpenError != nil {
		c.Set("responseErrorCode", constants.ErrFilePostOpenFileError)
		c.Set("responseErrorDetails", fileOpenError)
		return
	}

	defer f.Close()

	uuid := uuid.New().String()
	fileExtension := utils.GetFileExtensionByFileName(reqData.File.Filename)
	fileKey := fmt.Sprintf("%s.%s", uuid, fileExtension)

	_, s3Error := h.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(constants.FilesBucketName),
		Key:         aws.String(fileKey),
		ContentType: aws.String(reqData.File.Header.Get("Content-Type")),
		Body:        f,
	})

	if s3Error != nil {
		c.Set("responseErrorCode", constants.ErrFilePostS3Error)
		c.Set("responseErrorDetails", s3Error)
		return
	}

	responseData := responses.CreateFileResponseFromFileKey(fileKey)
	c.Set("responseData", responseData)
}
