package main

import (
	"day04/gen/restapi"
	"day04/gen/restapi/operations"
	"flag"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"log"
)

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
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "World"
			}

			greeting := fmt.Sprintf("Hello, %s!", name)
			return operations.NewGetGreetingOK().WithPayload(greeting)
		})

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
