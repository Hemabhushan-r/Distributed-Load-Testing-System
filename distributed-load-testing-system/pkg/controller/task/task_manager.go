package task

import (
	"distributed-load-testing-system/pkg/controller/models"
	"distributed-load-testing-system/pkg/controller/worker"
	"fmt"
	"log"
)

//focal-pgdg

func Distribute(taskRequest models.TaskRequest) error {
	workers := worker.GetAvailableWorkers()
	if len(workers) == 0 {
		return fmt.Errorf("no available workers")
	}

	for _, worker_var := range workers {
		go func(w worker.Worker) {
			if err := w.ExecuteTask(taskRequest); err != nil {
				log.Printf("error executing task on worker %s: %v", w.ID, err)
			}
		}(worker_var)
	}
	return nil
}
