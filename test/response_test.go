package test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/AMFarhan21/fres"
)

func TestCustom(t *testing.T) {
	types := fres.SuccessResponse{
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

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func TestResponseJSON(t *testing.T) {
	result, _ := json.Marshal(fres.Response.StatusNotFound("record not found"))
	expected := `{"success":false,"message":"Resource not found","error":"record not found"}`

	log.Print(result)

	log.Print("--------------------------------------------------------")

	if string(result) != expected {
		t.Fatalf("expected %s, got %s", expected, string(result))
	}
}
