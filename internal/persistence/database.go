package persistence

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rherlt/reval/ent"
	"github.com/rherlt/reval/ent/request"
	"github.com/rherlt/reval/ent/response"
	"github.com/rherlt/reval/internal/config"
)

type NotFoundError interface {
	error
}

func SetupDb() error {

	if config.Current.Db_Type == "sqlite" {
		return setupSqlite()
	}

	return nil
}

var p_client *ent.Client
var p_err error

func GetClient() (*ent.Client, error) {

	if p_client == nil {
		if config.Current.Db_Type == "sqlite" {
			p_client, p_err = ent.Open("sqlite3", config.Current.Db_Sqlite_Connection)
		}
	}

	return p_client, p_err
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

	response, err := client.Response.Query().
		Order(
			// responses without evaluations are sorted first.
			response.ByEvaluationsCount(),
		).
		First(ctx)

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

func GetDemoUser(ctx context.Context) (*ent.User, error) {

	client, err := GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %w", err)
	}

	response, err := client.User.
		Query().
		First(ctx)

	return response, err
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

	evaluation, err := client.Evaluation.Create().
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
