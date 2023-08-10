package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"

	"github.com/gin-gonic/gin"
)

type EvaluationApiServerInterface struct {
	evaluationapi.ServerInterface
}

func GetSwagger(c *gin.Context) {

	swagger, error := evaluationapi.GetSwagger()

	if error != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.PureJSON(http.StatusOK, swagger)
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

	var result = evaluations[currentEvaluation%len(evaluations)]
	currentEvaluation++
	return result
}

func (si EvaluationApiServerInterface) PostEvaluation(c *gin.Context, params evaluationapi.PostEvaluationParams) {

	requestBody := new(evaluationapi.PostEvaluationRequest)

	if err := c.BindJSON(&requestBody); err != nil {
		var ei = evaluationapi.ErrorInformation{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, ei)
		return
	}

	var found = idExists(requestBody.Id, evaluations)
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	switch requestBody.EvaluationResult {
	case evaluationapi.Negative:
		evaluations[requestBody.Id].Evaluations.NumNegative++
	case evaluationapi.Positive:
		evaluations[requestBody.Id].Evaluations.NumPositive++
	case evaluationapi.Neutral:
		evaluations[requestBody.Id].Evaluations.NumNeutral++
	}

	c.Status(http.StatusOK)
}

func idExists(value int32, data []evaluationapi.GetEvaluationResponse) (exists bool) {

	for _, search := range data {
		if search.Id == value {
			return true
		}
	}
	return false
}
