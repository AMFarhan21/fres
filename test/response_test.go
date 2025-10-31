package test

import (
	"log"
	"reflect"
	"testing"

	"github.com/AMFarhan21/fres/custom"
)

func TestResponse(t *testing.T) {
	types := custom.SuccessResponse{
		Success: true,
		Message: "Success",
		Data: map[interface{}]interface{}{
			"status": 200,
			200:      "yes",
		},
	}

	expectedSuccess := reflect.Bool
	expectedMessage := reflect.String

	if reflect.TypeOf(types.Success).Kind() != expectedSuccess {
		log.Fatal("Type of success should be boolean")
	}

	if reflect.TypeOf(types.Message).Kind() != expectedMessage {
		log.Fatal("Type of message should be string")
	}
}
