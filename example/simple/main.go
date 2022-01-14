package main

import (
	"fmt"

	"github.com/goreq/goreq"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	g := goreq.New(
		goreq.WithBaseURL("https://my-json-server.typicode.com/hadihammurabi/flutter-webservice"),
		goreq.WithAfterResponseHandler(),
	)

	resp, err := g.Get("/contacts", nil)
	panicOnError(err)
	defer resp.Body.Close()

	var data interface{}
	err = resp.Json(&data)
	panicOnError(err)

	fmt.Println(data)
}
