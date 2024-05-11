package responses

type createScoreResponseScore struct {
	ID string `json:"id"`
}

type createScoreResponseData struct {
	Score createScoreResponseScore `json:"score"`
}

func CreateScoreResponseFromStoreData(scoreID string) createScoreResponseData {
	res := createScoreResponseData{}

	res.Score = createScoreResponseScore{
		ID: scoreID,
	}

	return res
}
