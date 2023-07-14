package main

import (
	"day04/config"
	"day04/gen/restapi"
	"day04/gen/restapi/operations"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
	"net/http"
)

var candies config.Candies

func main() {
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandyAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()

	server.Port = *portFlag

	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(
		func(params operations.BuyCandyParams) middleware.Responder {
			err := candies.FillModels()
			if err != nil {
				server.Logf("unexpected internal error\n")
				return middleware.Error(http.StatusInternalServerError, nil, nil)
			}

			candyMap := candies.GetMapModels()

			money := swag.Int64Value(params.Order.Money)
			candyType := swag.StringValue(params.Order.CandyType)
			candyCount := swag.Int64Value(params.Order.CandyCount)

			if _, ok := candyMap[candyType]; !ok || candyCount <= 0 {
				badResp := &operations.BuyCandyBadRequestBody{Error: "oops... no such candies or invalid candies count :("}
				return operations.NewBuyCandyBadRequest().WithPayload(badResp)
			}

			if money < candyMap[candyType]*candyCount {
				badResp := &operations.BuyCandyPaymentRequiredBody{
					Error: fmt.Sprintf("You need %d more money!", candyCount*candyMap[candyType]-money),
				}
				return operations.NewBuyCandyPaymentRequired().WithPayload(badResp)
			}

			okResp := &operations.BuyCandyCreatedBody{
				Change: money - (candyMap[candyType] * candyCount),
				Thanks: "Thank you!",
			}
			return operations.NewBuyCandyCreated().WithPayload(okResp)
		})

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
