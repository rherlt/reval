package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

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

	//err = createUser(ctx, client)
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
				//SetDate()
				Save(ctx)
		}

		if err != nil {
			fmt.Println(err)
		}

		res, err := client.Response.Create().
			SetFrom(eval.Response.From).
			SetBody(eval.Response.Body).
			//SetDate()
			SetRequestId(req.ID).
			Save(ctx)

		_ = res

		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println("Done...")
}

func LoadDataFromFile(path string) *[]evaluationapi.GetEvaluationResponse {
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

	var evaluations []evaluationapi.GetEvaluationResponse

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(bytes, &evaluations)

	return &evaluations
}
