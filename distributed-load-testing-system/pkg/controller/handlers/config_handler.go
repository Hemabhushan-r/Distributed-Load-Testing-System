package handlers

import (
	"distributed-load-testing-system/pkg/controller/models"
	"distributed-load-testing-system/pkg/controller/validators"
	"distributed-load-testing-system/pkg/storage"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateConfig(writer http.ResponseWriter, request *http.Request) {
	var config models.TestConfig
	writer.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(request.Body).Decode(&config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateConfig(config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	last_inserted_id, err := storage.SaveConfig(config)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"test_id": last_inserted_id,
	}
	response_data, _ := json.Marshal(response)
	writer.Write(response_data)
}

func GetConfig(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	config, err := storage.GetConfig(id)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(config)
}

func UpdateConfig(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var config models.TestConfig
	if err := json.NewDecoder(request.Body).Decode(&config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validators.ValidateConfig(config); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.UpdateConfig(id, config); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func DeleteConfig(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	if err := storage.DeleteConfig(id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
