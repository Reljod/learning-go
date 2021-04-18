package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://thatcopy.pw/catapi/rest/"

func main() {
	response, err := http.Get(url)
	if err != nil {
		response := fmt.Sprintf("error http get: %v", err)
		panic(response)
	}

	fmt.Printf("Response type: %T\n", response)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		response := fmt.Sprintf("error reading response body: %v", err)
		panic(response)
	}

	content := string(bytes)
	fmt.Print(content)

	catApi := ExtractCatApiFromResponse(content)
	fmt.Printf("Cat API: %+v\n", catApi.Id)
}

func ExtractCatApiFromResponse(content string) CatApi {
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		response := fmt.Sprintf("error extracting: %v", err)
		panic(response)
	}

	var catApi CatApi
	errApi := json.Unmarshal(byte(content), &catApi)
	if errApi != nil {
		response := fmt.Sprintf("error decoding: %v", errApi)
		panic(response)
	}

	return catApi
}

type CatApi struct {
	Id     int64
	Url    string
	WebUrl string
	X      float64
	Y      float64
}
