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
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello world\n")
	//response := helloWorldResponse{Message: "HelloWorld"}
	//data, err := json.Marshal(response) // error가 발생하지 않는다면 err은 nil값으로 들어간다.
	//if err != nil {                     // nil 예외처리
	//	panic("Ooops")
	//}
	//fmt.Fprint(w, string(data))

	response := helloWorldResponse{Message: "HelloWorld"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}
