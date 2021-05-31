package pokemon

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/leoomi/GoWebService/src/pokemon/models"
	"github.com/leoomi/GoWebService/src/pokemon/service/mock"
)

func TestPost(t *testing.T) {
	w := httptest.NewRecorder()

	controller := pokemonController{}

	controller.post(w, nil)

	body, _ := io.ReadAll(w.Body)

	if string(body) != "bandana" {
		t.Error("wrong body")
	}
}

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()

	serviceResult := models.Pokemon{
		PokedexNumber: 1,
		Name:          "Name",
	}
	mockParams := mock.ServiceMockParams{
		GetReturn: serviceResult,
	}
	service := mock.NewServiceMock(mockParams)

	controller := pokemonController{
		service: service,
	}

	controller.get(w, nil)

	body, _ := io.ReadAll(w.Body)

	expected, _ := json.Marshal(serviceResult)
	if bytes.Equal(expected, body) {
		t.Error("wrong body")
	}
}
