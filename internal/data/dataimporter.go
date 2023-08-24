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

func importFromFile(ctx context.Context, client *ent.Client, filename string) {
	umbrella := LoadDataFromFile(filename)
	header := umbrella.Header
	entries := umbrella.Entries

	if header == nil {
		fmt.Print("No header found... skip import\n", len(*entries))
		return
	}
	printHeader(header)

	if entries == nil {
		fmt.Print("No entries found... skip import\n", len(*entries))
		return
	}

	fmt.Printf("start import %d records...\n", len(*entries))

	name := header.Name
	if stringIsNilOrEmpty(name) {
		name = &filename
	}

	scenario, err := client.Scenario.Create().
		SetNillableExternalId(nil).
		SetName(*name).
		SetNillableDescription(header.Description).
		SetNillableDate(tryParseTime(*header.Date)).
		SetNillableSystemprompt(header.SystemPrompt).
		Save(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, entry := range *entries {

		if stringIsNilOrEmpty(entry.Id) {
			fmt.Printf("EntryId is not given, skip Entry with index: %d\n", i)
			continue
		}

		//try to load existing request by external id
		req, err := client.Request.
			Query().
			Where(request.ExternalId(*entry.Id)).
			First(ctx)

		//in case of error create new record
		if err != nil {

			if stringIsNilOrEmpty(entry.Request.Body) {
				fmt.Printf("Request.Body is not given, skip Entry with index %d and Id %s\n", i, *entry.Id)
				continue
			}

			req, err = client.Request.
				Create().
				SetExternalId(*entry.Id).
				SetBody(*entry.Request.Body).
				SetNillableFrom(entry.Request.From).
				SetNillableSubject(entry.Request.Subject).
				SetNillableDate(tryParseTime(*entry.Request.Date)).
				Save(ctx)

			if err != nil {
				fmt.Printf("Error while inserting new Request for Entry with index %d and Id %s\n", i, *entry.Id)
				fmt.Println(err)
			}

		}

		if err != nil {
			fmt.Println(err)
		}

		//Create response
		res, err := client.Response.Create().
			SetRequestID(req.ID).
			SetScenarioID(scenario.ID).
			SetBody(*entry.Response.Body).
			SetNillableFrom(entry.Response.From).
			SetNillableSubject(entry.Response.Subject).
			SetNillableDate(tryParseTime(*entry.Response.Date)).
			Save(ctx)

		if err != nil {
			fmt.Println(err)
		}

		//only import rating if available
		if entry.Rating != nil {

			//try to get user by external Id
			user, err := persistence.GetUserByExternalId(ctx, *entry.Rating.From)
			var userId uuid.UUID

			//if user does not exists try to create one
			//in case of error create new record
			if err != nil {

				userId, err = persistence.UpsertUser(ctx, *entry.Rating.From, *entry.Rating.From, "LLM")
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
				SetEvaluationResult(mapEvaluationResult(*entry.Rating.Value)).
				SetResponseId(res.ID).
				SetNillableDate(tryParseTime(*entry.Rating.Date)).
				SetUserId(userId).
				Exec(ctx)

			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Done...")
}

func LoadDataFromFile(path string) Umbrella {
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

	var umbrella Umbrella

	// we unmarshal our byteArray which contains our
	json.Unmarshal(bytes, &umbrella)

	return umbrella
}

func printHeader(header *Header) {
	fmt.Printf("File header:\nName: %s\nDate: %s\nDescription: %s\n", *header.Name, *header.Date, *header.Description)

	if header.SystemPrompt != nil || *header.SystemPrompt != "" {
		fmt.Printf("System Prompt:\n%s\n\n", *header.SystemPrompt)
	}
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

func stringIsNilOrEmpty(str *string) bool {
	if str == nil || *str == "" {
		return true
	}
	return false
}
