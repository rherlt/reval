package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/persistence"

	"github.com/gin-gonic/gin"
)

type EvaluationApiServerInterface struct {
	evaluationapi.ServerInterface
}

var evaluations []evaluationapi.GetEvaluationResponse
var currentEvaluation = 0

func LoadDataFromFile() {
	// Open jsonFile from dataPath
	fmt.Println("Open data json from: " + config.Current.DataPath)
	jsonFile, err := os.Open(config.Current.DataPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	bytes, _ := io.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(bytes, &evaluations)
}

func (si EvaluationApiServerInterface) GetServerOptions() evaluationapi.GinServerOptions {
	return evaluationapi.GinServerOptions{
		BaseURL: config.Current.Gin_Api_BaseUrl,
	}
}

func (si EvaluationApiServerInterface) GetEvaluation(c *gin.Context, params evaluationapi.GetEvaluationParams) {

	var response = si.getNext()
	c.IndentedJSON(http.StatusOK, response)
}

func (si EvaluationApiServerInterface) getNext() evaluationapi.GetEvaluationResponse {

	ctx := context.Background()

	response, err := persistence.GetNextResponse(ctx)
	if err != nil {
		fmt.Println(err)
	}

	request, err := persistence.GetRequestById(ctx, response.RequestId)

	if err != nil {
		fmt.Println(err)
	}

	var res = evaluationapi.GetEvaluationResponse{
		Id: response.ID.String(),
		Response: evaluationapi.Message{
			From:    response.From,
			Subject: response.Subject,
			Body:    response.Body,
			Date:    response.Date.GoString(),
		},
		Evaluations: evaluationapi.Evaluations{
			NumNegative: 0,
			NumNeutral:  0,
			NumPositive: 0,
		},
		Request: evaluationapi.Message{
			From:    request.From,
			Subject: request.Subject,
			Body:    request.Body,
			Date:    request.Date.GoString(),
		},
	}

	return res
}

func (si EvaluationApiServerInterface) PostEvaluation(c *gin.Context, params evaluationapi.PostEvaluationParams) {

	ctx := context.Background()
	requestBody := new(evaluationapi.PostEvaluationRequest)

	if err := c.BindJSON(&requestBody); err != nil {
		var ei = evaluationapi.ErrorInformation{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, ei)
		return
	}

	user, err := persistence.GetDemoUser(ctx)

	evaluation, err := persistence.CreateEvaluationForResponseId(ctx, requestBody.Id, string(requestBody.EvaluationResult), user.ID)
	if err == nil {
		fmt.Println(err)
	}

	if evaluation == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
