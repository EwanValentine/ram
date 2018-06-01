package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Response(data interface{}, code int) (events.APIGatewayProxyResponse, error) {
	json, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(json),
	}, nil
}

func ServerError(code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       http.StatusText(code),
	}, nil
}

func ClientError(code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       http.StatusText(code),
	}, nil
}
