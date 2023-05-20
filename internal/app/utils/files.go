package utils

import "strings"

func GetFileExtensionByFileName(fileName string) string {
	splitted := strings.Split(fileName, ".")
	return splitted[len(splitted)-1]
}
