package util

import (
	"bytes"
	"encoding/json"

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
	OK                 = 200
	CLIENT_ERROR       = 400
	REDIRECT_SEE_OTHER = 303
	NOT_FOUND          = 404
)

func Jsonify(body map[string]interface{}) string {
	var buf bytes.Buffer
	marshalled, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	json.HTMLEscape(&buf, marshalled)
	return buf.String()
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

func CreateResponse(status int, body map[string]interface{}, headers map[string]string) Response {
	return Response{
		StatusCode:      status,
		IsBase64Encoded: false,
		Body:            Jsonify(body),
		Headers:         CombineHeaders(defaultHeaders, headers),
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
