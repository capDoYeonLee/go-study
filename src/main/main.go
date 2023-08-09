package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n ", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

type helloWorldResponse struct {
	Message string
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello world\n")
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))

}
