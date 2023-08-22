package statistics

import "github.com/rherlt/reval/internal/api/evaluationapi"

func CalculateRatingScore(resultStatistics *[]evaluationapi.NameValuePair) evaluationapi.RatingScore {

	var result float32 = 0
	var total int32 = 0

	for _, item := range *resultStatistics {
		total = total + item.Value
		switch item.Name {
		case string(evaluationapi.Positive):
			result = result + float32(item.Value)
		case string(evaluationapi.Negative):
			result = result + float32(item.Value)
		}
	}

	score := evaluationapi.RatingScore{
		Min:   -1,
		Value: result / float32(total),
		Max:   1,
	}

	return score
}
