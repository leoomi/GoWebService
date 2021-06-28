package pokemon

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestPost(t *testing.T) {
	w := httptest.NewRecorder()

	controller := pokemonController{}

	controller.post(w, nil, nil)

	body, _ := io.ReadAll(w.Body)

	if string(body) != "banana" {
		t.Error("wrong body")
	}
}

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()

	serviceResult := Pokemon{
		PokedexNumber: 1,
		Name:          "Name",
	}
	mockParams := ServiceMockParams{
		GetReturn: serviceResult,
	}
	service := NewServiceMock(mockParams)

	controller := pokemonController{
		service: service,
	}

	controller.get(w, nil, httprouter.Params{})

	body, _ := io.ReadAll(w.Body)

	expected, _ := json.Marshal(serviceResult)
	if bytes.Equal(expected, body) {
		t.Error("wrong body")
	}
}
