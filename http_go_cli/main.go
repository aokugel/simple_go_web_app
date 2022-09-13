package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type bird struct {
	Species     string `json:"Species"`
	Description string `json:"Description"`
}

type flock struct {
	Birds  []bird `json:"Birds"`
	Number int    `json:"Number`
}

func main() {
	endpoint := os.Args[1]
	textParse := flag.String("api endpoint", endpoint, "Where to actually make the api requeest")

	flag.Parse()
	fmt.Printf("Your cli thingy was %s \n", *textParse)

	response, getErr := http.Get("http://localhost:8080/" + *textParse)

	if getErr != nil {
		fmt.Println(getErr.Error())
		os.Exit(1)
	}

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}
	flock := flock{}

	unmarshallErr := json.Unmarshal(body, &flock)

	if unmarshallErr != nil {
		log.Fatal(unmarshallErr)
	}

	fmt.Println(flock.Birds[1].Species)
}
