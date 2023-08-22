// Package evaluationapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.2 DO NOT EDIT.
package evaluationapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

// Defines values for PostEvaluationRequestEvaluationResult.
const (
	Negative PostEvaluationRequestEvaluationResult = "negative"
	Neutral  PostEvaluationRequestEvaluationResult = "neutral"
	Positive PostEvaluationRequestEvaluationResult = "positive"
)

// ErrorInformation Information about the error.
type ErrorInformation struct {
	// Code Some unique error code.
	Code int32 `json:"code"`

	// Message Error message.
	Message string `json:"message"`
}

// Evaluations Evaluation statistics.
type Evaluations struct {
	// NumNegative number of negative evaluations.
	NumNegative int32 `json:"numNegative"`

	// NumNeutral number of neutral evaluations.
	NumNeutral int32 `json:"numNeutral"`

	// NumPositive Number of positive evaluations.
	NumPositive int32 `json:"numPositive"`
}

// GetEvaluationResponse The data of the next evaluation.
type GetEvaluationResponse struct {
	// Evaluations Evaluation statistics.
	Evaluations Evaluations `json:"evaluations"`

	// Id Unique id of the evaluation.
	Id string `json:"id"`

	// Request Message.
	Request Message `json:"request"`

	// Response Message.
	Response Message `json:"response"`
}

// GetStatisticsResponse The statistics of the evaluations grouped by scenario
type GetStatisticsResponse struct {
	Scenarios *[]ScenarioStatistics `json:"scenarios,omitempty"`
}

// Message Message.
type Message struct {
	// Body e.g. E-Mail body. The actual message.
	Body string `json:"body"`

	// Date string containing the date when the email was received in ISO 8601 format.
	Date string `json:"date"`

	// From Name of the author.
	From string `json:"from"`

	// Subject e.g. E-Mail subject.
	Subject string `json:"subject"`
}

// NameValuePair The statistics of the evaluations
type NameValuePair struct {
	// Name Name of the category.
	Name string `json:"name"`

	// Value Amout of rated question/response pairs in this category.
	Value int32 `json:"value"`
}

// PostEvaluationRequest The result of the current evaluation.
type PostEvaluationRequest struct {
	EvaluationResult PostEvaluationRequestEvaluationResult `json:"evaluationResult"`

	// Id Unique id of the evaluation.
	Id string `json:"id"`
}

// PostEvaluationRequestEvaluationResult defines model for PostEvaluationRequest.EvaluationResult.
type PostEvaluationRequestEvaluationResult string

// RatingScore defines model for RatingScore.
type RatingScore struct {
	// Max Maximum of rating Score.
	Max float32 `json:"max"`

	// Min Minimum of rating Score.
	Min float32 `json:"min"`

	// Value Rating score of the scenario.
	Value float32 `json:"value"`
}

// ScenarioStatistics Statistics per Scenario.
type ScenarioStatistics struct {
	// Description The description of the Scenario.
	Description *string `json:"description,omitempty"`

	// Id Unique id of the evaluation.
	Id string `json:"id"`

	// Name The name of the Scenario
	Name               string          `json:"name"`
	ProgressStatistics []NameValuePair `json:"progressStatistics"`
	RatingScore        RatingScore     `json:"ratingScore"`
	ResultStatistics   []NameValuePair `json:"resultStatistics"`

	// TotalResponseCount The amount of questions and response evaluated in this scenario.
	TotalResponseCount int32 `json:"totalResponseCount"`
}

// Authorization Bearer Token.
type Authorization = string

// GetEvaluationParams defines parameters for GetEvaluation.
type GetEvaluationParams struct {
	// Authorization JWT token with authorization information.
	Authorization *Authorization `json:"Authorization,omitempty"`
}

// PostEvaluationParams defines parameters for PostEvaluation.
type PostEvaluationParams struct {
	// Authorization JWT token with authorization information.
	Authorization *Authorization `json:"Authorization,omitempty"`
}

// GetStatisticsParams defines parameters for GetStatistics.
type GetStatisticsParams struct {
	// Authorization JWT token with authorization information.
	Authorization *Authorization `json:"Authorization,omitempty"`
}

// PostEvaluationJSONRequestBody defines body for PostEvaluation for application/json ContentType.
type PostEvaluationJSONRequestBody = PostEvaluationRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Gets the next message for evaluation.
	// (GET /evaluation)
	GetEvaluation(c *gin.Context, params GetEvaluationParams)
	// Posts the evaluation of the current evaluation response.
	// (POST /evaluation)
	PostEvaluation(c *gin.Context, params PostEvaluationParams)
	// Gets the statistics.
	// (GET /statistics)
	GetStatistics(c *gin.Context, params GetStatisticsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetEvaluation operation middleware
func (siw *ServerInterfaceWrapper) GetEvaluation(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEvaluationParams

	headers := c.Request.Header

	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization Authorization
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Authorization, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Authorization: %w", err), http.StatusBadRequest)
			return
		}

		params.Authorization = &Authorization

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetEvaluation(c, params)
}

// PostEvaluation operation middleware
func (siw *ServerInterfaceWrapper) PostEvaluation(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostEvaluationParams

	headers := c.Request.Header

	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization Authorization
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Authorization, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Authorization: %w", err), http.StatusBadRequest)
			return
		}

		params.Authorization = &Authorization

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostEvaluation(c, params)
}

// GetStatistics operation middleware
func (siw *ServerInterfaceWrapper) GetStatistics(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStatisticsParams

	headers := c.Request.Header

	// ------------- Optional header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization Authorization
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for Authorization, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter Authorization: %w", err), http.StatusBadRequest)
			return
		}

		params.Authorization = &Authorization

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetStatistics(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/evaluation", wrapper.GetEvaluation)
	router.POST(options.BaseURL+"/evaluation", wrapper.PostEvaluation)
	router.GET(options.BaseURL+"/statistics", wrapper.GetStatistics)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//0qyVwOaGgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
