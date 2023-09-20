package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rherlt/reval/ent"
	"github.com/rherlt/reval/ent/evaluation"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/scenario"
	"github.com/rherlt/reval/internal/api/evaluationapi"
)

func GetTotalResponseCountByScenarioId(ctx context.Context, scenarioId uuid.UUID) int32 {

	client, err := GetClient()
	if err != nil {
		fmt.Errorf("failed to get database client: %w", err)
		return 0
	}

	var v []struct {
		Count int32 `json:"count"`
	}

	err = client.Response.
		Query().
		Where(response.ScenarioId(scenarioId)).
		Aggregate(ent.Count()).
		Scan(ctx, &v)

	if err != nil {
		fmt.Errorf("failed to number of all scenarios: %w", err)
		return 0
	}

	return v[0].Count
}

func GetResultStatisticsByScenarioId(ctx context.Context, scenarioId uuid.UUID) []evaluationapi.NameValuePair {

	client, err := GetClient()
	if err != nil {
		fmt.Errorf("failed to get database client: %w", err)
		return []evaluationapi.NameValuePair{}
	}

	var v []struct {
		Count            int    `json:"count"`
		EvaluationResult string `json:"evaluation_result"`
	}

	err = client.Scenario.
		Query().
		Where(scenario.ID(scenarioId)).
		QueryResponses().
		QueryEvaluations().
		GroupBy(evaluation.FieldEvaluationResult).
		Aggregate(ent.Count()).
		Scan(ctx, &v)

	if err != nil {
		fmt.Errorf("failed to number of all scenarios: %w", err)
		return []evaluationapi.NameValuePair{}
	}

	res := []evaluationapi.NameValuePair{}

	for _, item := range v {
		res = append(res, evaluationapi.NameValuePair{
			Name:  item.EvaluationResult,
			Value: int32(item.Count),
		})
	}

	return res
}

func GetAgreementByScenarioId(ctx context.Context, scenarioId uuid.UUID) evaluationapi.RatingScore {

	score := evaluationapi.RatingScore{
		Min:   0,
		Value: 0,
		Max:   1,
	}

	client, err := GetClient()
	if err != nil {
		fmt.Errorf("Failed to get database client: %w", err)
		return score
	}

	// Get all responses with the specified scenarioId
	responses, err := client.Response.Query().
		Where(response.ScenarioId(scenarioId)).
		All(ctx)

	if err != nil {
		fmt.Errorf("Failed to get responses from database: %w", err)
		return score
	}

	allResponses := len(responses)
	matchingEvaluations := 0.0

	for _, response := range responses {
		evaluations, errr := client.Evaluation.Query().
			Where(evaluation.ResponseId(response.ID)).
			WithUser().
			All(ctx)

		if errr != nil {
			fmt.Errorf("Failed to get evaluations from database: %w", errr)
			continue
		}

		userPositive := 0
		userNegative := 0
		var chatGPTEval string
		var userEval string

		for _, evaluation := range evaluations {
			if evaluation.Edges.User.Name == "gpt3.5-turbo" {
				chatGPTEval = evaluation.EvaluationResult
			} else {
				if evaluation.EvaluationResult == "positive" {
					userPositive++
				} else if evaluation.EvaluationResult == "negative" {
					userNegative++
				}
			}
		}

		if (userPositive+userNegative == 0) || chatGPTEval == "" {
			allResponses--
			continue
		}

		if userPositive > userNegative {
			userEval = "positive"
		} else if userNegative < userPositive {
			userEval = "negative"
		} else {
			//userEval == "positive"
			userEval = "negative"
		}

		if userEval == chatGPTEval {
			matchingEvaluations++
		}

	}

	var agreement float64
	if allResponses == 0 {
		agreement = 0.0
	} else {
		agreement = matchingEvaluations / float64(allResponses)
	}

	score.Value = float32(agreement)

	return score
}
