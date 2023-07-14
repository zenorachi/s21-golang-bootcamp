package main

import (
	"day04/gen/restapi"
	"day04/gen/restapi/operations"
	"day04/model"
	"flag"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

var candies model.Candies

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
				log.Println(err.Error())
				// TODO: ???
				return middleware.Error(400, nil, nil)
			}

			candyMap := candies.GetMapModels()

			money := swag.Int64Value(params.Order.Money)
			candyType := swag.StringValue(params.Order.CandyType)
			candyCount := swag.Int64Value(params.Order.CandyCount)

			if _, ok := candyMap[candyType]; !ok {
				return middleware.Error(402, nil, nil)
			}

			//greeting := fmt.Sprintf("Hello, %s!", name)
			//return operations.NewGetGreetingOK().WithPayload(greeting)
		})

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
