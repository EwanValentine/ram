package handler

import (
	"encoding/json"
	"log"
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

func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	log.Error("ERROR:", err)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func ClientError(code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       http.StatusText(code),
	}, nil
}
