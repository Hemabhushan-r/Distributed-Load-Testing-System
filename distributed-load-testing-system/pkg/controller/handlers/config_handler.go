package handlers

import (
	"distributed-load-testing-system/pkg/controller/models"
	"distributed-load-testing-system/pkg/controller/validators"
	"encoding/json"
	"net/http"
)

func CreateConfig(writer http.ResponseWriter, request *http.Request) {
	var config models.TestConfig

	if err := json.NewDecoder(request.Body).Decode(&config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateConfig(config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
