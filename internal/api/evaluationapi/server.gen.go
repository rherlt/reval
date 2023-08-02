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

// Evaluations Evaluations of previous users.
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
	// Evaluations Evaluations of previous users.
	Evaluations Evaluations `json:"evaluations"`

	// Id Unique id of the evaluation.
	Id int32 `json:"id"`

	// Request Request message.
	Request Message `json:"request"`

	// Response Request message.
	Response Message `json:"response"`
}

// Message Request message.
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

// PostEvaluationRequest The result of the current evaluation.
type PostEvaluationRequest struct {
	EvaluationResult PostEvaluationRequestEvaluationResult `json:"evaluationResult"`

	// Id Unique id of the message evaluation.
	Id int32 `json:"id"`
}

// PostEvaluationRequestEvaluationResult defines model for PostEvaluationRequest.EvaluationResult.
type PostEvaluationRequestEvaluationResult string

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

// PostEvaluationJSONRequestBody defines body for PostEvaluation for application/json ContentType.
type PostEvaluationJSONRequestBody = PostEvaluationRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Gets the next message for evaluation
	// (GET /evaluation)
	GetEvaluation(c *gin.Context, params GetEvaluationParams)
	// Posts the evaluation of the current evaluation response.
	// (POST /evaluation)
	PostEvaluation(c *gin.Context, params PostEvaluationParams)
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
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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