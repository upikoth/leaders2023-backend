package responses

type createFileResponseFile struct {
	Name string `json:"name"`
}

type createFileStationsResponseData struct {
	File createFileResponseFile `json:"file"`
}

func CreateFileResponseFromFileKey(key string) createFileStationsResponseData {
	res := createFileStationsResponseData{
		File: createFileResponseFile{
			Name: key,
		},
	}

	return res
}
