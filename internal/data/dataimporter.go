package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/rherlt/reval/ent"
	"github.com/rherlt/reval/ent/request"
	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/persistence"
)

func ImportData() error {

	files, err := filepath.Glob(config.Current.Data_Import_Glob)
	if err != nil {
		return err
	}
	fmt.Println("Import from Files:", files)

	ctx := context.Background()
	client, _ := persistence.GetClient()

	if err != nil {
		fmt.Printf("Import from Files: %s", err)
	}

	for _, file := range files {
		importFromFile(ctx, client, file)
	}

	return err
}

// TODO fix
func createUser(ctx context.Context, client *ent.Client) error {
	u, err := client.User.Create().
		SetName("user").
		Save(ctx)

	var _ = u

	return err
}

func importFromFile(ctx context.Context, client *ent.Client, filename string) {
	evals := LoadDataFromFile(filename)
	fmt.Printf("start import %d records...\n", len(*evals))

	now := time.Now()
	scDescription := fmt.Sprintf("This scenario with %d evaluations was imported from the file '%s' at '%s'", len(*evals), filename, now.String())

	scenario, err := client.Scenario.Create().
		SetNillableExternalId(nil).
		SetName(filename).
		SetDesctiption(scDescription).
		SetDate(time.Now()).
		Save(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, eval := range *evals {

		//try to load existing request by external id
		req, err := client.Request.Query().
			Where(request.ExternalId(eval.Id)).
			First(ctx)

		//in case of error create new record
		if err != nil {

			req, err = client.Request.Create().
				SetExternalId(eval.Id).
				SetFrom(eval.Request.From).
				SetBody(eval.Request.Body).
				SetNillableDate(tryParseTime(eval.Request.Date)).
				Save(ctx)
		}

		if err != nil {
			fmt.Println(err)
		}

		//Create response
		res, err := client.Response.Create().
			SetFrom(eval.Response.From).
			SetBody(eval.Response.Body).
			SetScenarioId(scenario.ID).
			SetRequestId(req.ID).
			SetNillableDate(tryParseTime(eval.Response.Date)).
			Save(ctx)

		if err != nil {
			fmt.Println(err)
		}

		//try to get user by external Id
		user, err := persistence.GetUserByExternalId(ctx, eval.Rating.From)
		var userId uuid.UUID

		//if user does not exists try to create one
		//in case of error create new record
		if err != nil {

			userId, err = persistence.UpsertUser(ctx, eval.Rating.From, eval.Rating.From, "LLM")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			userId = user.ID
		}

		if err != nil {
			fmt.Println(err)
		}

		//create evaluation
		err = client.Evaluation.Create().
			SetEvaluationResult(mapEvaluationResult(eval.Rating.Value)).
			SetResponseId(res.ID).
			SetNillableDate(tryParseTime(eval.Rating.Date)).
			SetUserId(userId).
			Exec(ctx)

		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println("Done...")
}

func LoadDataFromFile(path string) *[]Evaluations {
	// Open jsonFile from dataPath
	fmt.Println("Open data json from: " + path)
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	bytes, _ := io.ReadAll(jsonFile)

	var evaluations []Evaluations

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(bytes, &evaluations)

	return &evaluations
}

func mapEvaluationResult(str string) string {

	switch str {
	case "Gut", "positive":
		return string(evaluationapi.Positive)
	case "Schlecht", "negative":
		return string(evaluationapi.Negative)

	}
	return string(evaluationapi.Neutral)
}

func tryParseTime(str string) *time.Time {

	time, err := time.Parse("2006-01-02", str)

	if err != nil {
		return nil
	}

	return &time
}
