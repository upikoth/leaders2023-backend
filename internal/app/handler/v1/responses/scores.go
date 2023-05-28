package responses

type createScoreResponseScore struct {
	Id int `json:"id"`
}

type createScoreResponseData struct {
	Score createScoreResponseScore `json:"score"`
}

func CreateScoreResponseFromStoreData(scoreId int) createScoreResponseData {
	res := createScoreResponseData{}

	res.Score = createScoreResponseScore{
		Id: scoreId,
	}

	return res
}
