package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/rherlt/reval/ent"
	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/persistence"
)

func main() {

	config.Configure()
	persistence.SetupDb()

	ctx := context.Background()
	client, _ := persistence.GetClient()

	defer client.Close()
	/*req, err := client.Request.Create().
		SetID("request1").
		SetFrom("chatgpt/gpt-3.5-turbo").
		SetBody("Sehr geehrte Damen und Herren,\nich wollte nur kurz nachfragen, ob meine Rechnung eingegangen ist bzw. ob die Abbuchung von meinem Konto erfolgreich war. Über eine kurze Rückmeldung wäre ich sehr dankbar.\nMit freundlichen Grüßen\n[Name des Kunden]").
		//AddResponses(res).
		Save(ctx)

	if err != nil {
		log.Panic(err)
	}

	res, err := client.Response.Create().
		SetID("response1").
		SetFrom("chatgpt/gpt-3.5-turbo").
		SetBody("Sehr geehrter Herr/Frau [Name des Kunden],\n\nvielen Dank für Ihre Anfrage. Gerne teile ich Ihnen mit, dass Ihre Rechnung erfolgreich bei uns eingegangen ist. Die Abbuchung von Ihrem Konto war ebenfalls erfolgreich.\n\nFür weitere Fragen oder Informationen stehe ich Ihnen gerne zur Verfügung.\n\nMit freundlichen Grüßen\n[Dein Name]\nKundenbetreuer Energieversorgung").
		SetRequestId(req.ID).
		Save(ctx)

	*/

	createUser(ctx, client)
	importFromFile(ctx, client, "../../tmp/Llama-2-13b-chat-hf_responses.json")
	importFromFile(ctx, client, "../../tmp/vicuna-33b-v1.3_responses.json")

}

func createUser(ctx context.Context, client *ent.Client) {
	client.User.Create().
		SetUsername("user").
		Save(ctx)
}

func importFromFile(ctx context.Context, client *ent.Client, filename string) {
	evals := LoadDataFromFile(filename)
	fmt.Printf("%d evaluations loaded from: %s", len(*evals), filename)

	for _, eval := range *evals {
		// element is the element from someSlice for where we are

		req, err := client.Request.Create().
			SetExternalId(eval.Id).
			SetFrom(eval.Request.From).
			SetBody(eval.Request.Body).
			//SetDate()
			Save(ctx)

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
