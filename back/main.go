/*
 * Fast api desc
 *
 *  This is a sample server Petstore server.  You can find out more about      Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters.
 *
 * API version: 1.0.0
 * Contact: verycooleagle@gamail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	sw "back/go"
	"log"
	"net/http"
)

const port = ":6060"

func main() {

	err := sw.ENV.Init()

	if (err != nil) {
		log.Fatalf("[0;31m Error on initializing envirement: %s :[39m", err)
	}

	err = sw.DatabaseInit()

	if (err != nil) {
		log.Fatalf("[0;31m Error on initializing database connection: %s :[39m", err)
	}


	log.Printf("Server started on port" + port)

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(port, router))
}
