package constants

import (
	"errors"
)

var (
	ErrRouteNotFound     = errors.New("1000")
	ErrUserNotAuthorized = errors.New("1100")

	ErrUsersGetDbError   = errors.New("1200")
	ErrUsersGetForbidden = errors.New("1201")

	ErrUserGetNotValidRequestData = errors.New("1300")
	ErrUserGetNotFoundById        = errors.New("1301")
	ErrUserGetDbError             = errors.New("1302")
	ErrUserGetByPhoneDbError      = errors.New("1303")
	ErrUserGetByPhoneUserNotExist = errors.New("1304")
	ErrUserGetForbidden           = errors.New("1305")

	ErrUserPostNotValidRequestData = errors.New("1400")
	ErrUserPostPhoneExist          = errors.New("1401")
	ErrUserPostDbError             = errors.New("1402")
	ErrUserPostCreateHash          = errors.New("1403")
	ErrUserPostForbidden           = errors.New("1404")

	ErrUserDeleteNotValidRequestData = errors.New("1500")
	ErrUserDeleteNotFoundById        = errors.New("1501")
	ErrUserDeleteDbError             = errors.New("1502")
	ErrUserDeleteForbidden           = errors.New("1503")

	ErrUserPatchNotValidRequestData = errors.New("1600")
	ErrUserPatchPhoneExist          = errors.New("1601")
	ErrUserPatchDbError             = errors.New("1602")
	ErrUserPatchNotFoundById        = errors.New("1603")
	ErrUserPatchForbidden           = errors.New("1604")

	ErrSessionPostNotValidRequestData = errors.New("1700")
	ErrSessionPostUserNotExist        = errors.New("1701")
	ErrSessionPostCreateJwtToken      = errors.New("1702")

	ErrMetroStationsGetDbError = errors.New("1800")

	ErrCreativeSpacePostNotValidRequestData = errors.New("1900")
	ErrCreativeSpacePostDbError             = errors.New("1901")
	ErrCreativeSpacePostForbidden           = errors.New("1902")

	ErrCreativeSpacesGetDbError = errors.New("2000")

	ErrCreativeSpaceGetNotValidRequestData = errors.New("2100")
	ErrCreativeSpaceGetNotFoundById        = errors.New("2101")
	ErrCreativeSpaceGetDbError             = errors.New("2102")

	ErrCreativeSpaceDeleteNotValidRequestData = errors.New("2200")
	ErrCreativeSpaceDeleteNotFoundById        = errors.New("2201")
	ErrCreativeSpaceDeleteDbError             = errors.New("2202")
	ErrCreativeSpaceDeleteS3Error             = errors.New("2203")

	ErrCreativeSpacePatchNotValidRequestData = errors.New("2300")
	ErrCreativeSpacePatchNotFoundById        = errors.New("2301")
	ErrCreativeSpacePatchDbError             = errors.New("2302")
	ErrPatchSpacePostForbidden               = errors.New("2303")

	ErrSessionGetNotValidRequestData = errors.New("2400")

	ErrAddressesGetNotValidRequestData = errors.New("2500")
	ErrAddressesGetDadataError         = errors.New("2501")

	ErrFilePostNotValidRequestData = errors.New("2600")
	ErrFilePostOpenFileError       = errors.New("2601")
	ErrFilePostS3Error             = errors.New("2602")

	ErrFileDeleteNotValidRequestData = errors.New("2700")
	ErrFileDeleteS3Error             = errors.New("2701")

	ErrCalendarConvertNotValidRequestData = errors.New("2800")

	ErrCalendarConvertFromLinkNotValidRequestData = errors.New("2900")

	ErrBookingPostDbError             = errors.New("3000")
	ErrBookingPostNotValidRequestData = errors.New("3001")
	ErrBookingPostForbidden           = errors.New("3002")
)

//nolint:gochecknoglobals // Пока добавил в игнор.
var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound:     "Метод не найден",
	ErrUserNotAuthorized: "Пользователь не авторизован",

	ErrUsersGetDbError:   "Не удалось получить список пользователей",
	ErrUsersGetForbidden: "Недостаточно прав",

	ErrUserGetNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserGetNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserGetDbError:             "Не удалось получить информацию о пользователе",
	ErrUserGetByPhoneDbError:      "Не удалось получить информацию о пользователе",
	ErrUserGetByPhoneUserNotExist: "Пользователя с такими телефоном не существует",
	ErrUserGetForbidden:           "Недостаточно прав",

	ErrUserPostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserPostPhoneExist:          "Пользователь с переданным телефоном уже существует",
	ErrUserPostDbError:             "Не удалось создать пользователя",
	ErrUserPostCreateHash:          "Ошибка при создании пользователя",
	ErrUserPostForbidden:           "Недостаточно прав",

	ErrUserDeleteNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserDeleteNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserDeleteDbError:             "Не удалось удалить пользователя",
	ErrUserDeleteForbidden:           "Недостаточно прав",

	ErrUserPatchNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserPatchPhoneExist:          "Пользователь с переданным телефоном уже существует",
	ErrUserPatchDbError:             "Не удалось обновить пользователя",
	ErrUserPatchNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserPatchForbidden:           "Недостаточно прав",

	ErrSessionPostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrSessionPostUserNotExist:        "Пользователя с такими телефоном и паролем не существует",
	ErrSessionPostCreateJwtToken:      "Ошибка при входе в систему, попробуйте позже",

	ErrMetroStationsGetDbError: "Не удалось получить список станций метро",

	ErrCreativeSpacePostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrCreativeSpacePostDbError:             "Ошибка при создании креативной площадки",
	ErrCreativeSpacePostForbidden:           "Недостаточно прав",

	ErrCreativeSpacesGetDbError: "Не удалось получить список креативных пространств",

	ErrCreativeSpaceGetNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrCreativeSpaceGetNotFoundById:        "Креативное пространство с указанным id не найдено",
	ErrCreativeSpaceGetDbError:             "Не удалось получить информацию о креативной площадке",

	ErrCreativeSpaceDeleteNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrCreativeSpaceDeleteNotFoundById:        "Креативное пространство с указанным id не найдено",
	ErrCreativeSpaceDeleteDbError:             "Не удалось удалить информацию о креативной площадке",
	ErrCreativeSpaceDeleteS3Error:             "Не удалось удалить фотографии из s3",

	ErrCreativeSpacePatchNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrCreativeSpacePatchNotFoundById:        "Креативное пространство с указанным id не найдено",
	ErrCreativeSpacePatchDbError:             "Не удалось обновить информацию о креативной площадке",
	ErrPatchSpacePostForbidden:               "Недостаточно прав",

	ErrSessionGetNotValidRequestData: "Ошбика при валидации параметров запроса",

	ErrAddressesGetNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrAddressesGetDadataError:         "Не удалось получить список адресов",

	ErrFilePostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrFilePostOpenFileError:       "Не удалось открыть файл",
	ErrFilePostS3Error:             "Не удалось отправить файл на s3",

	ErrFileDeleteNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrFileDeleteS3Error:             "Не удалось удалить файл из s3",

	ErrCalendarConvertNotValidRequestData: "Ошбика при валидации параметров запроса",

	ErrCalendarConvertFromLinkNotValidRequestData: "Ошбика при валидации параметров запроса",

	ErrBookingPostDbError:             "Ошибка при бронировании площадки",
	ErrBookingPostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrBookingPostForbidden:           "Недостаточно прав",
}
