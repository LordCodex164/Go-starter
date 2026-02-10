package main

import (
	"fmt"
	"log"
)

type LogEventError struct {
	name string
	message string
	errorType    string
}

const APIError = "apiError"

//create an error method for the LogEvent struct which allows it to implement the error interface
func (l *LogEventError) Error() string {
	fmt.Println("errorType", l.errorType)
	if l.errorType == APIError {
		return fmt.Sprintln("error while fetching api")
	}
	fmt.Println("API", APIError)
	fmt.Println("errorType", l.errorType)
	return fmt.Sprintln("something went wrong")
}

func fetchAPI(apiKey string) ([]string, error) {
	if apiKey == "" {
		return nil, &LogEventError{
			errorType: "apiError",
			message: fmt.Sprintf("error while fetching api with this key: %s", apiKey),
		}
	}
	return []string{"loading", "data"}, nil
}

func main() {
	data, err := fetchAPI("")
	if err != nil {
		log.Println(err)
		return;
	}
	fmt.Println("data>>", data)
}
