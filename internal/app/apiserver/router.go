package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/upikoth/leaders2023-backend/docs"
	"github.com/upikoth/leaders2023-backend/internal/app/constants"
	"github.com/upikoth/leaders2023-backend/internal/app/model"
)

func (s *APIServer) initRoutes() {
	docs.SwaggerInfo.Schemes = []string{}
	s.router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	proxiesErr := s.router.SetTrustedProxies(nil)

	if proxiesErr != nil {
		log.Println(proxiesErr)
	}

	s.router.Use(CORSMiddleware())
	s.router.Use(formatResponse())

	s.router.GET("/api/v1/health", s.handler.V1.CheckHealth)

	s.router.POST("/api/v1/session", s.handler.V1.CreateSession)

	authorized := s.router.Use(checkAuthorization(s.config.JwtSecret))

	authorized.GET("/api/v1/session", s.handler.V1.GetSession)

	authorized.GET("/api/v1/users", s.handler.V1.GetUsers)

	authorized.GET("/api/v1/users/:id", s.handler.V1.GetUser)
	authorized.POST("/api/v1/user", s.handler.V1.CreateUser)
	authorized.PATCH("/api/v1/users/:id", s.handler.V1.PatchUser)
	authorized.DELETE("/api/v1/users/:id", s.handler.V1.DeleteUser)

	authorized.GET("/api/v1/metroStations", s.handler.V1.GetMetroStations)

	authorized.GET("/api/v1/creativeSpaces", s.handler.V1.GetCreativeSpaces)

	authorized.GET("/api/v1/creativeSpaces/:id", s.handler.V1.GetCreativeSpace)
	authorized.POST("/api/v1/creativeSpace", s.handler.V1.CreateCreativeSpace)
	authorized.PATCH("/api/v1/creativeSpaces/:id", s.handler.V1.PatchCreativeSpace)
	authorized.DELETE("/api/v1/creativeSpaces/:id", s.handler.V1.DeleteCreativeSpace)

	authorized.GET("/api/v1/addresses", s.handler.V1.GetAddresses)

	authorized.POST("/api/v1/file", s.handler.V1.CreateFile)
	authorized.DELETE("/api/v1/files/:fileName", s.handler.V1.DeleteFile)

	authorized.POST("/api/v1/calendar/convert", s.handler.V1.ConvertCalendar)
	authorized.POST("/api/v1/calendar/convertFromLink", s.handler.V1.ConvertCalendarFromLink)

	authorized.GET("/api/v1/bookings", s.handler.V1.GetBookings)

	authorized.GET("/api/v1/bookings/:id", s.handler.V1.GetBooking)
	authorized.POST("/api/v1/booking", s.handler.V1.CreateBooking)
	authorized.PATCH("/api/v1/bookings/:id", s.handler.V1.PatchBooking)
	authorized.DELETE("/api/v1/bookings/:id", s.handler.V1.DeleteBooking)

	authorized.POST("/api/v1/score", s.handler.V1.CreateScore)

	s.router.NoRoute(func(c *gin.Context) {
		c.Set("responseCode", http.StatusNotFound)
		c.Set("responseErrorCode", constants.ErrRouteNotFound)
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, MyAuthorization",
		)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func formatResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		code, isCodeExist := c.Get("responseCode")
		data, isDataExist := c.Get("responseData")
		errorCode, isErrorCodeExist := c.Get("responseErrorCode")
		errorDetails, isErrorDetailsExist := c.Get("responseErrorDetails")
		description := ""

		if isErrorCodeExist {
			description = constants.ErrDescriptionByCode[errorCode.(error)]
		}

		if !isErrorDetailsExist {
			errorDetails = ""
		}

		_, err := strconv.Atoi(fmt.Sprintf("%v", errorDetails))

		if err == nil && errorDetails != "" {
			errorCode, _ = errorDetails.(error)
			description = constants.ErrDescriptionByCode[errorDetails.(error)]
			errorDetails = ""
		}

		if !isCodeExist {
			code = http.StatusOK
		}

		if !isDataExist {
			data = map[string]string{}
		}

		if isErrorCodeExist {
			response := model.ResponseError{}
			response.Success = false
			response.Error = &model.ResponseErrorField{
				Code:        fmt.Sprintf("%v", errorCode),
				Description: description,
				Details:     fmt.Sprintf("%v", errorDetails),
			}
			c.JSON(code.(int), response)
		} else {
			response := model.ResponseSuccess{}
			response.Success = true
			response.Data = data
			c.JSON(code.(int), response)
		}
	}
}

func checkAuthorization(jwtSecret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.Request.Header.Get("MyAuthorization")

		if jwtToken == "" {
			c.Set("responseCode", http.StatusUnauthorized)
			c.Set("responseErrorCode", constants.ErrUserNotAuthorized)
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			jwtToken, &model.JwtTokenClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return jwtSecret, nil
			})

		claims, isClaimsValid := token.Claims.(*model.JwtTokenClaims)

		if err != nil || !isClaimsValid || !token.Valid {
			c.Set("responseCode", http.StatusUnauthorized)
			c.Set("responseErrorCode", constants.ErrUserNotAuthorized)
			c.Abort()
			return
		}

		c.Set("userData", claims.UserData)
		c.Next()
	}
}
