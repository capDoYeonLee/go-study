package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {

	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler) // DefaultServeMux.HandleFunc(pattern string, handler Handler)

	log.Printf("Server starting on port %v\n ", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	var requestHello helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestHello)
	if err != nil {
		http.Error(w, "Bad requset", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + requestHello.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// 요청의 유효성 검사
// Go의 미들웨어 체인 패턴
// 미들웨어 체인 패턴은 HTTP 요청을 중간에서 가로채고 수정, 검증, 추가 작업 등을 수행한 다음 요청을 다음 핸들러로 전달하는 역할
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)

}

type hellowWorldHandler struct{}

func newHelloWorldhandler() http.Handler {
	return hellowWorldHandler{}
}

func (h hellowWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)

}

func fetchGoogle(t *testing.T) {

	r, _ := http.NewRequest("GET", "https://google.com", nil)
	// r은 HTTP 요청(Request) 객체를 나타내는 변수

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)

	}
}
