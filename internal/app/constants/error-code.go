package constants

import (
	"errors"
)

var (
	ErrRouteNotFound     = errors.New("1000")
	ErrUserNotAuthorized = errors.New("1100")

	ErrUsersGetDbError = errors.New("1200")

	ErrUserGetNotValidRequestData = errors.New("1300")
	ErrUserGetNotFoundById        = errors.New("1301")
	ErrUserGetDbError             = errors.New("1302")
	ErrUserGetByPhoneDbError      = errors.New("1303")
	ErrUserGetByPhoneUserNotExist = errors.New("1304")

	ErrUserPostNotValidRequestData = errors.New("1400")
	ErrUserPostPhoneExist          = errors.New("1401")
	ErrUserPostDbError             = errors.New("1402")
	ErrUserPostCreateHash          = errors.New("1403")

	ErrUserDeleteNotValidRequestData = errors.New("1500")
	ErrUserDeleteNotFoundById        = errors.New("1501")
	ErrUserDeleteDbError             = errors.New("1502")

	ErrUserPatchNotValidRequestData = errors.New("1600")
	ErrUserPatchPhoneExist          = errors.New("1601")
	ErrUserPatchDbError             = errors.New("1602")
	ErrUserPatchNotFoundById        = errors.New("1603")

	ErrSessionPostNotValidRequestData = errors.New("1700")
	ErrSessionPostUserNotExist        = errors.New("1701")
	ErrSessionPostCreateJwtToken      = errors.New("1702")

	ErrMetroStationsGetDbError = errors.New("1800")

	ErrCreativeSpacePostNotValidRequestData = errors.New("1900")
	ErrCreativeSpacePostDbError             = errors.New("1901")

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

	ErrSessionGetNotValidRequestData = errors.New("1400")

	ErrAddressesGetNotValidRequestData = errors.New("1500")
	ErrAddressesGetDadataError         = errors.New("1501")

	ErrFilePostNotValidRequestData = errors.New("1600")
	ErrFilePostOpenFileError       = errors.New("1601")
	ErrFilePostS3Error             = errors.New("1602")

	ErrFileDeleteNotValidRequestData = errors.New("1700")
	ErrFileDeleteS3Error             = errors.New("1701")
)

//nolint:gochecknoglobals // Пока добавил в игнор.
var ErrDescriptionByCode = map[error]string{
	ErrRouteNotFound:     "Метод не найден",
	ErrUserNotAuthorized: "Пользователь не авторизован",

	ErrUsersGetDbError: "Не удалось получить список пользователей",

	ErrUserGetNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserGetNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserGetDbError:             "Не удалось получить информацию о пользователе",
	ErrUserGetByPhoneDbError:      "Не удалось получить информацию о пользователе",
	ErrUserGetByPhoneUserNotExist: "Пользователя с такими телефоном не существует",

	ErrUserPostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserPostPhoneExist:          "Пользователь с переданным телефоном уже существует",
	ErrUserPostDbError:             "Не удалось создать пользователя",
	ErrUserPostCreateHash:          "Ошибка при создании пользователя",

	ErrUserDeleteNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserDeleteNotFoundById:        "Пользователь с указанным id не найден",
	ErrUserDeleteDbError:             "Не удалось удалить пользователя",

	ErrUserPatchNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrUserPatchPhoneExist:          "Пользователь с переданным телефоном уже существует",
	ErrUserPatchDbError:             "Не удалось обновить пользователя",
	ErrUserPatchNotFoundById:        "Пользователь с указанным id не найден",

	ErrSessionPostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrSessionPostUserNotExist:        "Пользователя с такими телефоном и паролем не существует",
	ErrSessionPostCreateJwtToken:      "Ошибка при входе в систему, попробуйте позже",

	ErrMetroStationsGetDbError: "Не удалось получить список станций метро",

	ErrCreativeSpacePostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrCreativeSpacePostDbError:             "Ошибка при создании креативной площадки",

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

	ErrSessionGetNotValidRequestData: "Ошбика при валидации параметров запроса",

	ErrAddressesGetNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrAddressesGetDadataError:         "Не удалось получить список адресов",

	ErrFilePostNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrFilePostOpenFileError:       "Не удалось открыть файл",
	ErrFilePostS3Error:             "Не удалось отправить файл на s3",

	ErrFileDeleteNotValidRequestData: "Ошбика при валидации параметров запроса",
	ErrFileDeleteS3Error:             "Не удалось удалить файл из s3",
}
