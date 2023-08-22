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
