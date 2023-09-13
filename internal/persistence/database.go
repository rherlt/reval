package persistence

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rherlt/reval/ent"
	"github.com/rherlt/reval/ent/evaluation"
	"github.com/rherlt/reval/ent/request"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/ent/user"
	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
)

type NotFoundError interface {
	error
}

var p_client *ent.Client
var p_err error

func SetupDb() error {

	if config.Current.Db_Type == "sqlite" {
		return setupSqlite()
	}

	return nil
}

func DbExistis() bool {

	if config.Current.Db_Type == "sqlite" {
		return dbExistsSqlite()
	}

	return false
}

func dbExistsSqlite() bool {

	//extract filename from sqlite connection string
	filename := strings.Split(strings.Split(config.Current.Db_Sqlite_Connection, "file:")[1], "?")[0]

	//check if exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetClient() (*ent.Client, error) {

	if p_client == nil {
		if config.Current.Db_Type == "sqlite" {
			p_client, p_err = ent.Open("sqlite3", config.Current.Db_Sqlite_Connection)
		}
	}

	return p_client, p_err
}

func CloseClient() error {

	if p_client != nil {
		if config.Current.Db_Type == "sqlite" {
			p_client.Close()
			p_client = nil
		}
	}

	return p_err
}

func setupSqlite() error {

	client, err := ent.Open("sqlite3", config.Current.Db_Sqlite_Connection)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return err
}

func GetNextResponse(ctx context.Context) (*ent.Response, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	var v []struct {
		Id              uuid.UUID `json:"id"`
		EvaluationCount int       `json:"evalCount"`
	}

	//get min evaluation count of all responses... smallest possible result is 0
	err = client.Response.
		Query().
		Limit(1).
		Aggregate(func(s *sql.Selector) string {
			joinT := sql.Table(response.EvaluationsTable)
			s.LeftJoin(joinT).
				On(s.C(response.FieldID), joinT.C(response.EvaluationsColumn))
			s.GroupBy(s.C(response.FieldID), joinT.C(response.EvaluationsColumn))
			s.OrderBy("evalCount")
			return sql.As(sql.Count(joinT.C(response.EvaluationsColumn)), "evalCount")
		}).
		Scan(ctx, &v)

	minEvaluationCount := v[0].EvaluationCount

	//get all responseIds with number of evaluations equal to minEvaluationCount
	err = client.Response.
		Query().
		Select(response.FieldID).
		Aggregate(func(s *sql.Selector) string {
			joinT := sql.Table(response.EvaluationsTable)

			s.LeftJoin(joinT).
				On(s.C(response.FieldID), joinT.C(response.EvaluationsColumn))
			s.GroupBy(s.C(response.FieldID), joinT.C(response.EvaluationsColumn))
			s.Having(
				sql.EQ(
					"evalCount",
					minEvaluationCount),
			)
			s.OrderBy("evalCount")
			return sql.As(sql.Count(joinT.C(response.EvaluationsColumn)), "evalCount")
		}).
		Scan(ctx, &v)

	//generate random between 0 and number of responses with minEvaluationCount
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(v))

	//load random response by its id
	response, err := client.Response.Get(ctx, v[rnd].Id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Next responseId: %s\n", response.ID)
	return response, err
}

func GetRequestById(ctx context.Context, requestId uuid.UUID) (*ent.Request, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	request, err := client.Request.Query().
		Where(request.ID(requestId)).
		First(ctx)

	return request, err
}

func GetEvaluationCountByResponseId(ctx context.Context, responseId uuid.UUID) (int, int, int, error) {

	client, err := GetClient()
	if err != nil {
		return -1, -1, -1, fmt.Errorf("failed to get database client: %w", err)
	}

	var v []struct {
		Count            int    `json:"count"`
		EvaluationResult string `json:"evaluation_result"`
	}

	err = client.Evaluation.Query().
		Where(evaluation.ResponseId(responseId)).
		GroupBy(evaluation.FieldEvaluationResult).
		Aggregate(ent.Count()).
		Scan(ctx, &v)

	positive, negative, neutral := 0, 0, 0

	for i := range v {
		switch v[i].EvaluationResult {
		case string(evaluationapi.Positive):
			positive = v[i].Count
		case string(evaluationapi.Negative):
			negative = v[i].Count
		case string(evaluationapi.Neutral):
			neutral = v[i].Count
		}
	}

	return positive, negative, neutral, nil
}

func GetResponseById(ctx context.Context, responseId uuid.UUID) (*ent.Response, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	response, err := client.Response.Query().
		Where(response.ID(responseId)).
		First(ctx)

	return response, err
}

func GetUserByExternalId(ctx context.Context, externalId string) (*ent.User, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	user, err := client.User.
		Query().
		Where(user.ExternalId(externalId)).
		First(ctx)

	return user, err
}

func UpsertUser(ctx context.Context, externalId, name, userType string) (uuid.UUID, error) {

	client, err := GetClient()
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get database client: %w", err)
	}

	//insert or update... if user exists, update name column
	userId, err := client.User.
		Create().
		SetExternalId(externalId).
		SetName(name).
		SetType(userType).
		OnConflict().
		UpdateName().
		// Use the new values that were set on create.
		ID(ctx)

	return userId, err
}

func GetDemoUser(ctx context.Context) (*ent.User, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	user, err := client.User.
		Query().
		First(ctx)

	return user, err
}

func CreateEvaluationForResponseId(ctx context.Context, responseId string, evaluationResult string, userId uuid.UUID) (*ent.Evaluation, error) {

	response, err := GetResponseById(ctx, uuid.MustParse(responseId))
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	evaluation, err := client.Evaluation.
		Create().
		SetUserID(userId).
		SetResponseID(response.ID).
		SetEvaluationResult(evaluationResult).
		SetDate(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return evaluation, err
}

func GetScenarios(ctx context.Context) ([]*ent.Scenario, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	scenarios, err := client.Scenario.
		Query().
		All(ctx)

	return scenarios, err
}

func ProgressStatistics(ctx context.Context, scenarioId uuid.UUID, totalCount int32) []evaluationapi.NameValuePair {

	client, err := GetClient()
	if err != nil {
		fmt.Errorf("failed to get database client: %w", err)
		return []evaluationapi.NameValuePair{}
	}

	var v []struct {
		Count int `json:"count"`
	}

	err = client.Evaluation.
		Query().
		Where(
			evaluation.And(
				evaluation.HasUserWith(user.Type(config.Current.Oidc_Authority)),
				evaluation.HasResponseWith(response.ScenarioId(scenarioId)),
			),
		).
		Aggregate(ent.Count()).
		Scan(ctx, &v)

	var currentCount int32 = int32(v[0].Count)

	result := [...]evaluationapi.NameValuePair{
		{
			Name:  "rated",
			Value: currentCount,
		},
		{
			Name:  "unrated",
			Value: totalCount - currentCount,
		},
	}

	return result[:]
}
