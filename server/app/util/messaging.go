package util

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

var (
	defaultHeaders = map[string]string{
		"Content-Type": "application/json",
	}
)

const (
	OK                    = 200
	REDIRECT_SEE_OTHER    = 303
	CLIENT_ERROR          = 400
	NOT_FOUND             = 404
	UNPROCESSABLE_ENTITY  = 422
	INTERNAL_SERVER_ERROR = 500
	SERVICE_UNAVAILABLE   = 503
)

func Stringify(body map[string]interface{}) string {
	var buf bytes.Buffer
	marshalled, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	json.HTMLEscape(&buf, marshalled)
	return buf.String()
}

func MustMarshal(body interface{}) []byte {
	data, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return data
}

func Empty() Response {
	return Ok(nil, nil)
}

func Ok(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(OK, body, headers)
}

func Redirect(to string) Response {
	return CreateResponse(REDIRECT_SEE_OTHER, nil, map[string]string{
		"Location": to,
	})
}

func BasicMessage(body string) Response {
	return CreateResponse(OK, map[string]interface{}{
		"message": body,
	}, nil)
}

func ClientError(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(CLIENT_ERROR, body, headers)
}

func NotFound() Response {
	return NotFoundWithHeaders(nil)
}

func NotFoundWithHeaders(headers map[string]string) Response {
	return CreateResponse(NOT_FOUND, nil, headers)
}

func UnprocessableEntity() Response {
	return UnprocessableEntityWithHeaders(nil)
}

func UnprocessableEntityWithHeaders(headers map[string]string) Response {
	return CreateResponse(UNPROCESSABLE_ENTITY, nil, headers)
}

func InternalServerError(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(INTERNAL_SERVER_ERROR, body, headers)
}

func ServiceUnavailable(body map[string]interface{}, headers map[string]string) Response {
	return CreateResponse(SERVICE_UNAVAILABLE, body, headers)
}

func CreateResponse(status int, body map[string]interface{}, headers map[string]string) Response {
	return Response{
		IsBase64Encoded: false,
		StatusCode:      status,
		Headers:         CombineHeaders(defaultHeaders, headers),
		Body:            Stringify(body),
	}
}

func Panic(err error) (response *Response, ok bool) {
	if err == nil {
		return nil, true
	}
	errorResponse := ClientError(nil, nil)
	return &errorResponse, false
}

func CombineHeaders(header1 map[string]string, header2 map[string]string) map[string]string {
	both := map[string]string{}
	for k, v := range header1 {
		both[k] = v
	}
	for k, v := range header2 {
		both[k] = v
	}
	return both
}

func ResponseIsBad(resp *http.Response) bool {
	return resp.StatusCode > 400
}
