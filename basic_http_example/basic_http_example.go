package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const HelloWorldString = "Hello World\n"

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id, string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	/*
		data, err := json.Marshal(response)
		if err != nil {
			panic("Ooops")
		}

		fmt.Fprint(w, string(data))
	*/
}
