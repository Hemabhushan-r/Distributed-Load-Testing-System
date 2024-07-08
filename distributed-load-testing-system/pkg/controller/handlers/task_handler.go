package handlers

import (
	"distributed-load-testing-system/pkg/controller/models"
	"distributed-load-testing-system/pkg/controller/task"
	"encoding/json"
	"net/http"
)

func DistributeTask(w http.ResponseWriter, r *http.Request) {
	var taskRequest models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := task.Distribute(taskRequest); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
