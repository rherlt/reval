package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rherlt/reval/internal/api/evaluationapi"
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/oidc"
	"github.com/rherlt/reval/internal/persistence"

	"github.com/gin-gonic/gin"
)

type EvaluationApiServerInterface struct {
	evaluationapi.ServerInterface
}

func (si EvaluationApiServerInterface) GetServerOptions() evaluationapi.GinServerOptions {
	return evaluationapi.GinServerOptions{
		BaseURL: config.Current.Gin_Api_BaseUrl,
	}
}

func (si EvaluationApiServerInterface) GetEvaluation(c *gin.Context, params evaluationapi.GetEvaluationParams) {

	ctx := context.TODO()

	response, err := persistence.GetNextResponse(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := persistence.GetRequestById(ctx, response.RequestId)

	if err != nil {
		fmt.Println(err)
	}

	pos, neg, neu, err := persistence.GetEvaluationCountByResponseId(ctx, response.ID)

	if err != nil {
		fmt.Println(err)
	}

	var respBody = evaluationapi.GetEvaluationResponse{
		Id: response.ID.String(),
		Response: evaluationapi.Message{
			From:    response.From,
			Subject: response.Subject,
			Body:    response.Body,
			Date:    response.Date.GoString(),
		},
		Evaluations: evaluationapi.Evaluations{
			NumNegative: int32(neg),
			NumNeutral:  int32(neu),
			NumPositive: int32(pos),
		},
		Request: evaluationapi.Message{
			From:    request.From,
			Subject: request.Subject,
			Body:    request.Body,
			Date:    request.Date.GoString(),
		},
	}

	c.IndentedJSON(http.StatusOK, respBody)
}

func (si EvaluationApiServerInterface) PostEvaluation(c *gin.Context, params evaluationapi.PostEvaluationParams) {

	ctx := context.TODO()
	requestBody := new(evaluationapi.PostEvaluationRequest)

	if err := c.BindJSON(&requestBody); err != nil {
		var ei = evaluationapi.ErrorInformation{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, ei)
		return
	}

	sub := c.GetStringMap(oidc.OidcUserClaimsKey)["sub"].(string)
	name := c.GetStringMap(oidc.OidcUserClaimsKey)["name"].(string)

	userId, err := persistence.UpsertUser(ctx, sub, name, config.Current.Oidc_Authority)
	if err != nil {
		http.Error(c.Writer, "error while upserting user", http.StatusInternalServerError)
		return
	}
	evaluation, err := persistence.CreateEvaluationForResponseId(ctx, requestBody.Id, string(requestBody.EvaluationResult), userId)
	if err != nil {
		http.Error(c.Writer, "error while inserting evaluation", http.StatusInternalServerError)
		return
	}

	if evaluation == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
